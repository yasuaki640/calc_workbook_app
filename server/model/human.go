package model

import (
	"gorm.io/gorm"
	"time"
)

type Human struct {
	gorm.Model
	Name         string    `json:"name" validate:"required"`
	Sex          byte      `json:"sex" validate:"required"`
	Birthday     time.Time `json:"birthday"`
	Description  string    `json:"description"`
	Favorability byte      `json:"favorability" validate:"required,gte=0,lte=10"`
}

func (Human) TableName() string {
	return "humans"
}
