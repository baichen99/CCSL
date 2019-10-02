package models

import (
	uuid "github.com/satori/go.uuid"
)

// LexicalVideo  Lexical Database for Chinese National Sign Language
type LexicalVideo struct {
	Base
	LexicalWord    LexicalWord `gorm:"PRELOAD:true" json:"word"`
	LexicalWordID  uuid.UUID   `gorm:"NOT NULL" json:"wordID"`
	Performer      Performer   `gorm:"PRELOAD:true" json:"performer"`
	PerformerID    uuid.UUID   `gorm:"NOT NULL" json:"performerID"`                     // 表演者
	ConstructType  string      `gorm:"DEFAULT:NULL" json:"constructType"`               // 构词方式
	ConstructWords string      `gorm:"DEFAULT:NULL" json:"constructWords"`              // 构词词语
	VideoPath      string      `gorm:"DEFAULT:NULL" json:"videoPath"`                   // 视频文件路径
	LeftSigns      []Sign      `gorm:"many2many:lexical_left_sign;" json:"leftSigns"`   // 左手手势
	RightSigns     []Sign      `gorm:"many2many:lexical_right_sign;" json:"rightSigns"` // 右手手势
}

type LexicalVideoResult struct {
	ID             string `json:"id"`
	Initial        string `json:"initial"`
	Chinese        string `json:"chinese"`
	English        string `json:"english"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	Region         string `json:"region"`
	Gender         string `json:"gender"`
	ConstructType  string `json:"constructType"`
	ConstructWords string `json:"constructWords"`
	LeftSigns      string `json:"leftSigns"`
	RightSigns     string `json:"rightSigns"`
	VideoPath      string `json:"videoPath"`
}
