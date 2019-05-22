package controllers

import (
	"ccsl/models"
)

const (
	data      string = "data"
	message   string = "message"
	success   string = "success" // 200 OK
	errinfo   string = "error"
	pageKey   string = "page"
	totalKey  string = "total"
	limitKey  string = "limit"
	paramsKey string = "params"
)

// >>> USER <<<
// ============
type userCreateForm struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,numeric|email"`
	Password string `json:"password" validate:"omitempty"`
	UserType string `json:"userType" validate:"required,oneof=admin user"`
}

func (f userCreateForm) ConvertToModel() (user models.User) {
	user = models.User{
		Name:     f.Name,
		Username: f.Username,
		Password: f.Password,
		UserType: f.UserType,
	}
	return
}

type userUpdateForm struct {
	Name     string `json:"name" validate:"omitempty"`
	Username string `json:"username" validate:"omitempty,numeric|email"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

func (f userUpdateForm) ConvertToModel() (user models.User) {
	user = models.User{
		Name:     f.Name,
		Username: f.Username,
		Password: f.Password,
	}
	return
}

type userLoginForm struct {
	Username string `json:"username" validate:"required,numeric|email"`
	Password string `json:"password" validate:"required"`
}

// >>> WORD <<<
// ============

type wordCreateForm struct {
	Initial string `json:"initial" validate:"required"`
	Chinese string `json:"chinese" validate:"required"`
	English string `json:"english" validate:"required"`
	Type    string `json:"type" validate:"required"`
}

func (f wordCreateForm) ConvertToModel() (word models.Word) {
	word = models.Word{
		Initial: f.Initial,
		Chinese: f.Chinese,
		English: f.English,
		Type:    f.Type,
	}
	return
}

type wordUpdateForm struct {
	Initial string `json:"initial" validate:"omitempty"`
	Chinese string `json:"chinese" validate:"omitempty"`
	English string `json:"english" validate:"omitempty"`
	Type    string `json:"type" validate:"omitempty"`
}

func (f wordUpdateForm) ConvertToModel() (word models.Word) {
	word = models.Word{
		Initial: f.Initial,
		Chinese: f.Chinese,
		English: f.English,
		Type:    f.Type,
	}
	return
}

// >>> PERFORMER <<<
// ============

type performerCreateForm struct {
	Name   string `json:"name" validate:"required"`
	Region string `json:"region" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}

func (f performerCreateForm) ConvertToModel() (performer models.Performer) {
	performer = models.Performer{
		Name:   f.Name,
		Region: f.Region,
		Gender: f.Gender,
	}
	return
}

type performerUpdateForm struct {
	Name   string `json:"name" validate:"omitempty"`
	Region string `json:"region" validate:"omitempty"`
	Gender string `json:"gender" validate:"omitempty"`
}

func (f performerUpdateForm) ConvertToModel() (performer models.Performer) {
	performer = models.Performer{
		Name:   f.Name,
		Region: f.Region,
		Gender: f.Gender,
	}
	return
}
