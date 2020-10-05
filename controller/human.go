package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/model"
	"github.com/yasuaki640/go-crud/service"
	"net/http"
)

func HumanList(c *gin.Context) {
	humanService := service.HumanService{}
	humanList := humanService.GetHumans()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":   "success",
		"humanList": humanList,
	})

}

func HumanAdd(c *gin.Context) {
	human := model.Human{}
	err := c.Bind(&human)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	humanService := service.HumanService{}
	err = humanService.InsertHuman(&human)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func HumanDelete(c *gin.Context) {
	id := c.Param("id")
}
