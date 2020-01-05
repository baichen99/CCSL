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
	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("users.user_type", parameters.UserType),
		utils.FilterByColumn("users.state", parameters.State),
		utils.SearchByColumn("users.username", parameters.Username),
		utils.SearchByColumn("users.name", parameters.Name),
	)

	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&users).Count(&count).Error

	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.
			Order(orderQuery).
			Limit(parameters.Limit).
			Offset(parameters.Limit * (parameters.Page - 1)).
			Find(&users).
			Error
	} else {
		err = db.
			Order(orderQuery).
			Find(&users).
			Error
	}
	return
}

// GetLoginHistoryList returns users login history
func (s *UserService) GetLoginHistoryList(parameters utils.GetUserLoginListParameters) (list []models.LoginHistory, count int, err error) {
	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("login_histories.user_id", parameters.UserID),
		utils.FilterByColumn("login_histories.status", parameters.Status),
		utils.FilterByColumn("login_histories.ip", parameters.IP),
	)
	// Fetching the total number of rows based on the conditions provided.
	err = db.Model(&list).Count(&count).Error
	if err != nil {
		return
	}

	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.
			Set("gorm:auto_preload", true).
			Order(orderQuery).
			Limit(parameters.Limit).
			Offset(parameters.Limit * (parameters.Page - 1)).
			Find(&list).
			Error
	} else {
		err = db.
			Set("gorm:auto_preload", true).
			Order(orderQuery).
			Find(&list).
			Error
	}
	return
}

// CreateUser creates a user
func (s *UserService) CreateUser(user models.User) (err error) {
	if user.Password != "" {
		user.Password, err = utils.HashPassword(user.Password)
	}
	err = s.PG.LogMode(true).Create(&user).Error
	return
}

// CreateLoginHistory creates login histroy
func (s *UserService) CreateLoginHistory(info models.LoginHistory) (err error) {
	err = s.PG.LogMode(true).Create(&info).Error
	return
}

// GetUser gets user by id, username or email
func (s *UserService) GetUser(key string, value string) (user models.User, err error) {
	db := s.PG.LogMode(false)
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
	err = s.PG.LogMode(true).
		Where("id = ?", userID).
		Take(&user).
		Updates(updatedData).
		Error
	return
}

// DeleteUser soft deletes a user model
func (s *UserService) DeleteUser(userID string) (err error) {
	var user models.User
	err = s.PG.LogMode(true).
		Where("id = ?", userID).
		Delete(&user).
		Error
	return
}
