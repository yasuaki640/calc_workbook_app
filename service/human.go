package service

import (
	"github.com/yasuaki640/go-crud/model"
)

type HumanService struct{}



func (HumanService) GetHumans() []model.Human {
	db := Init()
	humans := make([]model.Human, 0)
	db.Find(&humans)
	return humans
}
