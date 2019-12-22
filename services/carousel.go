package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// CarouselInterface struct
type CarouselInterface interface {
	GetCarouselList(parameters utils.GetCarouselListParameters) (carousels []models.Carousel, count int, err error)
	GetCarousel(carouselID string) (carousel models.Carousel, err error)
	CreateCarousel(carousel models.Carousel) (err error)
	UpdateCarousel(carouselID string, updatedData map[string]interface{}) (err error)
	DeleteCarousel(carouselID string) (err error)
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

	err = db.Model(&carousels).Count(&count).Error
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if err != nil {
		return
	}
	if parameters.Limit != 0 {
		err = db.Preload("Creator").Order("importance desc").Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&carousels).Error
	} else {
		err = db.Preload("Creator").Order("importance desc").Order(orderQuery).Find(&carousels).Error
	}
	return
}

// GetCarousel returns carousel with given id
func (s *CarouselService) GetCarousel(carouselID string) (carousel models.Carousel, err error) {
	err = s.PG.Preload("Creator").Where("id = ?", carouselID).Take(&carousel).Model(&carousel).Related(&carousel.Creator).Find(&carousel.Creator).Error
	return
}

// CreateCarousel creates a new carousel
func (s *CarouselService) CreateCarousel(carousel models.Carousel) (err error) {
	err = s.PG.Create(&carousel).Error
	return
}

// UpdateCarousel updates carousel with given id
func (s *CarouselService) UpdateCarousel(carouselID string, updatedData map[string]interface{}) (err error) {
	var carousel models.Carousel
	err = s.PG.Where("id = ?", carouselID).First(&carousel).Updates(updatedData).Error
	return
}

// DeleteCarousel soft deletes a carousel with given id
func (s *CarouselService) DeleteCarousel(carouselID string) (err error) {
	var carousel models.Carousel
	err = s.PG.Where("id = ?", carouselID).Delete(&carousel).Error
	return
}
