package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/humans", controller.HumanAdd)
		v1.GET("/humans", controller.HumanList)
		v1.PUT("/humans", controller.HumanUpdate)
		v1.DELETE("/humans/:id", controller.HumanDelete)
	}

	router.Run(":8000")
}
