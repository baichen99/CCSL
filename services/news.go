package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

type NewsInterface interface {
	GetNewsList(parameters utils.GetNewsListParameters) (news []models.News, count int, err error)
	GetNews(id string) (news models.News, err error)
	CreateNews(news models.News) (err error)
	UpdateNews(id string, updatedData map[string]interface{}) (err error)
	DeleteNews(id string) (err error)
}

type NewsService struct {
	PG *gorm.DB
}

func NewNewsService(pg *gorm.DB) NewsInterface {
	return &NewsService{
		PG: pg,
	}
}

func (s *NewsService) GetNewsList(parameters utils.GetNewsListParameters) (news []models.News, count int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("news.language", parameters.Language),
			utils.FilterByColumn("news.type", parameters.Type),
			utils.FilterByColumn("news.column", parameters.Column),
			utils.SearchByColumn("news.title", parameters.Title),
			utils.SearchByColumn("news.text", parameters.Text),
			utils.FilterByColumn("news.state", parameters.State),
		)

	err = db.
		Model(&models.News{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Preload("Creator", func(pg *gorm.DB) *gorm.DB {
			return pg.Select("id, name")
		}).
		Order("importance desc").
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&news).
		Error

	return
}

func (s *NewsService) GetNews(id string) (news models.News, err error) {
	err = s.PG.
		Preload("Creator", func(pg *gorm.DB) *gorm.DB {
			return pg.Select("id, name")
		}).
		Where("id = ?", id).
		Take(&news).
		Error
	return
}

func (s *NewsService) CreateNews(news models.News) (err error) {
	err = s.PG.
		Create(&news).
		Error
	return
}

func (s *NewsService) UpdateNews(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.News{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

func (s *NewsService) DeleteNews(id string) (err error) {
	err = s.PG.
		Where("id = ?", id).
		Delete(&models.News{}).
		Error
	return
}
