package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/controller"
	"github.com/yasuaki640/go-crud/service"
)

func main() {
	service.Init()
	router := gin.Default()

	v1 := router.Group("/human/v1")
	{
		v1.GET("/list", controller.HumanList)
	}

	router.Run(":8000")

}
