package services

import (
	"ccsl/models"
	"ccsl/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// LexicalVideoInterface struct
type LexicalVideoInterface interface {
	GetVideosList(parameters utils.GetVideoListParameters) (videos []models.LexicalVideo, count int, err error)
	CreateVideo(video models.LexicalVideo) (err error)
	GetVideo(videoID string) (video models.LexicalVideo, err error)
	UpdateVideo(videoID string, updatedData map[string]interface{}) (err error)
	DeleteVideo(videoID string) (err error)
}

// LexicalVideoService implements Video interface
type LexicalVideoService struct {
	PG *gorm.DB
}

// NewLexicalVideoService returns new Video serivce
func NewLexicalVideoService(pg *gorm.DB) LexicalVideoInterface {
	return &LexicalVideoService{
		PG: pg,
	}
}

// GetVideosList returns video list
func (s *LexicalVideoService) GetVideosList(parameters utils.GetVideoListParameters) (videos []models.LexicalVideo, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.

	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("lexicons.id", parameters.LexiconID),
		utils.SearchByColumn("lexicons.chinese", parameters.Chinese),
		utils.FilterByArray("lexicons.english", parameters.English, " "),
		utils.FilterByColumn("lexicons.pos", parameters.Pos),
		utils.FilterByColumn("lexicons.initial", parameters.Initial),
		utils.FilterByColumn("performers.region_id", parameters.RegionID),
		utils.FilterByColumn("performers.gender", parameters.Gender),
		utils.FilterByColumn("performers.id", parameters.PerformerID),
		utils.FilterByColumn("lexical_videos.word_formation", parameters.WordFormation),
		utils.SearchInList("lexical_videos.morpheme", parameters.Morpheme),
		utils.FilterInList("lexical_videos.left_signs_id", parameters.LeftSignID),
		utils.FilterInList("lexical_videos.right_signs_id", parameters.RightSignID),
		utils.FilterInList("lexical_videos.left_signs_id || lexical_videos.right_signs_id", parameters.SignID),
	)

	orderQuery := fmt.Sprintf("%s %s", parameters.OrderBy, parameters.Order)

	queryExp := db.Joins("JOIN performers ON performers.id = lexical_videos.performer_id").Joins("JOIN lexicons ON lexicons.id = lexical_videos.lexicon_id").Joins("JOIN districts ON districts.code = performers.region_id").Order(orderQuery)

	err = queryExp.Find(&videos).Count(&count).Error

	if err != nil {
		return
	}

	// TODO: ORM auto_preload is too slow and generated stupid sql query, will find a way to optimize this query
	if parameters.Limit != 0 {
		err = queryExp.Set("gorm:auto_preload", true).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&videos).Error
	} else {
		err = queryExp.Set("gorm:auto_preload", true).Find(&videos).Error
	}
	return
}

// CreateVideo creates a new video
func (s *LexicalVideoService) CreateVideo(video models.LexicalVideo) (err error) {
	// Don't allow to auto create signs here
	err = s.PG.LogMode(true).Set("gorm:association_autocreate", false).Create(&video).Error
	return
}

// GetVideo returns video with given id
func (s *LexicalVideoService) GetVideo(videoID string) (video models.LexicalVideo, err error) {
	err = s.PG.LogMode(true).
		Set("gorm:auto_preload", true).
		Where("id = ?", videoID).
		Take(&video).Error
	return
}

// UpdateVideo returns video with given id
func (s *LexicalVideoService) UpdateVideo(videoID string, updatedData map[string]interface{}) (err error) {
	var video models.LexicalVideo
	db := s.PG.LogMode(true)
	// Don't allow to update signs here
	err = db.Set("gorm:association_autoupdate", false).
		Where("id = ?", videoID).
		First(&video).
		Updates(updatedData).Error
	return
}

// DeleteVideo soft deletes a video
func (s *LexicalVideoService) DeleteVideo(videoID string) (err error) {
	var video models.LexicalVideo
	db := s.PG.LogMode(true)
	err = db.Where("id = ?", videoID).
		Find(&video).
		Delete(&video).Error
	return
}
