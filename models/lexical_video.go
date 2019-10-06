package models

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// LexicalVideo mdoel for Lexical Database
type LexicalVideo struct {
	Base
	LexicalWord    LexicalWord    `gorm:"FOREIGNKEY:LexicalWordID" json:"word"`
	LexicalWordID  uuid.UUID      `gorm:"NOT NULL;INDEX:lexical_word_id" json:"-"`
	Performer      Performer      `gorm:"FOREIGNKEY:PerformerID" json:"performer"`
	PerformerID    uuid.UUID      `gorm:"NOT NULL;INDEX:performer_id" json:"-"`                   // 表演者
	ConstructType  string         `gorm:"DEFAULT:NULL" json:"constructType"`                      // 构词方式
	ConstructWords pq.StringArray `gorm:"TYPE:varchar(100)[];DEFAULT:NULL" json:"constructWords"` // 构词词语
	VideoPath      string         `gorm:"DEFAULT:NULL" json:"videoPath"`                          // 视频文件路径
	LeftSigns      []Sign         `gorm:"MANY2MANY:lexical_left_sign" json:"leftSigns"`           // 左手手势
	RightSigns     []Sign         `gorm:"MANY2MANY:lexical_right_sign" json:"rightSigns"`         // 右手手势
}

// VideoSign model
type VideoSign struct {
	LexicalVideoID uuid.UUID
	SignID         uuid.UUID
}

// VideoSearchResult model
type VideoSearchResult struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	Initial        string    `json:"initial"`
	Chinese        string    `json:"chinese"`
	English        string    `json:"english"`
	Pos            string    `json:"pos"`
	Name           string    `json:"name"`
	RegionID       int       `json:"regionID"`
	Region         string    `json:"region"`
	Gender         string    `json:"gender"`
	ConstructType  string    `json:"constructType"`
	ConstructWords string    `json:"constructWords"`
	VideoPath      string    `json:"videoPath"`
}
