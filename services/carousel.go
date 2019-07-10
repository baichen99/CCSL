package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

type CarouselInterface interface {
	GetCarouselList(parameters utils.GetCarouselListParameters) (carousels []models.Carousel, count int, err error)
	GetCarousel(carouselID string) (carousel models.Carousel, err error)
	CreateCarousel(carousel models.Carousel) (err error)
	UpdateCarousel(carouselID string, updatedData models.Carousel) (err error)
	DeleteCarousel(carouselID string) (err error)
}

type CarouselService struct {
	PG *gorm.DB
}

func NewCarouselService(pg *gorm.DB) CarouselInterface {
	return &CarouselService{
		PG: pg,
	}
}

func (s *CarouselService) GetCarouselList(parameters utils.GetCarouselListParameters) (carousels []models.Carousel, count int, err error) {
	db := s.PG.Scopes(
		utils.SearchByColumn("title", parameters.Title),
	)
	err = db.Model(&carousels).Count(&count).Error
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if err != nil {
		return
	}
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&carousels).Error
	} else {
		err = db.Order(orderQuery).Find(&carousels).Error
	}
	return
}

func (s *CarouselService) GetCarousel(carouselID string) (carousel models.Carousel, err error) {
	err = s.PG.Where("id = ?", carouselID).Take(&carousel).Error
	return
}

func (s *CarouselService) CreateCarousel(carousel models.Carousel) (err error) {
	err = s.PG.Create(&carousel).Error
	return
}

func (s *CarouselService) UpdateCarousel(carouselID string, updatedData models.Carousel) (err error) {
	var carousel models.Carousel
	err = s.PG.Where("id = ?", carouselID).Take(&carousel).Model(&carousel).Updates(updatedData).Error
	return
}

func (s *CarouselService) DeleteCarousel(carouselID string) (err error) {
	var carousel models.Carousel
	err = s.PG.Where("id = ?", carouselID).Delete(&carousel).Error
	return
}
