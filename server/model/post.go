package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	HumanID uint `gorm:"foreignKey:posts_humans_id_fk"`
	Caption string
}

func (Post) TableName() string {
	return "posts"
}
