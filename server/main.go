package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		humans := v1.Group("/humans")
		{
			humans.POST("", controller.HumanAdd)
			humans.GET("", controller.HumanList)
			humans.PUT("", controller.HumanUpdate)
			humans.DELETE("/:id", controller.HumanDelete)
		}
	}

	router.Run(":8000")
}
