package models

import (
	"regexp"
	"strconv"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// Lexicon model
type Lexicon struct {
	Base
	Initial string         `json:"initial" example:"W"`                                                                                              // 汉语音序
	Chinese string         `json:"chinese" example:"我"`                                                                                              // 汉语转写
	English string         `json:"english" example:"I"`                                                                                              // 英语转写
	Pos     pq.StringArray `gorm:"type:text[];not null;default:array[]::text[];index:pos" json:"pos" swaggertype:"array,string" example:"noun,verb"` // 词性
}

// LexicalVideo mdoel
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

// Handshape model 手形
type Handshape struct {
	Base
	Name  string `gorm:"index:name" json:"name"` // Name of handshape
	Image string `json:"image"`                  // Image path of handshape
	Glyph string `json:"glyph"`                  // Font glyph of handshape
}

// Handshapes is alias for []Handshape
type Handshapes []Handshape

// Len returns length
func (s Handshapes) Len() int {
	return len(s)
}

// Swap changes the position of element
func (s Handshapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns if e[i] is less than e[j]
func (s Handshapes) Less(i, j int) bool {
	nameI := s[i].Name
	nameJ := s[j].Name
	intI := 0
	intJ := 0
	reg := regexp.MustCompile(`^\d+`)
	strI := reg.FindAllString(nameI, -1)
	strJ := reg.FindAllString(nameJ, -1)
	if len(strI) > 0 {
		intI, _ = strconv.Atoi(strI[0])
	}
	if len(strJ) > 0 {
		intJ, _ = strconv.Atoi(strJ[0])
	}
	return intI < intJ
}
