package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)



func main() {
	router := gin.Default()

	humanRouter := router.Group("/human")
	{
		humanRouter.POST("/add", controller.HumanAdd)
	}

	router.Run(":8000")
}
