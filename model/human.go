package model

import "github.com/jinzhu/gorm"

type Human struct {
	gorm.Model
	Name   string
	Sex    byte
	Status string
}
