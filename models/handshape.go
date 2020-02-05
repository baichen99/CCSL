package models

import (
	"regexp"
	"strconv"
)

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
