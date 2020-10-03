package service

import (
	"github.com/yasuaki640/go-crud/model"
)

type HumanService struct{}



func (HumanService) GetHumans() []model.Human {
	db := InitDB()
	humans := make([]model.Human, 0)
	db.Find(&humans)
	return humans
}

func (HumanService) InsertHuman(human *model.Human) {
	db := InitDB()
}
