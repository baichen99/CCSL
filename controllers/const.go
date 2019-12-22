package controllers

import (
	"ccsl/models"
	"errors"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

const (
	data      string = "data"
	message   string = "message"
	success   string = "success" // 200 OK
	errinfo   string = "error"
	page      string = "page"
	total     string = "total"
	limit     string = "limit"
	language  string = "language"
	version   string = "version"
	ipAddress string = "ip"
)

var (
	errParams error = errors.New("ParamsError")
	errSQL    error = errors.New("SqlError")
	errAuth   error = errors.New("AuthFailed")
	errRole   error = errors.New("RoleError")
)

// >>> INFO <<<
// ============

type infoUpdateForm struct {
	Data *postgres.Jsonb `json:"data" validate:"required"`
}

// >>> CAROUSEL <<<
// ============

type carouselCreateForm struct {
	TitleZh    string `json:"titleZh" validate:"required"`
	TitleEn    string `json:"titleEn" validate:"required"`
	Image      string `json:"image" validate:"required"`
	State      string `json:"state" validate:"omitempty,oneof=draft published"`
	Importance int    `json:"importance" validate:"omitempty,max=5,min=0" `
}

func (f carouselCreateForm) ConvertToModel() (carousel models.Carousel) {
	carousel = models.Carousel{
		TitleZh:    f.TitleZh,
		TitleEn:    f.TitleEn,
		Image:      f.Image,
		State:      f.State,
		Importance: f.Importance,
	}
	return
}

type carouselUpdateForm struct {
	TitleZh    *string `json:"titleZh" validate:"omitempty"`
	TitleEn    *string `json:"titleEn" validate:"omitempty"`
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
	Degree        string `json:"degree" validate:"omitempty"`
	EmployerZh    string `json:"employerZh" validate:"omitempty"`
	EmployerEn    string `json:"employerEn" validate:"omitempty"`
	DescriptionZh string `json:"descriptionZh" validate:"omitempty"`
	DescriptionEn string `json:"descriptionEn" validate:"omitempty"`
}

func (f memberCreateForm) ConvertToModel() (member models.Member) {
	member = models.Member{
		Profile:       f.Profile,
		Type:          f.Type,
		NameZh:        f.NameZh,
		NameEn:        f.NameEn,
		Degree:        f.Degree,
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
	Degree        *string `json:"degree" validate:"omitempty"`
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
	State    string `json:"state" validate:"required,oneof=active inactive"`
}

func (f userCreateForm) ConvertToModel() (user models.User) {
	user = models.User{
		Name:     f.Name,
		Username: f.Username,
		Password: f.Password,
		UserType: f.UserType,
		State:    f.State,
	}
	return
}

type userUpdateForm struct {
	Avatar   *string `json:"avatar" validate:"omitempty"`
	Name     *string `json:"name" validate:"omitempty"`
	Username *string `json:"username" validate:"omitempty,numeric|email"`
	UserType *string `json:"userType" validate:"omitempty,oneof=admin user learner"`
	State    *string `json:"state" validate:"omitempty,oneof=active inactive"`
}

type userLoginForm struct {
	Username string `json:"username" validate:"required,numeric|email"`
	Password string `json:"password" validate:"required"`
}

type resetPasswordForm struct {
	Email string `json:"email" validate:"required"`
}

// >>> PERFORMER <<<
// ============

type performerCreateForm struct {
	Name     string `json:"name" validate:"required"`
	RegionID int    `json:"regionID" validate:"required,numeric,min=100000"`
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
	RegionID *int    `json:"regionID" validate:"omitempty,min=100000"`
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

// >>> LEXICON <<<
// ============

type lexiconCreateForm struct {
	Initial string `json:"initial" validate:"required"`
	Chinese string `json:"chinese" validate:"required"`
	English string `json:"english" validate:"required"`
	Pos     string `json:"pos" validate:"required"`
}

func (f lexiconCreateForm) ConvertToModel() (word models.Lexicon) {
	word = models.Lexicon{
		Initial: f.Initial,
		Chinese: f.Chinese,
		English: f.English,
		Pos:     f.Pos,
	}
	return
}

type lexiconUpdateForm struct {
	Initial *string `json:"initial" validate:"omitempty"`
	Chinese *string `json:"chinese" validate:"omitempty"`
	English *string `json:"english" validate:"omitempty"`
	Pos     *string `json:"pos" validate:"omitempty"`
}

// >>> LEXICAL_VIDEO <<<
// ============

type lexicalVideoCreateForm struct {
	PerformerID   string   `json:"performerID"  validate:"required,uuid4" `
	LexiconID     string   `json:"lexiconID" validate:"required,uuid4"`
	VideoPath     string   `json:"videoPath" validate:"required"`                // 视频文件路径
	WordFormation string   `json:"wordFormation" validate:"required"`            // 构词方式
	Morpheme      []string `json:"morpheme" validate:"omitempty"`                // 构词词语
	LeftSignsID   []string `json:"leftSignsID" validate:"omitempty,dive,uuid4"`  // 左手手势
	RightSignsID  []string `json:"rightSignsID" validate:"omitempty,dive,uuid4"` // 右手手势
}

func (f lexicalVideoCreateForm) ConvertToModel() (video models.LexicalVideo) {
	performerID, _ := uuid.FromString(f.PerformerID)
	lexiconID, _ := uuid.FromString(f.LexiconID)
	video = models.LexicalVideo{
		PerformerID:   performerID,
		LexiconID:     lexiconID,
		VideoPath:     f.VideoPath,
		WordFormation: f.WordFormation,
		Morpheme:      f.Morpheme,
	}
	return
}

type lexicalVideoUpdateForm struct {
	PerformerID   *string   `json:"performerID" validate:"omitempty,uuid4"`
	LexiconID     *string   `json:"lexiconID" validate:"omitempty,uuid4"`
	VideoPath     *string   `json:"videoPath" validate:"omitempty"`               // 视频文件路径
	WordFormation *string   `json:"wordFormation" validate:"omitempty"`           // 构词方式
	Morpheme      *[]string `json:"morpheme" validate:"omitempty"`                // 构词词语
	LeftSignsID   *[]string `json:"leftSignsID" validate:"omitempty,dive,uuid4"`  // 左手手势
	RightSignsID  *[]string `json:"rightSignsID" validate:"omitempty,dive,uuid4"` // 右手手势
}
