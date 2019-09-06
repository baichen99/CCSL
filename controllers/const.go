package controllers

import (
	"ccsl/models"
	"time"
)

const (
	data     string = "data"
	message  string = "message"
	success  string = "success" // 200 OK
	errinfo  string = "error"
	page     string = "page"
	total    string = "total"
	limit    string = "limit"
	language string = "language"
)

// >>> CAROUSEL <<<
// ============

type carouselCreateForm struct {
	Title  string `json:"title" validate:"required"`
	Image  string `json:"image" validate:"required"`
	Status string `json:"status" validate:"omitempty,oneof=draft published"`
}

func (f carouselCreateForm) ConvertToModel() (carousel models.Carousel) {
	carousel = models.Carousel{
		Title:  f.Title,
		Image:  f.Image,
		Status: f.Status,
	}
	return
}

type carouselUpdateForm struct {
	Title  *string `json:"title" validate:"omitempty"`
	Image  *string `json:"image" validate:"omitempty"`
	Status *string `json:"status" validate:"omitempty,oneof=draft published"`
}

// >>> NEWS <<<
// ============

type newsCreateForm struct {
	Column   string    `json:"column" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	Title    string    `json:"title" validate:"required"`
	Type     string    `json:"type" validate:"required,oneof=link document"`     // Type can be 'link' or 'document'
	Text     string    `json:"text" validate:"required"`                         // Text
	Language string    `json:"language"  validate:"omitempty,oneof=zh-CN en-US"` // Can be zh-CN or en-US
	Status   string    `json:"status" validate:"omitempty,oneof=draft published"`
}

func (f newsCreateForm) ConvertToModel() (news models.News) {
	news = models.News{
		Column:   f.Column,
		Date:     f.Date,
		Title:    f.Title,
		Type:     f.Type,
		Text:     f.Text,
		Language: f.Language,
		Status:   f.Status,
	}
	return
}

type newsUpdateForm struct {
	Column   *string    `json:"column"  validate:"omitempty"`
	Date     *time.Time `json:"date"  validate:"omitempty"`
	Title    *string    `json:"title"  validate:"omitempty"`
	Type     *string    `json:"type"  validate:"omitempty,oneof=link document"`
	Text     *string    `json:"text"  validate:"omitempty"`
	Language *string    `json:"language" validate:"omitempty,oneof=zh-CN en-US"`
	Status   *string    `json:"status" validate:"omitempty,oneof=draft published"`
}

// >>> USER <<<
// ============
type userCreateForm struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,numeric|email"`
	Password string `json:"password" validate:"omitempty"`
	UserType string `json:"userType" validate:"required,oneof=admin user learner"`
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
	Name     *string `json:"name" validate:"omitempty"`
	Username *string `json:"username" validate:"omitempty,numeric|email"`
	// Password *string `json:"password" validate:"omitempty,min=6"`
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
	Initial *string `json:"initial" validate:"omitempty"`
	Chinese *string `json:"chinese" validate:"omitempty"`
	English *string `json:"english" validate:"omitempty"`
	Type    *string `json:"type" validate:"omitempty"`
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
	Name   *string `json:"name" validate:"omitempty"`
	Region *string `json:"region" validate:"omitempty"`
	Gender *string `json:"gender" validate:"omitempty"`
}

// >>> VIDEO <<<
// ============

type videoCreateForm struct {
	ConstructType  string `json:"constructType" validate:"required"`  // 构词方式
	ConstructWords string `json:"constructWords" validate:"required"` // 构词词语
	LeftSign       string `json:"leftSign" validate:"required"`       // 左手手势
	RightSign      string `json:"rightSign" validate:"required"`      // 右手手势
	VideoPath      string `json:"videoPath" validate:"required"`      // 视频文件路径
}

func (f videoCreateForm) ConvertToModel() (video models.Video) {
	video = models.Video{
		ConstructType:  f.ConstructType,
		ConstructWords: f.ConstructWords,
		LeftSign:       f.LeftSign,
		RightSign:      f.RightSign,
		VideoPath:      f.VideoPath,
	}
	return
}

type videoUpdateForm struct {
	ConstructType  *string `json:"constructType" validate:"omitempty"`  // 构词方式
	ConstructWords *string `json:"constructWords" validate:"omitempty"` // 构词词语
	LeftSign       *string `json:"leftSign" validate:"omitempty"`       // 左手手势
	RightSign      *string `json:"rightSign" validate:"omitempty"`      // 右手手势
	VideoPath      *string `json:"videoPath" validate:"omitempty"`      // 视频文件路径
}
