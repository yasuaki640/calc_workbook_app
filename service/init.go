package service

import (
	"../model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	dsn := "sqlite.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
	}

	db.AutoMigrate(&model.Human{})

}
