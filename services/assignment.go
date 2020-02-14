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
	GetSubmittedAssignmentsList(parameters utils.GetSubmittedAssignmentsListParameters) (submittedAssignment []models.SubmittedAssignment, count int, err error)
	CreateSubmittedAssignment(submittedAssignment models.SubmittedAssignment) (err error)
	GetSubmittedAssignment(id string) (submittedAssignment models.SubmittedAssignment, err error)
	UpdateSubmittedAssignment(id string, updatedData map[string]interface{}) (err error)
	DeleteSubmittedAssignment(id string) (err error)
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
	err = s.PG.
		Where("id = ?", id).
		Delete(&models.Assignment{}).
		Error
	return
}

// GetSubmittedAssignmentsList return all submittedAssignmentes
func (s *AssignmentService) GetSubmittedAssignmentsList(parameters utils.GetSubmittedAssignmentsListParameters) (submittedAssignmentes []models.SubmittedAssignment, count int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("submitted_assignments.assignment_id", parameters.AssignmentID),
			utils.FilterByColumn("submitted_assignments.creator_id", parameters.CreatorID),
			utils.FilterByColumn("submitted_assignments.grader_id", parameters.GraderID),
			utils.SearchByColumn("submitted_assignments.answer", parameters.Answer),
			utils.SearchByColumn("submitted_assignments.comment", parameters.Comment),
		)

	err = db.
		Model(&models.SubmittedAssignment{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&submittedAssignmentes).Error

	return
}

// CreateSubmittedAssignment create a submittedAssignment
func (s *AssignmentService) CreateSubmittedAssignment(submittedAssignment models.SubmittedAssignment) (err error) {
	err = s.PG.
		Create(&submittedAssignment).
		Error
	return
}

// GetSubmittedAssignment get a submittedAssignment by id
func (s *AssignmentService) GetSubmittedAssignment(id string) (submittedAssignment models.SubmittedAssignment, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&submittedAssignment).
		Error
	return
}

// UpdateSubmittedAssignment update a submittedAssignment model
func (s *AssignmentService) UpdateSubmittedAssignment(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.SubmittedAssignment{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteSubmittedAssignment delete a submittedAssignment model
func (s *AssignmentService) DeleteSubmittedAssignment(id string) (err error) {
	err = s.PG.
		Where("id = ?", id).
		Delete(&models.SubmittedAssignment{}).
		Error
	return
}
