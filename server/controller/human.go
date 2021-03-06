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
		return
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
	err = humanService.UpdateHuman(&human)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func HumanList(c *gin.Context) {
	humanService := service.HumanService{}

	humans := humanService.GetHumans()
	c.JSON(http.StatusOK, gin.H{
		"humans": humans,
	},
	)

}

func HumanDelete(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	humanService := service.HumanService{}
	err = humanService.DeleteHuman(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
