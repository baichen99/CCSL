package models

import "github.com/lib/pq"

// Lexicon words for Lexical Database
type Lexicon struct {
	Base
	Initial string         `json:"initial" example:"W"`                                                                                              // 汉语音序
	Chinese string         `json:"chinese" example:"我"`                                                                                              // 汉语转写
	English string         `json:"english" example:"I"`                                                                                              // 英语转写
	Pos     pq.StringArray `gorm:"type:text[];not null;default:array[]::text[];index:pos" json:"pos" swaggertype:"array,string" example:"noun,verb"` // 词性
}
