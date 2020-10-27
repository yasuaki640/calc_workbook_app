package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	v1 := router.Group("/human/v1")
	{
		v1.POST("/add", controller.HumanAdd)
		v1.GET("/list", controller.HumanList)
		v1.PUT("/update", controller.HumanUpdate)
		v1.DELETE("/delete/:id", controller.HumanDelete)
	}

	router.Run(":8080")
}
