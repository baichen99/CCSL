package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

type SystemInterface interface {
	CreateJsError(jsErr models.JsError) (err error)
	GetJsErrorList(parameters utils.GetJsErrorListParameters) (errors []models.JsError, count int, err error)
}

type SystemService struct {
	PG *gorm.DB
}

func NewSystemService(pg *gorm.DB) SystemInterface {
	return &SystemService{
		PG: pg,
	}
}

func (s *SystemService) CreateJsError(jsErr models.JsError) (err error) {
	err = s.PG.Create(&jsErr).Error
	return
}

func (s *SystemService) GetJsErrorList(parameters utils.GetJsErrorListParameters) (errors []models.JsError, count int, err error) {
	db := s.PG
	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&errors).Count(&count).Error

	if err != nil {
		return
	}
	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&errors).Error
	} else {
		err = db.Order(orderQuery).Find(&errors).Error
	}
	return
}
