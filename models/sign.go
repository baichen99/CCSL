package models

import (
	"regexp"
	"strconv"
)

// Sign model 手形
type Sign struct {
	Base
	Name  string `gorm:"DEFAULT:NULL;INDEX:name" json:"name"` // 手形名称
	Image string `gorm:"DEFAULT:NULL" json:"image"`           // 手形图片路径
}

// Signs is alias for []Sign
type Signs []Sign

// Len returns length
func (s Signs) Len() int {
	return len(s)
}

// Swap changes the position of element
func (s Signs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns if e[i] is less than e[j]
func (s Signs) Less(i, j int) bool {
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
