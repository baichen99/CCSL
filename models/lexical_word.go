package models

// LexicalWord words for Lexical Database
type LexicalWord struct {
	Base
	Initial string `gorm:"NOT NULL" json:"initial"` // 首字母
	Chinese string `gorm:"NOT NULL" json:"chinese"` // 汉语转写
	English string `gorm:"NOT NULL" json:"english"` // 英语转写
	Pos     string `gorm:"DEFAULT:NULL" json:"pos"` // 词性
}
