package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/service"
	"net/http"
)

func HumanList(c *gin.Context) {
	humanService := service.HumanService{}
	humanList := humanService.GetHumans()
	c.JSONP(http.StatusOK, gin.H{
		"message":   "success get all humans",
		"humanList": humanList,
	})
}
