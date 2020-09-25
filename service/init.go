package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yasuaki640/go-crud/model"
)

func Init() {
	dsn := "sqlite3/db.sqlite3"
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
	}

	db.AutoMigrate(&model.Human{})

}
