package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type Human struct {
	gorm.Model
	Name   string
	Sex    byte
	Status string
}

func intDb() {
	db, err := gorm.Open("sqlite3", "sqlite.main")
	if err != nil {
		panic("Failed to Open DB.")
	}
	db.AutoMigrate(&Human{})
	defer db.Close()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "こんにちは Gin",
		})
	})
	router.Run(":8000")
}
