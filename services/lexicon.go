package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// LexiconInterface struct
type LexiconInterface interface {
	GetWordsList(parameters utils.GetLexiconListParameters) (words []models.Lexicon, count int, err error)
	CreateWord(word models.Lexicon) (err error)
	GetWord(id string) (word models.Lexicon, err error)
	UpdateWord(id string, updatedData map[string]interface{}) (err error)
	DeleteWord(id string) (err error)
}

// LexiconService implements lexicon interface
type LexiconService struct {
	PG *gorm.DB
}

// NewLexiconService returns new lexicon serivce
func NewLexiconService(pg *gorm.DB) LexiconInterface {
	return &LexiconService{
		PG: pg,
	}
}

// GetWordsList returns words list
func (s *LexiconService) GetWordsList(parameters utils.GetLexiconListParameters) (words []models.Lexicon, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.
		Scopes(
			utils.FilterInList("lexicons.pos", parameters.Pos),
			utils.FilterByColumn("lexicons.initial", parameters.Initial),
			utils.SearchByColumn("lexicons.chinese", parameters.Chinese),
			utils.FilterByArray("lexicons.english", parameters.English, " "),
		)

	// Fetching the total number of rows based on the conditions provided.
	err = db.
		Model(&models.Lexicon{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&words).
		Error

	return
}

// GetWord returns word with given id
func (s *LexiconService) GetWord(id string) (word models.Lexicon, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&word).
		Error
	return
}

// CreateWord creates a new word
func (s *LexiconService) CreateWord(word models.Lexicon) (err error) {
	err = s.PG.
		Create(&word).
		Error
	return
}

// UpdateWord updates word with given id
func (s *LexiconService) UpdateWord(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Lexicon{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteWord soft deletes a word with given id
func (s *LexiconService) DeleteWord(id string) (err error) {
	var word models.Lexicon
	err = s.PG.
		Where("id = ?", id).
		Delete(&word).
		Error
	return
}
