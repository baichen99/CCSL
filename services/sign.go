package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// SignInterface struct
type SignInterface interface {
	GetSignList(parameters utils.GetSignListParameters) (signs []models.Sign, count int, err error)
	CreateSign(sign models.Sign) (err error)
	GetSign(signID string) (sign models.Sign, err error)
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

func (s *SignService) GetSignList(parameters utils.GetSignListParameters) (signs []models.Sign, count int, err error) {
	db := s.PG.Scopes(
		utils.SearchByColumn("signs.name", parameters.Name),
	)
	err = db.Model(&signs).Count(&count).Error
	if err != nil {
		return
	}
	orderQuery := parameters.OrderBy + " " + parameters.Order
	// nameOrder is used to sort signs list by checking its contained number through a string typed value
	// If we don't do like this, signs order would be 1, 10, 11, ..., 2, 20, ..., ZH
	// This regular expression can help sort signs as we expected: 1, 2, 3, ..., 11, ..., ZH
	const nameOrder = `cast(substring(signs.name, '^\d+') as int) asc`
	if parameters.Limit != 0 {
		err = db.Order(nameOrder).Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&signs).Error
	} else {
		err = db.Order(nameOrder).Order(orderQuery).Find(&signs).Error
	}

	return
}

func (s *SignService) CreateSign(sign models.Sign) (err error) {
	err = s.PG.Create(&sign).Error
	return
}

func (s *SignService) GetSign(signID string) (sign models.Sign, err error) {
	err = s.PG.Where("id = ?", signID).Take(&sign).Error
	return
}

func (s *SignService) UpdateSign(signID string, updatedData map[string]interface{}) (err error) {
	var sign models.Sign
	err = s.PG.Where("id = ?", signID).First(&sign).Updates(updatedData).Error
	return
}

func (s *SignService) DeleteSign(signID string) (err error) {
	var sign models.Sign
	db := s.PG
	err = db.Where("id = ?", signID).Find(&sign).Delete(&sign).Error
	// Delete all the video associations related to this sign
	db.Model(&sign).Association("LexicalVideoLeft").Clear()
	db.Model(&sign).Association("LexicalVideoRight").Clear()
	return
}
