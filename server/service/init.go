package service

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=db " +
		"user=postgres " +
		"password=postgres " +
		"dbname=go-crud " +
		"port=5432 " +
		"sslmode=disable " +
		"TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
