package service

import (
	"github.com/yasuaki640/go-crud/model"
)

type PostService struct{}

func (hs *PostService) GetPosts() []model.Post {
	db := InitDB()

	posts := make([]model.Post, 0)
	db.Find(&posts)
	return posts
}
