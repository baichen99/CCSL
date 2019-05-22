package services

import "github.com/jinzhu/gorm"

type NewsInterface interface {
}

type NewsService struct {
	PG *gorm.DB
}

func NewNewsService(pg *gorm.DB) NewsInterface {
	return &NewsService{
		PG: pg,
	}
}
