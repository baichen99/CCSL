package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// HandshapeInterface struct
type HandshapeInterface interface {
	GetHandshapesList(parameters utils.GetHandshapeListParameters) (handshapes []models.Handshape, count int, err error)
	CreateHandshape(handshape models.Handshape) (err error)
	GetHandshape(id string) (handshape models.Handshape, err error)
	UpdateHandshape(id string, updatedData map[string]interface{}) (err error)
	DeleteHandshape(id string) (err error)
}

// HandshapeService implements user interface
type HandshapeService struct {
	PG *gorm.DB
}

// NewHandshapeService return a new HandshapeService
func NewHandshapeService(pg *gorm.DB) HandshapeInterface {
	return &HandshapeService{
		PG: pg,
	}
}

func (s *HandshapeService) GetHandshapesList(parameters utils.GetHandshapeListParameters) (handshapes []models.Handshape, count int, err error) {
	db := s.PG.
		Scopes(
			utils.SearchByColumn("handshapes.name", parameters.Name),
		)

	err = db.
		Model(&models.Handshape{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	// nameOrder is used to sort handshapes list by checking its contained number through a string typed value
	// If we don't do like this, handshapes order would be 1, 10, 11, ..., 2, 20, ..., ZH
	// This regular expression can help sort handshapes as we expected: 1, 2, 3, ..., 11, ..., ZH
	const nameOrder = `cast(substring(handshapes.name, '^\d+') as int) asc`

	err = db.
		Order(nameOrder).
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&handshapes).Error

	return
}

func (s *HandshapeService) CreateHandshape(handshape models.Handshape) (err error) {
	err = s.PG.
		Create(&handshape).
		Error
	return
}

func (s *HandshapeService) GetHandshape(id string) (handshape models.Handshape, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&handshape).
		Error
	return
}

func (s *HandshapeService) UpdateHandshape(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Handshape{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

func (s *HandshapeService) DeleteHandshape(id string) (err error) {
	var handshape models.Handshape
	tx := s.PG.Begin()
	err = tx.
		Where("id = ?", id).
		Delete(&handshape).
		Error
	if err != nil {
		tx.Rollback()
		return
	}

	// Delete all the video associations related to this handshape
	err = tx.Table("lexical_videos").
		UpdateColumn(
			"left_handshapes_id",
			gorm.Expr("array_remove(left_handshapes_id, ? )", id),
		).
		UpdateColumn(
			"right_handshapes_id",
			gorm.Expr("array_remove(right_handshapes_id, ? )", id),
		).
		Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
