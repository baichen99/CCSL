package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// AssignmentInterface struct
type AssignmentInterface interface {
	GetAssignmentsList(parameters utils.GetAssignmentListParameters) (assignments []models.Assignment, count int, err error)
	CreateAssignment(assignment models.Assignment) (err error)
	GetAssignment(id string) (assignment models.Assignment, err error)
	UpdateAssignment(id string, updatedData map[string]interface{}) (err error)
	DeleteAssignment(id string) (err error)
}

// AssignmentService implements assignment interface
type AssignmentService struct {
	PG *gorm.DB
}

// NewAssignmentService return a new AssignmentService
func NewAssignmentService(pg *gorm.DB) AssignmentInterface {
	return &AssignmentService{
		PG: pg,
	}
}

// GetAssignmentsList return all assignmentes
func (s *AssignmentService) GetAssignmentsList(parameters utils.GetAssignmentListParameters) (assignments []models.Assignment, count int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("assignments.course_id", parameters.CourseID),
			utils.FilterByColumn("assignmentes.type", parameters.Type),
			utils.SearchByColumn("assignmentes.title", parameters.Title),
			utils.SearchByColumn("assignmentes.content", parameters.Content),
		)

	err = db.
		Model(&models.Assignment{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&assignments).Error

	return
}

// CreateAssignment create a assignment
func (s *AssignmentService) CreateAssignment(assignment models.Assignment) (err error) {
	err = s.PG.
		Create(&assignment).
		Error
	return
}

// GetAssignment get a assignment by id
func (s *AssignmentService) GetAssignment(id string) (assignment models.Assignment, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&assignment).
		Error
	return
}

// UpdateAssignment update a assignment model
func (s *AssignmentService) UpdateAssignment(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Assignment{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteAssignment delete a assignment model
func (s *AssignmentService) DeleteAssignment(id string) (err error) {
	var assignment models.Assignment
	err = s.PG.
		Where("id = ?", id).
		Delete(&assignment).
		Error
	return
}
