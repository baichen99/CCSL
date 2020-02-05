package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// LexicalVideo mdoel for Lexical Database
type LexicalVideo struct {
	Base
	WordFormation     string         `gorm:"index:word_formation" json:"wordFormation" example:"simple"`                                                                                                                              // 构词方式
	VideoPath         string         `json:"videoPath" example:"lexical_videos/550e8400-e29b-41d4-a716-446655440000.mp4"`                                                                                                             // 视频路径
	LexiconID         uuid.UUID      `gorm:"type:uuid;not null;index:lexicon_id" json:"lexiconID" example:"550e8400-e29b-41d4-a716-446655440000"`                                                                                     // 词汇ID
	PerformerID       uuid.UUID      `gorm:"type:uuid;not null;index:performer_id" json:"performerID" example:"550e8400-e29b-41d4-a716-446655440000"`                                                                                 // 被试ID
	Morpheme          pq.StringArray `gorm:"type:text[];not null;default:array[]::text[];index:morpheme" json:"morpheme" swaggertype:"array,string" example:"hello,world"`                                                            // 构词语素
	LeftHandshapesID  pq.StringArray `gorm:"type:uuid[];not null;default:array[]::uuid[];index:left_handshapes_id" json:"leftHandshapesID" swaggertype:"array,string" format:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`   // 左手手形ID
	RightHandshapesID pq.StringArray `gorm:"type:uuid[];not null;default:array[]::uuid[];index:right_handshapes_id" json:"rightHandshapesID" swaggertype:"array,string" format:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"` // 右手手形ID
}
