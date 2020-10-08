package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/model"
	"github.com/yasuaki640/go-crud/service"
	"net/http"
	"strconv"
)

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

func HumanUpdate(c *gin.Context) {
	human := model.Human{}

	err := c.Bind(&human)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	humanService := service.HumanService{}
}

func HumanList(c *gin.Context) {
	humanService := service.HumanService{}

	humanList := humanService.GetHumans()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":   "success",
		"humanList": humanList,
	})

}

func HumanDelete(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	humanService := service.HumanService{}
	err = humanService.DeleteHuman(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
