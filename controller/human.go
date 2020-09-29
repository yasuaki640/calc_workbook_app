package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/service"
	"net/http"
)

func HumanList(c *gin.Context) {
	humanService := service.HumanService{}
	humanList := humanService.GetHumans()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"message":   "success",
		"humanList": humanList,
	},
	)
}
