package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../model"
)

func HumanAdd(c *gin.Context)  {
	human := model.Human{}
	err := c.Bind(&human)
	if err != nil{
		c.String(http.StatusBadRequest,"Bad request")
		return
	}
}