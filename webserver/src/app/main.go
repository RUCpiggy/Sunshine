package main

import (
	"config"
	"database/sql"
	"fmt"
	"rt"

	"github.com/gin-gonic/gin"
)

func main() {
	var version = config.Version

	fmt.Println(version)
	var path = "../../data/sql/sunshine.sqlite3"
	db, err := sql.Open("sqlite3", path)
	println(err)

	route := gin.Default()
	rt.GetRouter(route, db)
	route.Run(":6080")
}
