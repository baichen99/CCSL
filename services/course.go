package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// CourseInterface struct
type CourseInterface interface {
	GetCoursesList(parameters utils.GetCourseListParameters) (course []models.Course, count int, err error)
	CreateCourse(course models.Course) (err error)
	GetCourse(id string) (course models.Course, err error)
	UpdateCourse(id string, updatedData map[string]interface{}) (err error)
	DeleteCourse(id string) (err error)
}

// CourseService implements course interface
type CourseService struct {
	PG *gorm.DB
}

// NewCourseService return a new CourseService
func NewCourseService(pg *gorm.DB) CourseInterface {
	return &CourseService{
		PG: pg,
	}
}

// GetCoursesList return all coursees
func (s *CourseService) GetCoursesList(parameters utils.GetCourseListParameters) (coursees []models.Course, count int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("courses.class_id", parameters.ClassID),
			utils.SearchByColumn("coursees.name", parameters.Name),
			utils.SearchByColumn("coursees.detail", parameters.Content),
		)

	err = db.
		Model(&models.Course{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&coursees).Error

	return
}

// CreateCourse create a course
func (s *CourseService) CreateCourse(course models.Course) (err error) {
	err = s.PG.
		Create(&course).
		Error
	return
}

// GetCourse get a course by id
func (s *CourseService) GetCourse(id string) (course models.Course, err error) {
	err = s.PG.
		Preload("Assignments").
		Where("id = ?", id).
		Take(&course).
		Error
	return
}

// UpdateCourse update a course model
func (s *CourseService) UpdateCourse(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Course{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteCourse delete a course model
func (s *CourseService) DeleteCourse(id string) (err error) {
	var course models.Course
	err = s.PG.
		Where("id = ?", id).
		Delete(&course).
		Error
	return
}
