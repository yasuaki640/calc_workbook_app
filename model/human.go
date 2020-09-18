package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Human struct {
	gorm.Model
	Name   string
	Sex    byte
	Status string
	CreatedAt time.Time
	UpdateAt time.Time
	DeletedAt sql.NullTime

}
