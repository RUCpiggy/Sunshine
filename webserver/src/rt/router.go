package rt

import (
	"database/sql"
	"net/http"
	"sqlite"
	"strings"
	"structure"

	"github.com/gin-gonic/gin"
)

func GetRouter(router *gin.Engine, db *sql.DB) {

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
				c.JSON(http.StatusOK, gin.H{"state": "T", "message": "注册成功"})
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
				c.JSON(http.StatusOK, gin.H{"state": "T", "message": "登陆成功"})
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

	router.GET("/detail", func(c *gin.Context) {
		artworkID := c.Query("ArtworkID")
		judge := " WHERE artworkID = " + artworkID
		artworks, check := sqlite.SearchArt(db, judge)
		if check == true {
			c.JSON(http.StatusOK, gin.H{"state": "T", "message": artworks})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": "F", "message": "服务器错误"})
		}
	})

}
