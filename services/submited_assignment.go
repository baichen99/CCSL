package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// SubmittedAssignmentInterface struct
type SubmittedAssignmentInterface interface {
	GetSubmittedAssignmentsList(parameters utils.GetSubmittedAssignmentsListParameters) (submittedAssignment []models.SubmittedAssignments, count int, err error)
	CreateSubmittedAssignment(submittedAssignment models.SubmittedAssignments) (err error)
	GetSubmittedAssignment(id string) (submittedAssignment models.SubmittedAssignments, err error)
	UpdateSubmittedAssignment(id string, updatedData map[string]interface{}) (err error)
	DeleteSubmittedAssignment(id string) (err error)
}

// SubmittedAssignmentService implements submittedAssignment interface
type SubmittedAssignmentService struct {
	PG *gorm.DB
}

// NewSubmittedAssignmentService return a new SubmittedAssignmentService
func NewSubmittedAssignmentService(pg *gorm.DB) SubmittedAssignmentInterface {
	return &SubmittedAssignmentService{
		PG: pg,
	}
}

// GetSubmittedAssignmentsList return all submittedAssignmentes
func (s *SubmittedAssignmentService) GetSubmittedAssignmentsList(parameters utils.GetSubmittedAssignmentsListParameters) (submittedAssignmentes []models.SubmittedAssignments, count int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("submittedAssignments.assignment_id", parameters.AssignmentID),
			utils.FilterByColumn("submittedAssignments.grade_id", parameters.GradeID),
			utils.SearchByColumn("submittedAssignmentes.answer", parameters.Answer),
			utils.SearchByColumn("submittedAssignmentes.comment", parameters.Comment),
		)

	err = db.
		Model(&models.SubmittedAssignments{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Order("name asc").
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&submittedAssignmentes).Error

	return
}

// CreateSubmittedAssignment create a submittedAssignment
func (s *SubmittedAssignmentService) CreateSubmittedAssignment(submittedAssignment models.SubmittedAssignments) (err error) {
	err = s.PG.
		Create(&submittedAssignment).
		Error
	return
}

// GetSubmittedAssignment get a submittedAssignment by id
func (s *SubmittedAssignmentService) GetSubmittedAssignment(id string) (submittedAssignment models.SubmittedAssignments, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&submittedAssignment).
		Error
	return
}

// UpdateSubmittedAssignment update a submittedAssignment model
func (s *SubmittedAssignmentService) UpdateSubmittedAssignment(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.SubmittedAssignments{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteSubmittedAssignment delete a submittedAssignment model
func (s *SubmittedAssignmentService) DeleteSubmittedAssignment(id string) (err error) {
	var submittedAssignment models.SubmittedAssignments
	err = s.PG.
		Where("id = ?", id).
		Delete(&submittedAssignment).
		Error
	return
}
