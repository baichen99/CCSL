package services

import (
	"ccsl/models"
	"ccsl/utils"
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
)

// LexicalVideoInterface struct
type LexicalVideoInterface interface {
	GetVideosList(parameters utils.GetVideoListParameters) (videos []models.LexicalVideo, count int, err error)
	CreateVideo(video models.LexicalVideo, leftSignsID []string, rightSignsID []string) (err error)
	GetVideo(videoID string) (video models.LexicalVideo, err error)
	UpdateVideo(videoID string, updatedData map[string]interface{}, leftSignsID *[]string, rightSignsID *[]string) (err error)
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

	var (
		leftSignSubQuery  string
		rightSignSubQuery string
		signSubQuery      string
	)

	// filter by left sign id
	if parameters.LeftSignID != "" {
		leftSignSubQuery = fmt.Sprintf("SELECT lexical_video_id FROM lexical_left_sign WHERE sign_id = '%s' ", parameters.LeftSignID)
	}

	// filter by right sign id
	if parameters.RightSignID != "" {
		rightSignSubQuery = fmt.Sprintf("SELECT lexical_video_id FROM lexical_right_sign WHERE sign_id = '%s'", parameters.RightSignID)
	}

	// filter by either of them
	if parameters.SignID != "" {
		signSubQuery = fmt.Sprintf("SELECT lexical_video_id FROM lexical_left_sign WHERE sign_id = '%s' UNION SELECT lexical_video_id FROM lexical_right_sign WHERE sign_id = '%s'", parameters.SignID, parameters.SignID)
	}

	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("lexicons.id", parameters.LexiconID),
		utils.SearchByColumn("lexicons.chinese", parameters.Chinese),
		utils.FilterByArray("lexicons.english", parameters.English, " "),
		utils.FilterByColumn("lexicons.pos", parameters.Pos),
		utils.FilterByColumn("lexicons.initial", parameters.Initial),
		utils.FilterByColumn("performers.region_id", parameters.RegionID),
		utils.FilterByColumn("performers.gender", parameters.Gender),
		utils.FilterByColumn("performers.id", parameters.PerformerID),
		utils.FilterInSubQuery("lexical_videos.id", leftSignSubQuery),
		utils.FilterInSubQuery("lexical_videos.id", rightSignSubQuery),
		utils.FilterInSubQuery("lexical_videos.id", signSubQuery),
		utils.FilterByColumn("lexical_videos.word_formation", parameters.WordFormation),
		utils.SearchInList("lexical_videos.morpheme", parameters.Morpheme),
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

	// add sign indices to struct
	for index, video := range videos {
		sort.Sort(models.Signs(video.LeftSigns))
		sort.Sort(models.Signs(video.RightSigns))
		videos[index].LeftSignsID = []string{}
		videos[index].RightSignsID = []string{}
		for _, sign := range video.LeftSigns {
			videos[index].LeftSignsID = append(videos[index].LeftSignsID, sign.ID.String())
		}
		for _, sign := range video.RightSigns {
			videos[index].RightSignsID = append(videos[index].RightSignsID, sign.ID.String())
		}
	}
	return
}

// CreateVideo creates a new video
func (s *LexicalVideoService) CreateVideo(video models.LexicalVideo, leftSignsID []string, rightSignsID []string) (err error) {
	var (
		leftSigns  []models.Sign
		rightSigns []models.Sign
	)
	// Don't allow to auto create signs here
	err = s.PG.LogMode(true).Set("gorm:association_autocreate", false).Create(&video).Where("id IN (?)", leftSignsID).Find(&leftSigns).Where("id IN (?)", rightSignsID).Find(&rightSigns).Error
	s.PG.Model(&video).Association("LeftSigns").Append(leftSigns)
	s.PG.Model(&video).Association("RightSigns").Append(rightSigns)
	return
}

// GetVideo returns video with given id
func (s *LexicalVideoService) GetVideo(videoID string) (video models.LexicalVideo, err error) {
	err = s.PG.LogMode(true).Set("gorm:auto_preload", true).Where("id = ?", videoID).Take(&video).Error
	video.LeftSignsID = []string{}
	video.RightSignsID = []string{}
	sort.Sort(models.Signs(video.LeftSigns))
	sort.Sort(models.Signs(video.RightSigns))
	// add sign indices to struct
	for _, sign := range video.LeftSigns {
		video.LeftSignsID = append(video.LeftSignsID, sign.ID.String())
	}
	for _, sign := range video.RightSigns {
		video.RightSignsID = append(video.RightSignsID, sign.ID.String())
	}
	return
}

// UpdateVideo returns video with given id
func (s *LexicalVideoService) UpdateVideo(videoID string, updatedData map[string]interface{}, leftSignsID *[]string, rightSignsID *[]string) (err error) {
	var (
		leftSigns  []models.Sign
		rightSigns []models.Sign
		video      models.LexicalVideo
	)

	db := s.PG.LogMode(true)
	// Don't allow to update signs here
	if err = db.Set("gorm:association_autoupdate", false).Where("id = ?", videoID).First(&video).Updates(updatedData).Error; err != nil {
		return
	}
	// Update left signs associations manually
	if leftSignsID != nil {
		if err = db.Where("id IN (?)", *leftSignsID).Find(&leftSigns).Error; err != nil {
			return
		}
		if err = db.Model(&video).Association("LeftSigns").Replace(leftSigns).Error; err != nil {
			return
		}
	}
	// Update right signs associations manually
	if rightSignsID != nil {
		if err = db.Where("id IN (?)", *rightSignsID).Find(&rightSigns).Error; err != nil {
			return
		}
		if err = db.Model(&video).Association("RightSigns").Replace(rightSigns).Error; err != nil {
			return
		}
	}
	return
}

// DeleteVideo soft deletes a video
func (s *LexicalVideoService) DeleteVideo(videoID string) (err error) {
	var video models.LexicalVideo
	db := s.PG.LogMode(true)
	err = db.Where("id = ?", videoID).Find(&video).Delete(&video).Error
	// Delete related associations
	db.Model(&video).Association("LeftSigns").Clear()
	db.Model(&video).Association("RightSigns").Clear()
	return
}
