package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// LexicalWordInterface struct
type LexicalWordInterface interface {
	GetWordsList(parameters utils.GetWordListParameters) (words []models.LexicalWord, count int, err error)
	CreateWord(word models.LexicalWord) (err error)
	GetWord(wordID string) (word models.LexicalWord, err error)
	UpdateWord(wordID string, updatedData map[string]interface{}) (err error)
	DeleteWord(wordID string) (err error)
}

// LexicalWordService implements word interface
type LexicalWordService struct {
	PG *gorm.DB
}

// NewLexicalWordService returns new word serivce
func NewLexicalWordService(pg *gorm.DB) LexicalWordInterface {
	return &LexicalWordService{
		PG: pg,
	}
}

// GetWordsList returns words list
func (s *LexicalWordService) GetWordsList(parameters utils.GetWordListParameters) (words []models.LexicalWord, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("lexical_words.pos", parameters.Pos),
		utils.FilterByColumn("lexical_words.initial", parameters.Initial),
		utils.SearchByColumn("lexical_words.chinese", parameters.Chinese),
		utils.FilterByArray("lexical_words.english", parameters.English, " "),
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
func (s *LexicalWordService) CreateWord(word models.LexicalWord) (err error) {
	err = s.PG.LogMode(true).Create(&word).Error
	return
}

// GetWord returns word with given id
func (s *LexicalWordService) GetWord(wordID string) (word models.LexicalWord, err error) {
	err = s.PG.LogMode(false).Where("id = ?", wordID).Take(&word).Error
	return
}

// UpdateWord updates word with given id
func (s *LexicalWordService) UpdateWord(wordID string, updatedData map[string]interface{}) (err error) {
	var word models.LexicalWord
	err = s.PG.LogMode(true).Where("id = ?", wordID).First(&word).Updates(updatedData).Error
	return
}

// DeleteWord soft deletes a word with given id
func (s *LexicalWordService) DeleteWord(wordID string) (err error) {
	var word models.LexicalWord
	err = s.PG.LogMode(true).Where("id = ?", wordID).Delete(&word).Error
	return
}
