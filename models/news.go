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
	Column string `gorm:"NOT NULL" json:"column"`
	// Column：'新闻动态'-> news ,'研究成果' -> research,'学术活动->activity','通知公告'->notice
	Date     time.Time `gorm:"NOT NULL" json:"date"`
	Title    string    `gorm:"NOT NULL" json:"title"`
	Type     string    `gorm:"NOT NULL" json:"type"` // Type can be 'link' or 'document'
	Text     string    `gorm:"NOT NULL" json:"text"`
	Language string    `gorm:"NOT NULL;DEFAULT:'zh-CN'" json:"language"`
}
