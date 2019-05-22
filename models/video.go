package models

import (
	uuid "github.com/satori/go.uuid"
)

// Video model 视频
type Video struct {
	Base
	Word           Word      `gorm:"PRELOAD:true" json:"-"`
	WordID         uuid.UUID `json:"wordID"`
	Performer      Performer `gorm:"PRELOAD:true" json:"-"`
	PerformerID    uuid.UUID `gorm:"DEFAULT:NULL" json:"performerID"`    // 表演者
	ConstructType  string    `gorm:"DEFAULT:NULL" json:"constructType"`  // 构词方式
	ConstructWords string    `gorm:"DEFAULT:NULL" json:"constructWords"` // 构词词语
	LeftSign       string    `gorm:"DEFAULT:NULL" json:"leftSign"`       // 左手手势
	RightSign      string    `gorm:"DEFAULT:NULL" json:"rightSign"`      // 右手手势
	VideoPath      string    `gorm:"DEFAULT:NULL" json:"videoPath"`      // 视频文件路径
}

type VideoResult struct {
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
	LeftSign       string `json:"leftSign"`
	RightSign      string `json:"rightSign"`
	VideoPath      string `json:"videoPath"`
}
