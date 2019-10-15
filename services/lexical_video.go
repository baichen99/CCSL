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
	CreateVideo(video models.LexicalVideo, leftSignsID []string, rightSignsID []string) (err error)
	GetVideo(videoID string) (video models.LexicalVideo, err error)
	UpdateVideo(videoID string, updatedData map[string]interface{}, leftSignsID []string, rightSignsID []string) (err error)
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

	if parameters.LeftSignID != "" {
		leftSignSubQuery = fmt.Sprintf("SELECT lexical_video_id FROM lexical_left_sign WHERE sign_id = '%s' ", parameters.LeftSignID)
	}

	if parameters.RightSignID != "" {
		rightSignSubQuery = fmt.Sprintf("SELECT lexical_video_id FROM lexical_right_sign WHERE sign_id = '%s'", parameters.RightSignID)
	}

	if parameters.SignID != "" {
		signSubQuery = fmt.Sprintf("SELECT lexical_video_id FROM lexical_left_sign WHERE sign_id = '%s' UNION SELECT lexical_video_id FROM lexical_right_sign WHERE sign_id = '%s'", parameters.SignID, parameters.SignID)
	}

	db := s.PG.Scopes(
		utils.FilterByColumn("lexical_words.id", parameters.WordID),
		utils.SearchByColumn("lexical_words.chinese", parameters.Chinese),
		utils.FilterByArray("lexical_words.english", parameters.English, " "),
		utils.FilterByColumn("lexical_words.pos", parameters.Pos),
		utils.FilterByColumn("performers.region_id", parameters.RegionID),
		utils.FilterByColumn("performers.gender", parameters.Gender),
		utils.FilterByColumn("performers.id", parameters.PerformerID),
		utils.FilterInSubQuery("lexical_videos.id", leftSignSubQuery),
		utils.FilterInSubQuery("lexical_videos.id", rightSignSubQuery),
		utils.FilterInSubQuery("lexical_videos.id", signSubQuery),
		utils.FilterByColumn("lexical_videos.construct_type", parameters.ConstructType),
		utils.SearchInList("lexical_videos.construct_words", parameters.ConstructWords),
	)

	orderQuery := fmt.Sprintf("%s %s", parameters.OrderBy, parameters.Order)

	queryExp := db.Joins("JOIN performers ON performers.id = lexical_videos.performer_id").Joins("JOIN lexical_words ON lexical_words.id = lexical_videos.lexical_word_id").Joins("JOIN districts ON districts.code = performers.region_id").Order("lexical_words.initial asc").Order("lexical_words.id asc").Order("performers.region_id asc").Order(orderQuery)

	err = queryExp.Find(&videos).Count(&count).Error

	if err != nil {
		return
	}

	err = queryExp.Set("gorm:auto_preload", true).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&videos).Error

	for index, video := range videos {
		if len(video.LeftSigns) != 0 {
			for _, sign := range video.LeftSigns {
				videos[index].LeftSignsID = append(videos[index].LeftSignsID, sign.ID.String())
			}
		} else {
			videos[index].LeftSignsID = []string{}
		}

		if len(video.RightSigns) != 0 {
			for _, sign := range video.RightSigns {
				videos[index].RightSignsID = append(videos[index].RightSignsID, sign.ID.String())
			}
		} else {
			videos[index].RightSignsID = []string{}
		}
	}

	return
}

// CreateVideo creates a new video
func (s *LexicalVideoService) CreateVideo(video models.LexicalVideo, leftSignsID []string, rightSignsID []string) (err error) {
	err = s.PG.Set("gorm:association_autocreate", false).Create(&video).Error
	return
}

// GetVideo returns video with given id
func (s *LexicalVideoService) GetVideo(videoID string) (video models.LexicalVideo, err error) {
	err = s.PG.Set("gorm:auto_preload", true).Where("id = ?", videoID).Take(&video).Error
	if len(video.LeftSigns) != 0 {
		for _, sign := range video.LeftSigns {
			video.LeftSignsID = append(video.LeftSignsID, sign.ID.String())
		}
	} else {
		video.LeftSignsID = []string{}
	}
	if len(video.RightSigns) != 0 {
		for _, sign := range video.RightSigns {
			video.RightSignsID = append(video.RightSignsID, sign.ID.String())
		}
	} else {
		video.RightSignsID = []string{}
	}
	return
}

// UpdateVideo returns video with given id
func (s *LexicalVideoService) UpdateVideo(videoID string, updatedData map[string]interface{}, leftSignsID []string, rightSignsID []string) (err error) {
	var video models.LexicalVideo
	err = s.PG.Set("gorm:association_autoupdate", false).Model(&video).Where("id = ?", videoID).Updates(updatedData).Error
	return
}

// DeleteVideo soft deletes a video
func (s *LexicalVideoService) DeleteVideo(videoID string) (err error) {
	var video models.LexicalVideo
	err = s.PG.Where("id = ?", videoID).Delete(&video).Error
	return
}
