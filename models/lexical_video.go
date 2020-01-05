package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// LexicalVideo mdoel for Lexical Database
type LexicalVideo struct {
	Base
	Lexicon       Lexicon        `gorm:"FOREIGNKEY:LexiconID" json:"lexicon"`
	LexiconID     uuid.UUID      `gorm:"TYPE:uuid;NOT NULL;INDEX:lexicon_id" json:"lexiconID"`
	Performer     Performer      `gorm:"FOREIGNKEY:PerformerID" json:"performer"`
	PerformerID   uuid.UUID      `gorm:"TYPE:uuid;NOT NULL;INDEX:performer_id" json:"performerID"`
	WordFormation string         `gorm:"DEFAULT:NULL;INDEX:word_formation" json:"wordFormation"`
	Morpheme      pq.StringArray `gorm:"TYPE:text[];DEFAULT:array[]::text[];INDEX:morpheme" json:"morpheme"`
	LeftSignsID   pq.StringArray `gorm:"TYPE:uuid[];DEFAULT:array[]::uuid[];INDEX:left_signs_id" json:"leftSignsID"`
	RightSignsID  pq.StringArray `gorm:"TYPE:uuid[];DEFAULT:array[]::uuid[];INDEX:right_signs_id" json:"rightSignsID"`
	VideoPath     string         `gorm:"DEFAULT:NULL" json:"videoPath"`
}
