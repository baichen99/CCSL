package models


// Sign model 手形
type Sign struct {
	Base
	Name  string `gorm:"NOT NULL;UNIQUE;INDEX:name" json:"name"` // 手形名称
	Image string `gorm:"NOT NULL" json:"image"`                  // 手形图片路径
}
