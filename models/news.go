package models

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// Carousel model 首页轮播大图
type Carousel struct {
	Base
	Image      string    `gorm:"NOT NULL" json:"image"`
	Title      string    `gorm:"NOT NULL" json:"title"`
	Creator    User      `json:"creator"`
	CreatorID  uuid.UUID `gorm:"DEFAULT:NULL" json:"creatorID"`
	State      string    `gorm:"DEFAULT:'published'" json:"state"` // State can be 'draft' or 'published'
	Importance int       `gorm:"DEFAULT:0" json:"importance"`
}

// News model
type News struct {
	Base
	Column      string         `gorm:"NOT NULL" json:"column"`                   // Column：'新闻动态'-> news ,'研究成果' -> research,'学术活动->activity','通知公告'->notice
	Date        time.Time      `gorm:"NOT NULL" json:"date"`                     // Publish date
	Creator     User           `json:"creator"`                                  // Creator
	CreatorID   uuid.UUID      `gorm:"DEFAULT:NULL" json:"creatorID"`            // News creator
	Title       string         `gorm:"DEFAULT:NULL" json:"title"`                // News title
	Type        string         `gorm:"DEFAULT:NULL" json:"type"`                 // Type can be 'link' or 'document'
	Text        string         `gorm:"DEFAULT:NULL" json:"text"`                 // If is 'link' type, this field is a url, and document content fot 'ducument' type
	Language    string         `gorm:"NOT NULL;DEFAULT:'zh-CN'" json:"language"` // Language can be 'zh-CN' or 'en-US'
	Importance  int            `gorm:"DEFAULT:0" json:"importance"`
	State       string         `gorm:"DEFAULT:'published'" json:"state"` // State can be 'draft' or 'published'
	Attachments pq.StringArray `gorm:"TYPE:varchar(200)[];DEFAULT:array[]::varchar(200)[]" json:"attachments"`
}
