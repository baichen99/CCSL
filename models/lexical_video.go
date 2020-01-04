package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// LexicalVideo mdoel for Lexical Database
type LexicalVideo struct {
	Base
	Lexicon       Lexicon        `gorm:"FOREIGNKEY:LexiconID" json:"lexicon"`
	LexiconID     uuid.UUID      `gorm:"NOT NULL;INDEX:lexicon_id" json:"lexiconID"`
	Performer     Performer      `gorm:"FOREIGNKEY:PerformerID" json:"performer"`
	PerformerID   uuid.UUID      `gorm:"NOT NULL;INDEX:performer_id" json:"performerID"`                     // 表演者
	WordFormation string         `gorm:"DEFAULT:NULL;INDEX:word_formation" json:"wordFormation"`             // 构词方式
	Morpheme      pq.StringArray `gorm:"TYPE:text[];DEFAULT:array[]::text[];INDEX:morpheme" json:"morpheme"` // 构词词语
	VideoPath     string         `gorm:"DEFAULT:NULL" json:"videoPath"`
	LeftSignsID   []string       `gorm:"-" json:"leftSignsID"`
	RightSignsID  []string       `gorm:"-" json:"rightSignsID"`                          // 视频文件路径
	LeftSigns     []Sign         `gorm:"MANY2MANY:lexical_left_sign" json:"leftSigns"`   // 左手手势
	RightSigns    []Sign         `gorm:"MANY2MANY:lexical_right_sign" json:"rightSigns"` // 右手手势
}
