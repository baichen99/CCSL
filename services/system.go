package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// SystemInterface struct
type SystemInterface interface {
	CreateJsError(jsErr models.JsError) (err error)
	GetJsErrorList(parameters utils.GetJsErrorListParameters) (errors []models.JsError, count int, err error)
	GetCitiesList() (jProvinces []models.JSONProvince, err error)
	GetAppInfo(key string) (data models.Info, err error)
	UpdateAppInfo(key string, updatedData map[string]interface{}) (err error)
}

// SystemService implements system interface
type SystemService struct {
	PG *gorm.DB
}

func NewSystemService(pg *gorm.DB) SystemInterface {
	return &SystemService{
		PG: pg,
	}
}

func (s *SystemService) GetAppInfo(key string) (data models.Info, err error) {
	err = s.PG.
		Where("key = ?", key).
		Take(&data).
		Error
	return
}

func (s *SystemService) UpdateAppInfo(key string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Info{}).
		Where("key = ?", key).
		Updates(updatedData).
		Error
	return
}

func (s *SystemService) CreateJsError(jsErr models.JsError) (err error) {
	err = s.PG.
		Create(&jsErr).
		Error
	return
}

func (s *SystemService) GetJsErrorList(parameters utils.GetJsErrorListParameters) (errors []models.JsError, count int, err error) {
	db := s.PG
	// Fetching the total number of rows based on the conditions provided.
	err = db.
		Model(&errors).
		Count(&count).
		Error

	if err != nil {
		return
	}
	// Fetching the items to be returned by the query.

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&errors).
		Error

	return
}

// GetCitiesList returns json cities
func (s *SystemService) GetCitiesList() (jProvinces []models.JSONProvince, err error) {
	db := s.PG
	var provinces []models.Province
	err = db.Order("code asc").Find(&provinces).Error
	if err != nil {
		return
	}
	for _, p := range provinces {
		var cities []models.City
		var jCities []models.JSONCity
		err = db.Where("province_code = ?", p.Code).Find(&cities).Error
		if err != nil {
			return
		}
		for _, c := range cities {
			var districts []models.District
			var jDistricts []models.JSONDistrict
			err = db.Where("city_code = ?", c.Code).Find(&districts).Error
			if err != nil {
				return
			}
			for _, d := range districts {
				jDistricts = append(jDistricts, models.JSONDistrict{
					LabelZh: d.Name,
					Value:   d.Code,
				})
			}
			jCities = append(jCities, models.JSONCity{
				LabelZh:  c.Name,
				Value:    c.Code,
				Children: jDistricts,
			})
		}
		jProvinces = append(jProvinces, models.JSONProvince{
			LabelZh:  p.Name,
			Value:    p.Code,
			Children: jCities,
		})
	}
	return
}
