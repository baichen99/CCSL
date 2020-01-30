package services

import (
	"ccsl/models"
	"ccsl/utils"
	"errors"

	"github.com/jinzhu/gorm"
)

// UserInterface struct
type UserInterface interface {
	GetUsersList(parameters utils.GetUserListParameters) (users []models.User, count int, err error)
	CreateUser(user models.User) (err error)
	GetUser(key string, value string) (user models.User, err error)
	UpdateUser(userID string, updatedData map[string]interface{}) (err error)
	DeleteUser(userID string) (err error)
	GetLoginHistoryList(parameters utils.GetUserLoginListParameters) (list []models.LoginHistory, count int, err error)
	CreateLoginHistory(info models.LoginHistory) (err error)
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
	db := s.PG.
		Scopes(
			utils.FilterByColumn("users.state", parameters.State),
			utils.SearchByColumn("users.username", parameters.Username),
			utils.SearchByColumn("users.name", parameters.Name),
			utils.FilterInList("users.roles", parameters.Role),
		)

	// Fetching the total number of rows based on the conditions provided.
	err = db.
		Model(&models.User{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&users).
		Error

	return
}

// GetLoginHistoryList returns users login history
func (s *UserService) GetLoginHistoryList(parameters utils.GetUserLoginListParameters) (list []models.LoginHistory, count int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("login_histories.user_id", parameters.UserID),
			utils.FilterByColumn("login_histories.status", parameters.Status),
			utils.FilterByColumn("login_histories.ip", parameters.IP),
		)
	// Fetching the total number of rows based on the conditions provided.
	err = db.
		Model(&list).
		Count(&count).
		Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.

	err = db.
		Preload("User").
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&list).
		Error

	return
}

// CreateUser creates a user
func (s *UserService) CreateUser(user models.User) (err error) {
	if user.Password != "" {
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			return
		}
	}

	err = s.PG.
		Create(&user).
		Error
	return
}

// CreateLoginHistory creates login histroy
func (s *UserService) CreateLoginHistory(info models.LoginHistory) (err error) {
	err = s.PG.
		Create(&info).
		Error
	return
}

// GetUser gets user by id, username or email
func (s *UserService) GetUser(key string, value string) (user models.User, err error) {
	db := s.PG
	switch key {
	case "id":
		err = db.Where("id = ?", value).Take(&user).Error
	case "username":
		err = db.Where("username = ?", value).Take(&user).Error
	default:
		err = errors.New("Unsupported key")
	}
	return
}

// UpdateUser updates user model
func (s *UserService) UpdateUser(userID string, updatedData map[string]interface{}) (err error) {
	var user models.User
	err = s.PG.
		Where("id = ?", userID).
		Take(&user).
		Updates(updatedData).
		Error
	return
}

// DeleteUser soft deletes a user model
func (s *UserService) DeleteUser(userID string) (err error) {
	var user models.User
	err = s.PG.
		Where("id = ?", userID).
		Delete(&user).
		Error
	return
}
