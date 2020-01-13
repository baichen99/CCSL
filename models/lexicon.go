package models

// Lexicon words for Lexical Database
type Lexicon struct {
	Base
	Initial string `json:"initial" example:"W"` // 汉语音序
	Chinese string `json:"chinese" example:"我"` // 汉语转写
	English string `json:"english" example:"I"` // 英语转写
	Pos     string `json:"pos" example:"代词"`    // 词性
}
