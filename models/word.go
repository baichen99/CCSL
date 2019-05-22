package models

// Word model
type Word struct {
	Base
	Initial string `gorm:"NOT NULL" json:"initial"`        // 首字母
	Chinese string `gorm:"NOT NULL;UNIQUE" json:"chinese"` // 汉语转写
	English string `gorm:"NOT NULL" json:"english"`        // 英语转写
	Type    string `gorm:"DEFAULT:NULL" json:"type"`       // 词性
}
