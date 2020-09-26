package service

import (
	. "github.com/yasuaki640/go-crud/model"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	dsn := "sqlite3/db.sqlite3"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Human{})
	if err != nil {
		panic(err)
	}
}
