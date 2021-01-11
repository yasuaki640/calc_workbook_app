package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	HumanId uint `gorm:"foreignKey:posts_humans_id_fk"`
	Caption string
}

func (Article) TableName() string {
	return "articles"
}
