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

func (HumanService) InsertHuman(human *model.Human) error {
	db := InitDB()

	result := db.Create(&human)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s HumanService) DeleteHuman(id int) error {
	db := InitDB()
}
