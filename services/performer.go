package services

import (
	"ccsl/models"
	"ccsl/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

type PerformerInterface interface {
	GetPerformersList(parameters utils.GetPerformerListParameters) (performers []models.Performer, count int, err error)
	CreatePerformer(performer models.Performer) (err error)
	GetPerformer(performerID string) (performer models.Performer, err error)
	UpdatePerformer(performerID string, updatedData models.Performer) (err error)
	DeletePerformer(performerID string) (err error)
}

type PerformerService struct {
	PG *gorm.DB
}

func NewPerformerServices(pg *gorm.DB) PerformerInterface {
	return &PerformerService{
		PG: pg,
	}
}

func (s *PerformerService) GetPerformersList(parameters utils.GetPerformerListParameters) (performers []models.Performer, count int, err error) {

	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.Scopes(
		utils.FilterByColumn("gender", parameters.Gender),
		utils.SearchByColumn("region", parameters.Region),
		utils.SearchByColumn("name", parameters.Name),
	)

	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&performers).Count(&count).Error
	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := fmt.Sprintf("%s %s", parameters.OrderBy, parameters.Order)
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&performers).Error
	} else {
		err = db.Order(orderQuery).Find(&performers).Error
	}
	return
}

func (s *PerformerService) CreatePerformer(performer models.Performer) (err error) {
	err = s.PG.Create(&performer).Error
	return
}

func (s *PerformerService) GetPerformer(performerID string) (performer models.Performer, err error) {
	err = s.PG.Where("id = ?", performerID).Take(&performer).Error
	return
}

func (s *PerformerService) UpdatePerformer(performerID string, updatedData models.Performer) (err error) {
	var performer models.Performer
	err = s.PG.Where("id = ?", performerID).Take(&performer).Model(&performer).Updates(updatedData).Error
	return
}

func (s *PerformerService) DeletePerformer(performerID string) (err error) {
	var performer models.Performer
	err = s.PG.Where("id = ?", performerID).Delete(&performer).Error
	return
}
