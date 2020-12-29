package model

import (
	"gorm.io/gorm"
)

type Human struct {
	gorm.Model
	Name   string `json:"name" validate:"required"`
	Sex    byte   `json:"sex" validate:"required"`
	Status string `json:"status" validate:"required"`
}

func (Human) TableName() string {
	return "humans"
}
