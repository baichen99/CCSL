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
	db := s.PG.Scopes(
		utils.FilterByColumn("language", parameters.Language),
		utils.FilterByColumn("type", parameters.Type),
		utils.FilterByColumn("column", parameters.Column),
		utils.SearchByColumn("title", parameters.Title),
		utils.SearchByColumn("text", parameters.Text),
	)
	err = db.Model(&news).Count(&count).Error
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if err != nil {
		return
	}
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&news).Error
	} else {
		err = db.Order(orderQuery).Find(&news).Error
	}
	return
}

func (s *NewsService) GetNews(newsID string) (news models.News, err error) {
	err = s.PG.Where("id = ?", newsID).Take(&news).Error
	return
}

func (s *NewsService) CreateNews(news models.News) (err error) {
	err = s.PG.Create(&news).Error
	return
}

func (s *NewsService) UpdateNews(newsID string, updatedData map[string]interface{}) (err error) {
	var news models.News
	err = s.PG.Where("id = ?", newsID).Take(&news).Model(&news).Updates(updatedData).Error
	return
}

func (s *NewsService) DeleteNews(newsID string) (err error) {
	var news models.News
	err = s.PG.Where("id = ?", newsID).Delete(&news).Error
	return
}
