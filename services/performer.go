package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

type PerformerInterface interface {
	GetPerformersList(parameters utils.GetPerformerListParameters) (performers []models.Performer, count int, err error)
	CreatePerformer(performer models.Performer) (err error)
	GetPerformer(performerID string) (performer models.Performer, err error)
	UpdatePerformer(performerID string, updatedData map[string]interface{}) (err error)
	DeletePerformer(performerID string) (err error)
}

type PerformerService struct {
	PG *gorm.DB
}

func NewPerformerService(pg *gorm.DB) PerformerInterface {
	return &PerformerService{
		PG: pg,
	}
}

func (s *PerformerService) GetPerformersList(parameters utils.GetPerformerListParameters) (performers []models.Performer, count int, err error) {

	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("performers.gender", parameters.Gender),
		utils.FilterByColumn("performers.region_id", parameters.RegionID),
		utils.SearchByColumn("performers.name", parameters.Name),
	)

	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&performers).Count(&count).Error
	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.Preload("Region").Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&performers).Error
	} else {
		err = db.Preload("Region").Order(orderQuery).Find(&performers).Error
	}
	return
}

func (s *PerformerService) GetPerformer(performerID string) (performer models.Performer, err error) {
	err = s.PG.LogMode(false).Preload("Region").Where("id = ?", performerID).Take(&performer).Error
	return
}

func (s *PerformerService) CreatePerformer(performer models.Performer) (err error) {
	err = s.PG.LogMode(true).Create(&performer).Error
	return
}

func (s *PerformerService) UpdatePerformer(performerID string, updatedData map[string]interface{}) (err error) {
	var performer models.Performer
	err = s.PG.LogMode(true).Where("id = ?", performerID).First(&performer).Updates(updatedData).Error
	return
}

func (s *PerformerService) DeletePerformer(performerID string) (err error) {
	var performer models.Performer
	err = s.PG.LogMode(true).Where("id = ?", performerID).Delete(&performer).Error
	return
}
