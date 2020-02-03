package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// CarouselInterface struct
type CarouselInterface interface {
	GetCarouselList(parameters utils.GetCarouselListParameters) (carousels []models.Carousel, count int, err error)
	GetCarousel(id string) (carousel models.Carousel, err error)
	CreateCarousel(carousel models.Carousel) (err error)
	UpdateCarousel(id string, updatedData map[string]interface{}) (err error)
	DeleteCarousel(id string) (err error)
}

// CarouselService implements CarouselInterface
type CarouselService struct {
	PG *gorm.DB
}

// NewCarouselService returns new carousel serivce
func NewCarouselService(pg *gorm.DB) CarouselInterface {
	return &CarouselService{
		PG: pg,
	}
}

// GetCarouselList returns carousels list
func (s *CarouselService) GetCarouselList(parameters utils.GetCarouselListParameters) (carousels []models.Carousel, count int, err error) {
	db := s.PG.Scopes(
		utils.SearchByColumn("carousels.title_zh", parameters.TitleZh),
		utils.FilterByArray("carousels.title_en", parameters.TitleEn, " "),
		utils.FilterByColumn("carousels.state", parameters.State),
	)

	err = db.
		Model(&models.Carousel{}).
		Count(&count).Error

	if err != nil {
		return
	}

	err = db.
		Preload("Creator").
		Order("importance desc").
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&carousels).
		Error

	return
}

// GetCarousel returns carousel with given id
func (s *CarouselService) GetCarousel(id string) (carousel models.Carousel, err error) {
	err = s.PG.
		Preload("Creator").
		Where("id = ?", id).
		Take(&carousel).
		Error
	return
}

// CreateCarousel creates a new carousel
func (s *CarouselService) CreateCarousel(carousel models.Carousel) (err error) {
	err = s.PG.
		Create(&carousel).
		Error
	return
}

// UpdateCarousel updates carousel with given id
func (s *CarouselService) UpdateCarousel(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Carousel{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteCarousel soft deletes a carousel with given id
func (s *CarouselService) DeleteCarousel(id string) (err error) {
	var carousel models.Carousel
	err = s.PG.
		Where("id = ?", id).
		Delete(&carousel).
		Error
	return
}
