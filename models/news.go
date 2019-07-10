package models

import "time"

// Carousel model 首页轮播大图
type Carousel struct {
	Base
	Image string `gorm:"NOT NULL" json:"image"`
	Title string `gorm:"NOT NULL" json:"title"`
}

// News model
type News struct {
	Base
	Column   string    `gorm:"NOT NULL" json:"column"` // Column：'通知公告','研究成果','学术活动','资料下载'
	Date     time.Time `gorm:"NOT NULL" json:"date"`   // Data
	Title    string    `gorm:"NOT NULL" json:"title"`  // Title
	Type     string    `gorm:"NOT NULL" json:"type"`   // Type can be 'link' or 'text'
	Text     string    `gorm:"NOT NULL" json:"text"`   // Text
	Language string    `gorm:"NOT NULL" json:"language"`
}
