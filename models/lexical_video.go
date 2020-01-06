package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// LexicalVideo mdoel for Lexical Database
type LexicalVideo struct {
	Base
	WordFormation string         `gorm:"DEFAULT:NULL;INDEX:word_formation" json:"wordFormation" example:"simple"`                                                         // 构词方式
	VideoPath     string         `gorm:"DEFAULT:NULL" json:"videoPath" example:"lexical_videos/550e8400-e29b-41d4-a716-446655440000.mp4"`                                 // 视频路径
	LexiconID     uuid.UUID      `gorm:"TYPE:uuid;NOT NULL;INDEX:lexicon_id" json:"lexiconID" example:"550e8400-e29b-41d4-a716-446655440000"`                             // 词汇ID
	PerformerID   uuid.UUID      `gorm:"TYPE:uuid;NOT NULL;INDEX:performer_id" json:"performerID" example:"550e8400-e29b-41d4-a716-446655440000"`                         // 被试ID
	Morpheme      pq.StringArray `gorm:"TYPE:text[];DEFAULT:array[]::text[];INDEX:morpheme" json:"morpheme" example:"['hello','world']"`                                  // 构词语素
	LeftSignsID   pq.StringArray `gorm:"TYPE:uuid[];DEFAULT:array[]::uuid[];INDEX:left_signs_id" json:"leftSignsID" example:"['550e8400-e29b-41d4-a716-446655440000']"`   // 左手手形ID
	RightSignsID  pq.StringArray `gorm:"TYPE:uuid[];DEFAULT:array[]::uuid[];INDEX:right_signs_id" json:"rightSignsID" example:"['550e8400-e29b-41d4-a716-446655440000']"` // 右手手形ID
}
