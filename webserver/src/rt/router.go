package rt

import (
	"database/sql"
	"net/http"
	"sqlite"
	"strings"

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
		artMap["artworkID"] = c.Query("artworkID")
		artMap["artist"] = c.Query("artist")
		artMap["imageFileName"] = c.Query("imageFileName")
		artMap["title"] = c.Query("title")
		artMap["yearOfWork"] = c.Query("yearOfWork")
		artMap["genre"] = c.Query("genre")
		artMap["width"] = c.Query("width")
		artMap["height"] = c.Query("height")
		artMap["price"] = c.Query("price")
		artMap["view"] = c.Query("view")
		artMap["ownerID"] = c.Query("ownerID")
		artMap["orderID"] = c.Query("orderID")
		artMap["timeReleased"] = c.Query("timeReleased")
		judge := ""
		var sliceJudge []string
		for index, value := range artMap {
			if value != "" {
				singleJudge := index + "=" + "'" + value + "'"
				sliceJudge = append(sliceJudge, singleJudge)
			}
		}
		judge = strings.Join(sliceJudge, " and ")
		result, check := sqlite.SearchArt(db, judge)
		if check == true {
			var resultJson []gin.H
			for _, value := range result {
				singleResult := gin.H{
					"artist":        value.Artist,
					"artworkID":     value.ArtworkID,
					"description":   value.Description,
					"genre":         value.Genre,
					"height":        value.Height,
					"imageFileName": value.ImageFileName,
					"orderID":       value.OrderID,
					"ownerID":       value.OwnerID,
					"price":         value.Price,
					"timeReleased":  value.TimeReleased,
					"title":         value.Title,
					"view":          value.View,
					"width":         value.Width,
					"yearOfWork":    value.YearOfWork,
				}
				resultJson = append(resultJson, singleResult)
			}
			c.JSON(200, resultJson)
		}
	})

	router.GET("/user/:userID/*action", func(c *gin.Context) {
		name := c.Param("userID")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

}
