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
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&signs).Error
	} else {
		err = db.Order(orderQuery).Find(&signs).Error
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
	err = s.PG.Model(&sign).Where("id = ?", signID).Updates(updatedData).Error
	return
}

func (s *SignService) DeleteSign(signID string) (err error) {
	var sign models.Sign
	err = s.PG.Where("id = ?", signID).Find(&sign).Delete(&sign).Error
	s.PG.Model(&sign).Association("LexicalVideoLeft").Clear()
	s.PG.Model(&sign).Association("LexicalVideoRight").Clear()
	return
}
