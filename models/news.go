package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Carousel model 首页轮播大图
type Carousel struct {
	Base
	Image      string    `json:"image" example:"news/a.jpg"`
	TitleZh    string    `gorm:"INDEX:title_zh" json:"titleZh" example:"新闻标题"`
	TitleEn    string    `gorm:"INDEX:title_en" json:"titleEn" example:"News Title"`
	Creator    User      `json:"creator"`
	CreatorID  uuid.UUID `gorm:"NOT NULL" json:"creatorID" example:"550e8400-e29b-41d4-a716-446655440000"`
	State      string    `gorm:"DEFAULT:'published'" json:"state"` // State can be 'draft' or 'published'
	Importance int       `gorm:"DEFAULT:0" json:"importance"`
}

// News model
type News struct {
	Base
	Column     string    `json:"column"`                                   // Column:'新闻动态'-> news ,'研究成果' -> research,'学术活动->activity','通知公告'->notice
	Date       time.Time `json:"date"`                                     // Publish date
	Creator    User      `json:"creator"`                                  // Creator
	CreatorID  uuid.UUID `gorm:"TYPE:uuid;NOT NULL" json:"creatorID"`      // News creator
	Title      string    `json:"title"`                                    // News title
	Type       string    `json:"type"`                                     // Type can be 'link' or 'document'
	Text       string    `json:"text"`                                     // If is 'link' type, this field is a url, and document content fot 'ducument' type
	Language   string    `gorm:"NOT NULL;DEFAULT:'zh-CN'" json:"language"` // Language can be 'zh-CN' or 'en-US'
	Importance int       `gorm:"NOT NULL;DEFAULT:0" json:"importance"`
	State      string    `gorm:"NOT NULL;DEFAULT:'published'" json:"state"` // State can be 'draft' or 'published'
}
