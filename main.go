package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()

	humanRouter := router.Group("/human")
	{

	}

	router.Run(":8000")
}
