package models

// Lexicon words for Lexical Database
type Lexicon struct {
	Base
	Initial string `gorm:"DEFAULT:NULL" json:"initial"` // 首字母
	Chinese string `gorm:"DEFAULT:NULL" json:"chinese"` // 汉语转写
	English string `gorm:"DEFAULT:NULL" json:"english"` // 英语转写
	Pos     string `gorm:"DEFAULT:NULL" json:"pos"`     // 词性
}
