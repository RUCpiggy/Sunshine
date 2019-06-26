package rt

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sqlite"
	"strconv"
	"strings"
	"structure"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRouter(router *gin.Engine, db *sql.DB) {
	//useless
	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		user, check := sqlite.SearchUser(db, name)
		if check == true {
			c.JSON(200, gin.H{
				"userID":  user.UserID,
				"name":    user.Name,
				"email":   user.Email,
				"tel":     user.Tel,
				"address": user.Address,
				"balance": user.Balance,
			})
		} else {
			c.JSON(200, gin.H{
				"check": "false",
			})
		}

	})

	router.GET("/artworks", func(c *gin.Context) {
		artMap := make(map[string]string)
		artMap["artworkID"] = c.Query("ArtworkID")
		artMap["artist"] = c.Query("Artist")
		artMap["imageFileName"] = c.Query("ImageFileName")
		artMap["title"] = c.Query("Title")
		artMap["yearOfWork"] = c.Query("YearOfWork")
		artMap["genre"] = c.Query("Genre")
		artMap["width"] = c.Query("Width")
		artMap["height"] = c.Query("Height")
		artMap["price"] = c.Query("Price")
		artMap["view"] = c.Query("View")
		artMap["ownerID"] = c.Query("OwnerID")
		artMap["orderID"] = c.Query("OrderID")
		artMap["timeReleased"] = c.Query("TimeReleased")
		judge := ""
		var sliceJudge []string
		for index, value := range artMap {
			if value != "" {
				singleJudge := index + "=" + "'" + value + "'"
				sliceJudge = append(sliceJudge, singleJudge)
			}
		}
		judge = strings.Join(sliceJudge, " and ")
		judge = " WHERE " + judge
		artworks, check := sqlite.SearchArt(db, judge)
		if check == true {
			c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
		}
	})

	router.GET("/user/:userID/*action", func(c *gin.Context) {
		name := c.Param("userID")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/register", func(c *gin.Context) {
		var user structure.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, check := sqlite.SearchUser(db, user.Name)
		if check == false && (len(user.Name) > 5) {
			user.UserID = sqlite.MaxUserID(db) + 1
			sqlite.AddUser(db, user.UserID, user.Name, user.Email, user.Password, user.Tel, user.Address, user.Balance)
			_, check = sqlite.SearchUser(db, user.Name)
			if check == true {
				userIDSting, tocken, t := genTocken(user.UserID)
				c.JSON(http.StatusOK, gin.H{"state": "T", "message": "注册成功", "tocken": userIDSting + "#" + tocken + "#" + t})
			} else {
				c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户已存在"})
		}

	})

	router.POST("/login", func(c *gin.Context) {
		var user structure.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userInDB, check := sqlite.SearchUser(db, user.Name)
		if check == true {
			if userInDB.Password == user.Password {
				userIDString, tocken, t := genTocken(userInDB.UserID)
				c.JSON(http.StatusOK, gin.H{"state": "T", "message": "登陆成功", "tocken": userIDString + "#" + tocken + "#" + t})
			} else {
				c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户名或密码错误"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户名或密码错误"})
		}

	})

	router.StaticFS("/img", http.Dir("../../data/img"))

	router.GET("/hot", func(c *gin.Context) {
		var judge = " ORDER BY view DESC LIMIT 0,4"
		artworks, check := sqlite.SearchArt(db, judge)
		if check == true {
			c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
		}
	})

	router.GET("/new", func(c *gin.Context) {
		var judge = " ORDER BY timeReleased DESC LIMIT 0,3"
		artworks, check := sqlite.SearchArt(db, judge)
		if check == true {
			c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
		}
	})

	router.GET("/detail", func(c *gin.Context) {
		artworkID := c.Query("ArtworkID")
		judge := " WHERE artworkID = " + artworkID
		artworks, check := sqlite.SearchArt(db, judge)
		if check == true {
			sqlite.UpdateArtInfo(db, artworkID, "view", strconv.Itoa(artworks[0].View+1))
			c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
		}
	})

	router.POST("/chart", func(c *gin.Context) {
		artworkID := c.PostForm("ArtworkID")
		tocken := c.PostForm("tocken")
		chartID := sqlite.MaxCart(db)
		chartID = chartID + 1
		userID, check := isLogin(tocken)
		if check == true {
			judge := " WHERE artworkID = " + artworkID + " AND userID = " + userID
			_, check = sqlite.SearchCart(db, judge)
			if check == true {
				c.JSON(http.StatusOK, gin.H{"state": "F", "message": "该商品已经被添加进购物车"})
			} else {
				check = sqlite.AddCart(db, chartID, userID, artworkID)
				if check == true {
					c.JSON(http.StatusOK, gin.H{"state": "T", "message": "成功添加到购物车"})
				}
			}

		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户认证失败，请重新登陆"})
		}

	})

	router.GET("/search", func(c *gin.Context) {
		title := c.DefaultQuery("title", "0")
		description := c.DefaultQuery("description", "0")
		artist := c.DefaultQuery("artist", "0")
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		artMap := make(map[string]string)
		if title != "0" {
			artMap["title"] = title
		}
		if description != "0" {
			artMap["description"] = description
		}
		if artist != "0" {
			artMap["artist"] = artist
		}
		if (artist == "0" || artist == "") && (description == "0" || description == "") && (title == "0" || title == "") {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "请输入关键词"})
		} else {
			judge := ""
			var sliceJudge []string
			for index, value := range artMap {
				if value != "" {
					singleJudge := index + " LIKE " + "'%" + value + "%'"
					sliceJudge = append(sliceJudge, singleJudge)
				}
			}
			judge = strings.Join(sliceJudge, " and ")
			judge = " WHERE " + judge
			artworks, check := sqlite.SearchArt(db, judge)
			if check == true {
				if len(artworks) < 10 {
					c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks, "total": len(artworks)})
				} else {
					var groupNum int
					if len(artworks)%10 == 0 {
						groupNum = len(artworks) / 10
					} else {
						groupNum = len(artworks)/10 + 1
					}
					if page < 1 {
						page = 1
					}
					if page >= groupNum {
						c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks[10*(groupNum-1):], "total": len(artworks)})
					} else {
						c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks[10*(page-1) : 10*(page)], "total": len(artworks)})
					}

				}
			} else {
				c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
			}
		}

	})

	router.POST("/user", func(c *gin.Context) {
		tocken := c.PostForm("tocken")
		userID, check := isLogin(tocken)
		if check == true {
			user, check := sqlite.SearchUserByUserID(db, userID)
			if check == true {
				c.JSON(http.StatusOK, gin.H{"state": "T", "message": user})
			} else {
				c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户认证失败，请重新登陆"})
			}

		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户认证失败，请重新登陆"})
		}

	})

	router.POST("/img", func(c *gin.Context) {
		file, header, _ := c.Request.FormFile("file")
		filename := header.Filename
		out, err := os.Create("../../data/img/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)

		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", filename))
	})

	router.POST("/artwork", func(c *gin.Context) {
		var artwork structure.Artwork
		if err := c.BindJSON(&artwork); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userIDStr, check := isLogin(artwork.Tocken)
		if check == true {
			userID, err := strconv.Atoi(userIDStr)
			artwork.OwnerID = userID
			println(err)
			artwork.ArtworkID = sqlite.MaxArtworkID(db) + 1
			sqlite.AddArt(db, artwork.ArtworkID, artwork.Artist, artwork.ImageFileName, artwork.Title, artwork.Description, artwork.YearOfWork, artwork.Genre, artwork.Width, artwork.Height, artwork.Price, "0", artwork.OwnerID, 0, strconv.FormatInt(time.Now().Unix(), 10))

			c.JSON(http.StatusOK, gin.H{"state": "T", "message": "发布成功"})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户认证失败，请重新登陆"})

		}
	})

	router.GET("/chart", func(c *gin.Context) {
		tocken := c.DefaultQuery("tocken", "0")
		if tocken != "0" {
			userID, check := isLogin(tocken)
			if check == true {
				var judge = "SELECT artworks.artworkID, artworks.title, artworks.description, artworks.imageFileName, artworks.price FROM carts LEFT JOIN artworks ON carts.artworkID = artworks.artworkID WHERE carts.userID  =  " + userID
				artworks, check := sqlite.SearchArtInChart(db, judge)
				if check == true {
					c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks})
				} else {
					c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"state": "F", "message": "用户认证失败，请重新登陆"})
			}
		}

	})
}
