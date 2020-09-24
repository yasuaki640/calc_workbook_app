package controller

import (
	"github.com/yasuaki640/go-crud/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
)

func HumanAdd(c *gin.Context) {
	human := model.Human{}
	err := c.Bind(&human)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
}
