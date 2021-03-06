package service

import (
	"github.com/yasuaki640/go-crud/model"
)

type HumanService struct{}

func (hs *HumanService) InsertHuman(human *model.Human) error {
	db := InitDB()

	result := db.Create(&human)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (hs *HumanService) GetHumans() []model.Human {
	db := InitDB()

	humans := make([]model.Human, 0)
	db.Find(&humans)
	return humans
}

func (hs *HumanService) UpdateHuman(h *model.Human) error {
	db := InitDB()

	result := db.Updates(h)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (hs *HumanService) DeleteHuman(id int) error {
	db := InitDB()

	result := db.Delete(&model.Human{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
