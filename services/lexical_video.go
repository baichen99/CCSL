package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// LexicalVideoInterface struct
type LexicalVideoInterface interface {
	GetLexicalVideosList(parameters utils.GetLexicalVideoListParameters) (videos []models.LexicalVideo, count int, err error)
	CreateLexicalVideo(video models.LexicalVideo) (err error)
	GetLexicalVideo(id string) (video models.LexicalVideo, err error)
	UpdateLexicalVideo(id string, updatedData map[string]interface{}) (err error)
	DeleteLexicalVideo(id string) (err error)
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

// GetLexicalVideosList returns video list
func (s *LexicalVideoService) GetLexicalVideosList(parameters utils.GetLexicalVideoListParameters) (videos []models.LexicalVideo, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.
		Scopes(
			utils.FilterByColumn("lexicons.id", parameters.LexiconID),
			utils.SearchByColumn("lexicons.chinese", parameters.Chinese),
			utils.FilterByArray("lexicons.english", parameters.English, " "),
			utils.FilterInList("lexicons.pos", parameters.Pos),
			utils.FilterByColumn("lexicons.initial", parameters.Initial),
			utils.FilterByColumn("performers.region_id", parameters.RegionID),
			utils.FilterByColumn("performers.gender", parameters.Gender),
			utils.FilterByColumn("performers.id", parameters.PerformerID),
			utils.FilterByColumn("lexical_videos.word_formation", parameters.WordFormation),
			utils.SearchInList("lexical_videos.morpheme", parameters.Morpheme),
			utils.FilterInList("lexical_videos.left_handshapes_id", parameters.LeftHandshapeID),
			utils.FilterInList("lexical_videos.right_handshapes_id", parameters.RightHandshapeID),
			utils.FilterInList("lexical_videos.left_handshapes_id || lexical_videos.right_handshapes_id", parameters.HandshapeID),
		)

	query := db.Joins("JOIN performers ON performers.id = lexical_videos.performer_id").
		Joins("JOIN lexicons ON lexicons.id = lexical_videos.lexicon_id").
		Joins("JOIN districts ON districts.code = performers.region_id")

	err = query.Model(&models.LexicalVideo{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = query.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&videos).Error

	return
}

// CreateLexicalVideo creates a new video
func (s *LexicalVideoService) CreateLexicalVideo(video models.LexicalVideo) (err error) {
	err = s.PG.
		Create(&video).Error
	return
}

// GetLexicalVideo returns video with given id
func (s *LexicalVideoService) GetLexicalVideo(id string) (video models.LexicalVideo, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&video).Error
	return
}

// UpdateLexicalVideo returns video with given id
func (s *LexicalVideoService) UpdateLexicalVideo(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.LexicalVideo{}).
		Where("id = ?", id).
		Updates(updatedData).Error
	return
}

// DeleteLexicalVideo soft deletes a video
func (s *LexicalVideoService) DeleteLexicalVideo(id string) (err error) {
	err = s.PG.
		Where("id = ?", id).
		Delete(&models.LexicalVideo{}).
		Error
	return
}
