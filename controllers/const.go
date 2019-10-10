package controllers

import (
	"ccsl/models"
	"time"

	uuid "github.com/satori/go.uuid"
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
	version  string = "version"
)

// >>> CAROUSEL <<<
// ============

type carouselCreateForm struct {
	Title      string `json:"title" validate:"required"`
	Image      string `json:"image" validate:"required"`
	State      string `json:"state" validate:"omitempty,oneof=draft published"`
	Importance int    `json:"importance" validate:"omitempty,max=5,min=0" `
}

func (f carouselCreateForm) ConvertToModel() (carousel models.Carousel) {
	carousel = models.Carousel{
		Title:      f.Title,
		Image:      f.Image,
		State:      f.State,
		Importance: f.Importance,
	}
	return
}

type carouselUpdateForm struct {
	Title      *string `json:"title" validate:"omitempty"`
	Image      *string `json:"image" validate:"omitempty"`
	State      *string `json:"state" validate:"omitempty,oneof=draft published"`
	Importance *int    `json:"importance" validate:"omitempty,max=5,min=0" `
}

// >>> NEWS <<<
// ============

type newsCreateForm struct {
	Column     string    `json:"column" validate:"required"`
	Date       time.Time `json:"date" validate:"required"`
	Title      string    `json:"title" validate:"required"`
	Type       string    `json:"type" validate:"required,oneof=link document"`     // Type can be 'link' or 'document'
	Text       string    `json:"text" validate:"required"`                         // Text
	Language   string    `json:"language"  validate:"omitempty,oneof=zh-CN en-US"` // Can be zh-CN or en-US
	Importance int       `json:"importance" validate:"omitempty,max=5,min=0" `
	State      string    `json:"state" validate:"omitempty,oneof=draft published"`
}

func (f newsCreateForm) ConvertToModel() (news models.News) {
	news = models.News{
		Column:     f.Column,
		Date:       f.Date,
		Title:      f.Title,
		Type:       f.Type,
		Text:       f.Text,
		Language:   f.Language,
		Importance: f.Importance,
		State:      f.State,
	}
	return
}

type newsUpdateForm struct {
	Column     *string    `json:"column"  validate:"omitempty"`
	Date       *time.Time `json:"date"  validate:"omitempty"` // RFC3339 - example: 2000-12-30T00:00:00Z
	Title      *string    `json:"title"  validate:"omitempty"`
	Type       *string    `json:"type"  validate:"omitempty,oneof=link document"`
	Text       *string    `json:"text"  validate:"omitempty"`
	Language   *string    `json:"language" validate:"omitempty,oneof=zh-CN en-US"`
	Importance *int       `json:"importance" validate:"omitempty,max=5,min=0" `
	State      *string    `json:"state" validate:"omitempty,oneof=draft published"`
}

// >> MEMBER <<
// ============
type memberCreateForm struct {
	Profile       string `json:"profile" validate:"required"`
	Type          string `json:"type" validate:"required"`
	NameZh        string `json:"nameZh" validate:"required"`
	NameEn        string `json:"nameEn" validate:"required"`
	PositionZh    string `json:"positionZh" validate:"required"`
	PositionEn    string `json:"positionEn" validate:"required"`
	EmployerZh    string `json:"employerZh" validate:"required"`
	EmployerEn    string `json:"employerEn" validate:"required"`
	DescriptionZh string `json:"descriptionZh" validate:"required"`
	DescriptionEn string `json:"descriptionEn" validate:"required"`
}

func (f memberCreateForm) ConvertToModel() (member models.Member) {
	member = models.Member{
		Profile:       f.Profile,
		Type:          f.Type,
		NameZh:        f.NameZh,
		NameEn:        f.NameEn,
		PositionZh:    f.PositionZh,
		PositionEn:    f.PositionEn,
		EmployerZh:    f.EmployerZh,
		EmployerEn:    f.EmployerEn,
		DescriptionZh: f.DescriptionZh,
		DescriptionEn: f.DescriptionEn,
	}
	return
}

type memberUpdateForm struct {
	Profile       *string `json:"profile" validate:"omitempty"`
	Type          *string `json:"type" validate:"omitempty"`
	Text          *string `json:"text" validate:"omitempty"`
	NameZh        *string `json:"nameZh" validate:"omitempty"`
	NameEn        *string `json:"nameEn" validate:"omitempty"`
	PositionZh    *string `json:"positionZh" validate:"omitempty"`
	PositionEn    *string `json:"positionEn" validate:"omitempty"`
	EmployerZh    *string `json:"employerZh" validate:"omitempty"`
	EmployerEn    *string `json:"employerEn" validate:"omitempty"`
	DescriptionZh *string `json:"descriptionZh" validate:"omitempty"`
	DescriptionEn *string `json:"descriptionEn" validate:"omitempty"`
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

// >>> PERFORMER <<<
// ============

type performerCreateForm struct {
	Name     string `json:"name" validate:"required"`
	RegionID int    `json:"regionID" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
}

func (f performerCreateForm) ConvertToModel() (performer models.Performer) {
	performer = models.Performer{
		Name:     f.Name,
		RegionID: f.RegionID,
		Gender:   f.Gender,
	}
	return
}

type performerUpdateForm struct {
	Name     *string `json:"name" validate:"omitempty"`
	RegionID *int    `json:"regionID" validate:"omitempty"`
	Gender   *string `json:"gender" validate:"omitempty"`
}

// >>> SIGN <<<
// ============

type signCreateForm struct {
	Name  string `json:"name" validate:"required"`
	Image string `json:"image" validate:"required"`
}

func (f signCreateForm) ConvertToModel() (user models.Sign) {
	user = models.Sign{
		Name:  f.Name,
		Image: f.Image,
	}
	return
}

type signUpdateForm struct {
	Name  *string `json:"name" validate:"omitempty"`
	Image *string `json:"image" validate:"omitempty"`
}

// >>> LEXICAL_WORD <<<
// ============

type lexicalWordCreateForm struct {
	Initial string `json:"initial" validate:"required"`
	Chinese string `json:"chinese" validate:"required"`
	English string `json:"english" validate:"required"`
	Pos     string `json:"pos" validate:"required"`
}

func (f lexicalWordCreateForm) ConvertToModel() (word models.LexicalWord) {
	word = models.LexicalWord{
		Initial: f.Initial,
		Chinese: f.Chinese,
		English: f.English,
		Pos:     f.Pos,
	}
	return
}

type lexicalWordUpdateForm struct {
	Initial *string `json:"initial" validate:"omitempty"`
	Chinese *string `json:"chinese" validate:"omitempty"`
	English *string `json:"english" validate:"omitempty"`
	Pos     *string `json:"pos" validate:"omitempty"`
}

// >>> LEXICAL_VIDEO <<<
// ============

type lexicalVideoCreateForm struct {
	PerformerID    string   `json:"performerID"  validate:"required,uuid4" `
	LexicalWordID  string   `json:"lexicalWordID" validate:"required,uuid4"`
	VideoPath      string   `json:"videoPath" validate:"required"`      // 视频文件路径
	ConstructType  string   `json:"constructType" validate:"required"`  // 构词方式
	ConstructWords []string `json:"constructWords" validate:"required"` // 构词词语
	LeftSigns      []string `json:"leftSigns" validate:"omitempty"`     // 左手手势
	RightSigns     []string `json:"rightSigns" validate:"omitempty"`    // 右手手势
}

func (f lexicalVideoCreateForm) ConvertToModel() (video models.LexicalVideo) {
	performerID, _ := uuid.FromString(f.PerformerID)
	lexicalWordID, _ := uuid.FromString(f.LexicalWordID)
	video = models.LexicalVideo{
		PerformerID:    performerID,
		LexicalWordID:  lexicalWordID,
		VideoPath:      f.VideoPath,
		ConstructType:  f.ConstructType,
		ConstructWords: f.ConstructWords,
	}
	return
}

type lexicalVideoUpdateForm struct {
	PerformerID    *string   `json:"performerID" validate:"omitempty,uuid4"`
	LexicalWordID  *string   `json:"lexicalWordID" validate:"omitempty,uuid4"`
	VideoPath      *string   `json:"videoPath" validate:"omitempty"`      // 视频文件路径
	ConstructType  *string   `json:"constructType" validate:"omitempty"`  // 构词方式
	ConstructWords *[]string `json:"constructWords" validate:"omitempty"` // 构词词语
	LeftSigns      *[]string `json:"leftSign" validate:"omitempty"`       // 左手手势
	RightSigns     *[]string `json:"rightSign" validate:"omitempty"`      // 右手手势
}
