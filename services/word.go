package services

import (
	"ccsl/models"
	"ccsl/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// WordInterface struct
type WordInterface interface {
	GetWordsList(parameters utils.GetWordListParameters) (words []models.Word, count int, err error)
	CreateWord(word models.Word) (err error)
	GetWord(wordID string) (word models.Word, err error)
	UpdateWord(wordID string, updatedData models.Word) (err error)
	DeleteWord(wordID string) (err error)
}

// WordService implements word interface
type WordService struct {
	PG *gorm.DB
}

// NewWordService returns new word serivce
func NewWordService(pg *gorm.DB) WordInterface {
	return &WordService{
		PG: pg,
	}
}

// GetWordsList returns words list
func (s *WordService) GetWordsList(parameters utils.GetWordListParameters) (words []models.Word, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.Scopes(
		utils.FilterByColumn("type", parameters.Type),
		utils.FilterByColumn("initial", parameters.Initial),
		utils.SearchByColumn("chinese", parameters.Chinese),
		utils.SearchByColumn("english", parameters.English),
	)

	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&words).Count(&count).Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := fmt.Sprintf("%s %s", parameters.OrderBy, parameters.Order)
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&words).Error
	} else {
		err = db.Order(orderQuery).Find(&words).Error
	}

	return
}

// CreateWord creates a new word
func (s *WordService) CreateWord(word models.Word) (err error) {
	err = s.PG.Create(&word).Error
	return
}

// GetWord returns word with given id
func (s *WordService) GetWord(wordID string) (word models.Word, err error) {
	err = s.PG.Where("id = ?", wordID).Take(&word).Error
	return
}

// UpdateWord updates word with given id
func (s *WordService) UpdateWord(wordID string, updatedData models.Word) (err error) {
	var word models.Word
	err = s.PG.Where("id = ?", wordID).Take(&word).Model(&word).Updates(updatedData).Error
	return
}

// DeleteWord soft deletes a word with given id
func (s *WordService) DeleteWord(wordID string) (err error) {
	var word models.Word
	err = s.PG.Where("id = ?", wordID).Delete(&word).Error
	return
}
