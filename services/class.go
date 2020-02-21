package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// ClassInterface struct
type ClassInterface interface {
	GetClassesList(parameters utils.GetClassListParameters) (classes []models.Class, count int, err error)
	CreateClass(class models.Class) (err error)
	GetClass(id string) (class models.Class, err error)
	UpdateClass(id string, updatedData map[string]interface{}) (err error)
	DeleteClass(id string) (err error)
	CreateTeacher(id string, uid string) (err error)
	DeleteTeacher(id string, uid string) (err error)
	CreateStudent(id string, uid string) (err error)
	DeleteStudent(id string, uid string) (err error)
}

// ClassService implements class interface
type ClassService struct {
	PG *gorm.DB
}

// NewClassService return a new ClassService
func NewClassService(pg *gorm.DB) ClassInterface {
	return &ClassService{
		PG: pg,
	}
}

// GetClassesList return all classes
func (s *ClassService) GetClassesList(parameters utils.GetClassListParameters) (classes []models.Class, count int, err error) {
	db := s.PG.
		Scopes(
			utils.SearchByColumn("classes.name", parameters.Name),
		)

	err = db.
		Model(&models.Class{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&classes).Error

	return
}

// CreateClass create a class
func (s *ClassService) CreateClass(class models.Class) (err error) {
	err = s.PG.
		Create(&class).
		Error
	return
}

// GetClass get a class by id
func (s *ClassService) GetClass(id string) (class models.Class, err error) {
	err = s.PG.
		Preload("Teachers").
		Preload("Students").
		Preload("Courses").
		Where("id = ?", id).
		Take(&class).
		Error
	return
}

// UpdateClass update a class model
func (s *ClassService) UpdateClass(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Class{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteClass delete a class model
func (s *ClassService) DeleteClass(id string) (err error) {
	var class models.Class
	err = s.PG.
		Where("id = ?", id).
		Delete(&class).
		Error
	return
}

// CreateTeacher associate a class with a teacher
func (s *ClassService) CreateTeacher(id string, uid string) (err error) {
	var class models.Class
	var user models.User
	if user.ID, err = uuid.FromString(uid); err != nil {
		return
	}
	err = s.PG.
		Where("id = ?", id).
		Take(&class).
		Association("Teachers").
		Append(&user).
		Error
	return
}

// DeleteTeacher delete class's association with a teacher
func (s *ClassService) DeleteTeacher(id string, uid string) (err error) {
	var class models.Class
	var user models.User
	if user.ID, err = uuid.FromString(uid); err != nil {
		return
	}
	err = s.PG.
		Where("id = ?", id).
		Take(&class).
		Association("Teachers").
		Delete(&user).
		Error
	return
}

// CreateStudent associate a class with a student
func (s *ClassService) CreateStudent(id string, uid string) (err error) {
	var class models.Class
	var user models.User
	if user.ID, err = uuid.FromString(uid); err != nil {
		return
	}
	err = s.PG.
		Where("id = ?", id).
		Take(&class).
		Association("Students").
		Append(&user).
		Error
	return
}

// DeleteStudent associate a class with a student
func (s *ClassService) DeleteStudent(id string, uid string) (err error) {
	var class models.Class
	var user models.User
	if user.ID, err = uuid.FromString(uid); err != nil {
		return
	}
	err = s.PG.
		Where("id = ?", id).
		Take(&class).
		Association("Students").
		Delete(&user).
		Error
	return
}
