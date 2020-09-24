package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"github.com/yasuaki640/go-crud/controller"
)

func main() {
	router := gin.Default()

	humanRouter := router.Group("/human")
	{
		humanRouter.POST("/add", controller.HumanAdd)
	}

	router.Run(":8000")

}
