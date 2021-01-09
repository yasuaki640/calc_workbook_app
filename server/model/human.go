package model

import (
	"gorm.io/gorm"
	"time"
)

type Human struct {
	gorm.Model
	Name         string `validate:"required"`
	Sex          byte   `validate:"required"`
	Birthday     time.Time
	Description  string
	Favorability byte `validate:"required,gte=0,lte=10"`
}

func (Human) TableName() string {
	return "humans"
}
