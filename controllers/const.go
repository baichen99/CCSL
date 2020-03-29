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
	unread    string = "unread"
	language  string = "language"
	version   string = "version"
	ipAddress string = "ip"
)

var (
	errParams   error = errors.New("ParamsError")
	errSQL      error = errors.New("SqlError")
	errAuth     error = errors.New("AuthFailed")
	errRole     error = errors.New("RoleError")
	errOutdated error = errors.New("Outdated")
)

// ErrorResponse Response for error request
type ErrorResponse struct {
	Message string `json:"message" example:"error title"`
	ErrInfo string `json:"error" example:"error message"`
}

// SuccessResponse Response for successful operation
type SuccessResponse struct {
	Message string `json:"message" example:"success"`
}

// GetListResponse Response for get list request
type GetListResponse struct {
	Message string `json:"message" example:"success"`
	Page    int    `json:"page" example:"1"`
	Limit   int    `json:"limit" example:"10"`
	Total   int    `json:"total" example:"100"`
}

// >>> INFO <<<
// ============

type InfoUpdateForm struct {
	Data *postgres.Jsonb `json:"data" validate:"required"`
}

// >>> CAROUSEL <<<
// ============

type CarouselCreateForm struct {
	TitleZh    string `json:"titleZh" validate:"required"`
	TitleEn    string `json:"titleEn" validate:"required"`
	Image      string `json:"image" validate:"required"`
	State      string `json:"state" validate:"omitempty,oneof=draft published"`
	Importance int    `json:"importance" validate:"omitempty,max=5,min=0" `
}

func (f CarouselCreateForm) ConvertToModel() (carousel models.Carousel) {
	carousel = models.Carousel{
		TitleZh:    f.TitleZh,
		TitleEn:    f.TitleEn,
		Image:      f.Image,
		State:      f.State,
		Importance: f.Importance,
	}
	return
}

type CarouselUpdateForm struct {
	TitleZh    *string `json:"titleZh" validate:"omitempty"`
	TitleEn    *string `json:"titleEn" validate:"omitempty"`
	Image      *string `json:"image" validate:"omitempty"`
	State      *string `json:"state" validate:"omitempty,oneof=draft published"`
	Importance *int    `json:"importance" validate:"omitempty,max=5,min=0" `
}

// >>> NEWS <<<
// ============

type NewsCreateForm struct {
	Column     string    `json:"column" validate:"required"`
	Date       time.Time `json:"date" validate:"required"`
	Title      string    `json:"title" validate:"required"`
	Type       string    `json:"type" validate:"required,oneof=link document"`     // Type can be 'link' or 'document'
	Text       string    `json:"text" validate:"required"`                         // Text
	Language   string    `json:"language"  validate:"omitempty,oneof=zh-CN en-US"` // Can be zh-CN or en-US
	Importance int       `json:"importance" validate:"omitempty,max=5,min=0" `
	State      string    `json:"state" validate:"omitempty,oneof=draft published"`
}

func (f NewsCreateForm) ConvertToModel() (news models.News) {
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

type NewsUpdateForm struct {
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
type MemberCreateForm struct {
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

func (f MemberCreateForm) ConvertToModel() (member models.Member) {
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

type MemberUpdateForm struct {
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

// UserCreateForm user create form
type UserCreateForm struct {
	Avatar   string   `json:"avatar" validate:"omitempty" example:"https://ccsl.shu.edu.cn/public/assets/default.png"`
	Name     string   `json:"name" validate:"required" example:"Adrian Duan"`
	Username string   `json:"username" validate:"required,numeric|email" example:"adrianduan@icloud.com"`
	Password string   `json:"password" validate:"omitempty" example:"p@ssw0rd"`
	State    string   `json:"state" validate:"required,oneof=active inactive" example:"active" enums:"active,inactive"`
	Roles    []string `json:"roles" validate:"required,dive,oneof=super admin dbuser student teacher" swaggertype:"array,string" example:"admin" enums:"super,admin,dbuser,student,teacher"`
}

// ConvertToModel convert form to model
func (f UserCreateForm) ConvertToModel() (user models.User) {
	user = models.User{
		Avatar:   f.Avatar,
		Name:     f.Name,
		Username: f.Username,
		Password: f.Password,
		State:    f.State,
		Roles:    f.Roles,
	}
	return
}

// UserUpdateForm user update form
type UserUpdateForm struct {
	Avatar   *string   `json:"avatar" validate:"omitempty"  example:"https://ccsl.shu.edu.cn/public/assets/default.png"`
	Name     *string   `json:"name" validate:"omitempty" example:"Adrian Duan"`
	Username *string   `json:"username" validate:"omitempty,numeric|email" example:"adrianduan@icloud.com"`
	State    *string   `json:"state" validate:"omitempty,oneof=active inactive" example:"active" enums:"active,inactive"`
	Roles    *[]string `json:"roles"  validate:"omitempty,dive,oneof=super admin dbuser student teacher" swaggertype:"array,string" example:"admin" enums:"super,admin,dbuser,student,teacher"`
}

type UserLoginForm struct {
	Username string `json:"username" validate:"required,numeric|email"`
	Password string `json:"password" validate:"required"`
}

type ResetPasswordForm struct {
	Email string `json:"email" validate:"required"`
}

// GetUsersListResponse Response for GetUsersList
type GetUsersListResponse struct {
	GetListResponse
	Data []models.User `json:"data"`
}

type GetUserResponse struct {
	Message string      `json:"message" example:"success"`
	Data    models.User `json:"data"`
}

// >>> PERFORMER <<<
// ============

type PerformerCreateForm struct {
	Name     string `json:"name" validate:"required"`
	RegionID int    `json:"regionID" validate:"required,numeric,min=100000"`
	Gender   string `json:"gender" validate:"required"`
}

func (f PerformerCreateForm) ConvertToModel() (performer models.Performer) {
	performer = models.Performer{
		Name:     f.Name,
		RegionID: f.RegionID,
		Gender:   f.Gender,
	}
	return
}

type PerformerUpdateForm struct {
	Name     *string `json:"name" validate:"omitempty"`
	RegionID *int    `json:"regionID" validate:"omitempty,min=100000"`
	Gender   *string `json:"gender" validate:"omitempty"`
}

// >>> HANDSHAPE <<<
// ============

type HandshapeCreateForm struct {
	Name  string `json:"name" validate:"required"`
	Image string `json:"image" validate:"required"`
	Glyph string `json:"glyph" validate:"required"`
}

func (f HandshapeCreateForm) ConvertToModel() (user models.Handshape) {
	user = models.Handshape{
		Name:  f.Name,
		Image: f.Image,
		Glyph: f.Glyph,
	}
	return
}

type HandshapeUpdateForm struct {
	Name  *string `json:"name" validate:"omitempty"`
	Image *string `json:"image" validate:"omitempty"`
	Glyph *string `json:"glyph" validate:"omitempty"`
}

// >>> LEXICON <<<
// ============

type LexiconCreateForm struct {
	Initial string   `json:"initial" validate:"required"`
	Chinese string   `json:"chinese" validate:"required"`
	English string   `json:"english" validate:"required"`
	Pos     []string `json:"pos" validate:"required"`
}

func (f LexiconCreateForm) ConvertToModel() (word models.Lexicon) {
	word = models.Lexicon{
		Initial: f.Initial,
		Chinese: f.Chinese,
		English: f.English,
		Pos:     f.Pos,
	}
	return
}

type LexiconUpdateForm struct {
	Initial *string   `json:"initial" validate:"omitempty"`
	Chinese *string   `json:"chinese" validate:"omitempty"`
	English *string   `json:"english" validate:"omitempty"`
	Pos     *[]string `json:"pos" validate:"omitempty"`
}

// >>> LEXICAL_VIDEO <<<
// ============

type LexicalVideoCreateForm struct {
	PerformerID       string   `json:"performerID"  validate:"required,uuid4" `
	LexiconID         string   `json:"lexiconID" validate:"required,uuid4"`
	VideoPath         string   `json:"videoPath" validate:"required"`                     // 视频文件路径
	WordFormation     string   `json:"wordFormation" validate:"required"`                 // 构词方式
	Morpheme          []string `json:"morpheme" validate:"omitempty"`                     // 构词词语
	LeftHandshapesID  []string `json:"leftHandshapesID" validate:"omitempty,dive,uuid4"`  // 左手手势
	RightHandshapesID []string `json:"rightHandshapesID" validate:"omitempty,dive,uuid4"` // 右手手势
}

func (f LexicalVideoCreateForm) ConvertToModel() (video models.LexicalVideo) {
	performerID, _ := uuid.FromString(f.PerformerID)
	lexiconID, _ := uuid.FromString(f.LexiconID)
	video = models.LexicalVideo{
		PerformerID:       performerID,
		LexiconID:         lexiconID,
		VideoPath:         f.VideoPath,
		WordFormation:     f.WordFormation,
		Morpheme:          f.Morpheme,
		LeftHandshapesID:  f.LeftHandshapesID,
		RightHandshapesID: f.RightHandshapesID,
	}
	return
}

type LexicalVideoUpdateForm struct {
	PerformerID       *string   `json:"performerID" validate:"omitempty,uuid4"`
	LexiconID         *string   `json:"lexiconID" validate:"omitempty,uuid4"`
	VideoPath         *string   `json:"videoPath" validate:"omitempty"`                    // 视频文件路径
	WordFormation     *string   `json:"wordFormation" validate:"omitempty"`                // 构词方式
	Morpheme          *[]string `json:"morpheme" validate:"omitempty"`                     // 构词词语
	LeftHandshapesID  *[]string `json:"leftHandshapesID" validate:"omitempty,dive,uuid4"`  // 左手手势
	RightHandshapesID *[]string `json:"rightHandshapesID" validate:"omitempty,dive,uuid4"` // 右手手势
}

// GetLexicalVideosListResponse Response for GetLexicalVideosList
type GetLexicalVideosListResponse struct {
	GetListResponse
	Data []models.LexicalVideo `json:"data"`
}

type GetLexicalVideoResponse struct {
	Message string              `json:"message" example:"success"`
	Data    models.LexicalVideo `json:"data"`
}

// >>> CLASS <<<
// ============

type ClassCreateForm struct {
	Name      string `json:"name" validate:"required"`
	Details   string `json:"details" validate:"required"`
	Resources string `json:"resources" validate:"required"`
}

func (f ClassCreateForm) ConvertToModel() (class models.Class) {
	class = models.Class{
		Name:      f.Name,
		Details:   f.Details,
		Resources: f.Resources,
	}
	return
}

type ClassUpdateForm struct {
	Name      *string `json:"name"  validate:"omitempty"`
	Details   *string `json:"details"  validate:"omitempty"`
	Resources *string `json:"resources"  validate:"omitempty"`
}

type ClassStudentCreateForm struct {
	Username string `json:"username" validate:"required,numeric|email"`
	Name     string `json:"name" validate:"required"`
}

// GetClassListResponse Response for GetClassList
type GetClassListResponse struct {
	GetListResponse
	Data []models.Class `json:"data"`
}

// GetClassResponse Response for get class
type GetClassResponse struct {
	Message string       `json:"message" example:"success"`
	Data    models.Class `json:"data"`
}

// >>> COURSE <<<
// ============
type CourseCreateForm struct {
	ClassID string `json:"classID" validate:"required,uuid4"`
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (f CourseCreateForm) ConvertToModel() (course models.Course) {
	classID, _ := uuid.FromString(f.ClassID)
	course = models.Course{
		ClassID: classID,
		Name:    f.Name,
		Content: f.Content,
	}
	return
}

type CourseUpdateForm struct {
	ClassID *string `json:"classID" validate:"omitempty"`
	Name    *string `json:"name"  validate:"omitempty"`
	Content *string `json:"content"  validate:"omitempty"`
}

// GetCourseListResponse Response for GetCourseList
type GetCourseListResponse struct {
	GetListResponse
	Data []models.Course `json:"data"`
}

// GetCourseResponse Response for get course
type GetCourseResponse struct {
	Message string        `json:"message" example:"success"`
	Data    models.Course `json:"data"`
}

// >>> ASSIGNMENT <<<
// ============
type AssignmentCreateForm struct {
	CourseID string    `json:"courseID" validate:"required,uuid4"`
	Title    string    `json:"title" validate:"required"`
	Type     string    `json:"type" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	Deadline time.Time `json:"deadline" validate:"required"`
}

func (f AssignmentCreateForm) ConvertToModel() (assignment models.Assignment) {
	courseID, _ := uuid.FromString(f.CourseID)
	assignment = models.Assignment{
		CourseID: courseID,
		Title:    f.Title,
		Type:     f.Type,
		Content:  f.Content,
		Deadline: &f.Deadline,
	}
	return
}

type AssignmentUpdateForm struct {
	CourseID *string    `json:"courseID" validate:"omitempty,uuid4"`
	Title    *string    `json:"title"  validate:"omitempty"`
	Type     *string    `json:"type"  validate:"omitempty"`
	Content  *string    `json:"content"  validate:"omitempty"`
	Deadline *time.Time `json:"deadline" validate:"omitempty"`
}

// >>> POST <<<
// ============
type PostCreateForm struct {
	//CreatorID string `json:"creatorID" validate:"required,uuid4"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (f PostCreateForm) ConvertToModel() (post models.Post) {
	//creatorID, _ := uuid.FromString(f.CreatorID)
	post = models.Post{
		Title:   f.Title,
		Content: f.Content,
		//CreatorID: creatorID,
	}
	return
}

type PostUpdateForm struct {
	// json:"PostID" is not given by request's json, it's given by URL
	PostID  *string `json:"postID" validate:"omitempty,uuid4"`
	Title   *string `json:"title" validate:"omitempty"`
	Content *string `json:"content" validate:"omitempty"`
}

// >>> REPLY <<<
// ============
type ReplyCreateForm struct {
	//ReplyID string `json:"replyID" validate:"required"`
	PostID  string `json:"postID" validate:"required,uuid4"`
	Content string `json:"content" validate:"required"`
}

func (f ReplyCreateForm) ConvertToModel() (reply models.Reply) {
	postID, _ := uuid.FromString(f.PostID)
	reply = models.Reply{
		PostID:  postID,
		Content: f.Content,
	}
	return
}

type ReplyUpdateForm struct {
	// json:"replyID" is not given by request's json, it's given by URL
	ReplyID *string `json:"replyID" validate:"omitempty,uuid4"`
	Content *string `json:"content" validate:"omitempty"`
}

type GetReplyListResponse struct {
	GetListResponse
	Data []models.Reply `json:"data"`
}

type GetReplyResponse struct {
	Message string       `json:"message" example:"success"`
	Data    models.Reply `json:"data"`
}

// GetPostListResponse for GetPostList
type GetPostListResponse struct {
	GetListResponse
	Data []models.Post `json:"data"`
}

// GetPostResponse for GetPost
type GetPostResponse struct {
	Message string      `json:"message" example:"success"`
	Data    models.Post `json:"data"`
}

// GetAssignmentListResponse Response for GetAssignmentList
type GetAssignmentListResponse struct {
	GetListResponse
	Data []models.Assignment `json:"data"`
}

// GetAssignmentResponse Response for get assignment
type GetAssignmentResponse struct {
	Message string            `json:"message" example:"success"`
	Data    models.Assignment `json:"data"`
}

// >>> SUBMITTED ASSIGNMENT <<<
// ============
type SubmittedAssignmentCreateForm struct {
	AssignmentID string `json:"assignmentID" validate:"required,uuid4"`
	Answer       string `json:"answer" validate:"required"`
}

func (f SubmittedAssignmentCreateForm) ConvertToModel() (assignment models.SubmittedAssignment) {
	assignmentID, _ := uuid.FromString(f.AssignmentID)
	assignment = models.SubmittedAssignment{
		AssignmentID: assignmentID,
		Answer:       f.Answer,
	}
	return
}

type SubmittedAssignmentUpdateForm struct {
	AssignmentID *string `json:"assignmentID" validate:"omitempty,uuid4"`
	Answer       *string `json:"answer"  validate:"omitempty"`
	Grade        *int    `json:"grade"  validate:"omitempty"`
	Comment      *string `json:"comment"  validate:"omitempty"`
}

// GetSubmmitedAssignmentListResponse Response for GetSubmmitedAssignmentList
type GetSubmmitedAssignmentListResponse struct {
	GetListResponse
	Data []models.SubmittedAssignment `json:"data"`
}

// GetSubmittedAssignmentResponse Response for get submitted_assignment
type GetSubmittedAssignmentResponse struct {
	Message string                     `json:"message" example:"success"`
	Data    models.SubmittedAssignment `json:"data"`
}

// >>> NOTIFICATION <<<
// ============

// GetNotificationsListResponse Response for GetNotificationList
type GetNotificationsListResponse struct {
	GetListResponse
	Unread int                   `json:"unread"`
	Data   []models.Notification `json:"data"`
}

// GetNotificationResponse Response for get notification
type GetNotificationResponse struct {
	Message string              `json:"message" example:"success"`
	Data    models.Notification `json:"data"`
}
