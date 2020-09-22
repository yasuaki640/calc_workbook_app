package main

import (
	"go-crud/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
)

func main() {
	router := gin.Default()

	humanRouter := router.Group("/human")
	{
		humanRouter.POST("/add", controller.HumanAdd)
	}

	router.Run(":8000")

}
