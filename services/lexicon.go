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
	GetWord(wordID string) (word models.Lexicon, err error)
	UpdateWord(wordID string, updatedData map[string]interface{}) (err error)
	DeleteWord(wordID string) (err error)
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
	db := s.PG.Scopes(
		utils.FilterByColumn("lexicons.pos", parameters.Pos),
		utils.FilterByColumn("lexicons.initial", parameters.Initial),
		utils.SearchByColumn("lexicons.chinese", parameters.Chinese),
		utils.FilterByArray("lexicons.english", parameters.English, " "),
	)

	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&words).Count(&count).Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&words).Error
	} else {
		err = db.Order(orderQuery).Find(&words).Error
	}

	return
}

// CreateWord creates a new word
func (s *LexiconService) CreateWord(word models.Lexicon) (err error) {
	err = s.PG.Create(&word).Error
	return
}

// GetWord returns word with given id
func (s *LexiconService) GetWord(wordID string) (word models.Lexicon, err error) {
	err = s.PG.Where("id = ?", wordID).Take(&word).Error
	return
}

// UpdateWord updates word with given id
func (s *LexiconService) UpdateWord(wordID string, updatedData map[string]interface{}) (err error) {
	var word models.Lexicon
	err = s.PG.Where("id = ?", wordID).First(&word).Updates(updatedData).Error
	return
}

// DeleteWord soft deletes a word with given id
func (s *LexiconService) DeleteWord(wordID string) (err error) {
	var word models.Lexicon
	err = s.PG.Where("id = ?", wordID).Delete(&word).Error
	return
}
