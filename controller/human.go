package controller

import (
	"../model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
)

func HumanAdd(c *gin.Context)  {
	human := model.Human{}
	err := c.Bind(&human)
	if err != nil{
		c.String(http.StatusBadRequest,"Bad request")
		return
	}
}