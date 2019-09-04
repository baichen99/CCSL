package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// UserInterface struct
type UserInterface interface {
	GetUsersList(parameters utils.GetUserListParameters) (users []models.User, count int, err error)
	CreateUser(user models.User) (err error)
	GetUser(key string, value string) (user models.User, err error)
	UpdateUser(userID string, updatedData map[string]interface{}) (err error)
	DeleteUser(userID string) (err error)
}

// UserService implements user interface
type UserService struct {
	PG *gorm.DB
}

// NewUserService returns new user serivce
func NewUserService(pg *gorm.DB) UserInterface {
	return &UserService{
		PG: pg,
	}
}

// GetUsersList returns all users
func (s *UserService) GetUsersList(parameters utils.GetUserListParameters) (users []models.User, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.Scopes(
		utils.FilterByColumn("user_type", parameters.UserType),
		utils.SearchByColumn("username", parameters.SearchName),
	)

	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&users).Count(&count).Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&users).Error
	} else {
		err = db.Order(orderQuery).Find(&users).Error
	}

	return
}

// CreateUser creates a user
func (s *UserService) CreateUser(user models.User) (err error) {
	if user.Password != "" {
		user.Password, err = utils.HashPassword(user.Password)
	}
	err = s.PG.Create(&user).Error
	return
}

// GetUser gets user by id, username or email
func (s *UserService) GetUser(key string, value string) (user models.User, err error) {
	switch key {
	case "id":
		err = s.PG.Where("id = ?", value).Take(&user).Error
	case "username":
		err = s.PG.Where("username = ?", value).Take(&user).Error
	}
	return
}

// UpdateUser updates user model
func (s *UserService) UpdateUser(userID string, updatedData map[string]interface{}) (err error) {
	var user models.User
	err = s.PG.Model(&user).Where("id = ?", userID).Updates(updatedData).Error
	return
}

// DeleteUser soft deletes a user model
func (s *UserService) DeleteUser(userID string) (err error) {
	var user models.User
	err = s.PG.Where("id = ?", userID).Delete(&user).Error
	return
}
