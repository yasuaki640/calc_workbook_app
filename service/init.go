package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	dsn := "sqlite.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

}
