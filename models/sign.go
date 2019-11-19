package models

// Sign model 手形
type Sign struct {
	Base
	Name  string `gorm:"DEFAULT:NULL;INDEX:name" json:"name"` // 手形名称
	Image string `gorm:"DEFAULT:NULL" json:"image"`           // 手形图片路径
}
