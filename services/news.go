package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

type NewsInterface interface {
	GetNewsList(parameters utils.GetNewsListParameters) (news []models.News, count int, err error)
	GetNews(newsID string) (news models.News, err error)
	CreateNews(news models.News) (err error)
	UpdateNews(newsID string, updatedData map[string]interface{}) (err error)
	DeleteNews(newsID string) (err error)
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
		Preload("Creator").
		Order("importance desc").
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&news).
		Error

	return
}

func (s *NewsService) GetNews(newsID string) (news models.News, err error) {
	err = s.PG.
		Preload("Creator").
		Where("id = ?", newsID).
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

func (s *NewsService) UpdateNews(newsID string, updatedData map[string]interface{}) (err error) {
	var news models.News
	err = s.PG.
		Where("id = ?", newsID).
		Take(&news).
		Updates(updatedData).
		Error
	return
}

func (s *NewsService) DeleteNews(newsID string) (err error) {
	var news models.News
	err = s.PG.
		Where("id = ?", newsID).
		Take(&news).
		Delete(&news).
		Error
	return
}
