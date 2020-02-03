package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// SignInterface struct
type SignInterface interface {
	GetSignList(parameters utils.GetSignListParameters) (signs []models.Handshape, count int, err error)
	CreateSign(sign models.Handshape) (err error)
	GetSign(signID string) (sign models.Handshape, err error)
	UpdateSign(signID string, updatedData map[string]interface{}) (err error)
	DeleteSign(signID string) (err error)
}

// SignService implements user interface
type SignService struct {
	PG *gorm.DB
}

// NewSignService return a new SignService
func NewSignService(pg *gorm.DB) SignInterface {
	return &SignService{
		PG: pg,
	}
}

func (s *SignService) GetSignList(parameters utils.GetSignListParameters) (signs []models.Handshape, count int, err error) {
	db := s.PG.
		Scopes(
			utils.SearchByColumn("signs.name", parameters.Name),
		)

	err = db.
		Model(&models.Handshape{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	// nameOrder is used to sort signs list by checking its contained number through a string typed value
	// If we don't do like this, signs order would be 1, 10, 11, ..., 2, 20, ..., ZH
	// This regular expression can help sort signs as we expected: 1, 2, 3, ..., 11, ..., ZH
	const nameOrder = `cast(substring(signs.name, '^\d+') as int) asc`

	err = db.
		Order(nameOrder).
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&signs).Error

	return
}

func (s *SignService) CreateSign(sign models.Handshape) (err error) {
	err = s.PG.
		Create(&sign).
		Error
	return
}

func (s *SignService) GetSign(signID string) (sign models.Handshape, err error) {
	err = s.PG.
		Where("id = ?", signID).
		Take(&sign).
		Error
	return
}

func (s *SignService) UpdateSign(signID string, updatedData map[string]interface{}) (err error) {
	var sign models.Handshape
	err = s.PG.
		Where("id = ?", signID).
		Take(&sign).
		Updates(updatedData).
		Error
	return
}

func (s *SignService) DeleteSign(signID string) (err error) {
	var sign models.Handshape
	tx := s.PG.Begin()
	err = tx.
		Where("id = ?", signID).
		Delete(&sign).
		Error
	if err != nil {
		tx.Rollback()
		return
	}

	// Delete all the video associations related to this sign
	err = tx.Table("lexical_videos").
		UpdateColumn(
			"left_signs_id",
			gorm.Expr("array_remove(left_signs_id, ? )", signID),
		).
		UpdateColumn(
			"right_signs_id",
			gorm.Expr("array_remove(right_signs_id, ? )", signID),
		).
		Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
