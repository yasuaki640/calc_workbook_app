package model

import (
	"gorm.io/gorm"
)

type Human struct {
	gorm.Model
	Name      string
	Sex       byte
	Status    string
}

func (Human) TableName() string {
	return "humans"
}