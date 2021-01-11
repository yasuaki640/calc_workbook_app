package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yasuaki640/go-crud/service"
	"net/http"
)

func PostList(c *gin.Context) {
	postService := service.PostService{}

	posts := postService.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"Posts": posts,
	},
	)

}
