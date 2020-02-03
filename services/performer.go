package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// PerformerInterface struct
type PerformerInterface interface {
	GetPerformersList(parameters utils.GetPerformerListParameters) (performers []models.Performer, count int, err error)
	CreatePerformer(performer models.Performer) (err error)
	GetPerformer(id string) (performer models.Performer, err error)
	UpdatePerformer(id string, updatedData map[string]interface{}) (err error)
	DeletePerformer(id string) (err error)
}

// PerformerService implements performer interface
type PerformerService struct {
	PG *gorm.DB
}

// NewPerformerService returns new performer service
func NewPerformerService(pg *gorm.DB) PerformerInterface {
	return &PerformerService{
		PG: pg,
	}
}

// GetPerformersList returns all performers
func (s *PerformerService) GetPerformersList(parameters utils.GetPerformerListParameters) (performers []models.Performer, count int, err error) {

	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.
		Scopes(
			utils.FilterByColumn("performers.gender", parameters.Gender),
			utils.FilterByColumn("performers.region_id", parameters.RegionID),
			utils.SearchByColumn("performers.name", parameters.Name),
		)

	// Fetching the total number of rows based on the conditions provided.
	err = db.
		Model(&models.Performer{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&performers).
		Error

	return
}

// GetPerformer get performer by id
func (s *PerformerService) GetPerformer(id string) (performer models.Performer, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&performer).
		Error
	return
}

// CreatePerformer creates performer
func (s *PerformerService) CreatePerformer(performer models.Performer) (err error) {
	err = s.PG.
		Create(&performer).
		Error
	return
}

// UpdatePerformer updates performer with given id
func (s *PerformerService) UpdatePerformer(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Performer{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeletePerformer delete performer by id
func (s *PerformerService) DeletePerformer(id string) (err error) {
	var performer models.Performer
	err = s.PG.
		Where("id = ?", id).
		Delete(&performer).
		Error
	return
}
