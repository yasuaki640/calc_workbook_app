package service

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := "host=" + os.Getenv("DB_HOST") + " " +
		"user=" + os.Getenv("DB_USER") + " " +
		"password=" + os.Getenv("DB_PASSWORD") + " " +
		"dbname=" + os.Getenv("DB_NAME") + " " +
		"port=" + os.Getenv("DB_PORT") + " " +
		"sslmode=" + os.Getenv("DB_SSL_MODE") + " " +
		"TimeZone=" + os.Getenv("DB_TIME_ZONE")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
