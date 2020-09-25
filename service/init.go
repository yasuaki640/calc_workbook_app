package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yasuaki640/go-crud/model"
)

	dsn := "./sqlite3/db/sqlite3"
func Init() {
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
	}

	db.AutoMigrate(&model.Human{})

}
