package models

// Member mdoel for team members
type Member struct {
	Base
	Profile       string `gorm:"DEFAULT:NULL" json:"profile"`       // 照片文件路径
	Type          string `gorm:"DEFAULT:NULL" json:"type"`          // 类型：高级顾问、研究员……
	NameZh        string `gorm:"DEFAULT:NULL" json:"nameZh"`        // 姓名 中文
	NameEn        string `gorm:"DEFAULT:NULL" json:"nameEn"`        // 姓名 英文
	PositionZh    string `gorm:"DEFAULT:NULL" json:"positionZh"`    // 职称 中文
	PositionEn    string `gorm:"DEFAULT:NULL" json:"positionEn"`    // 职称 英文
	EmployerZh    string `gorm:"DEFAULT:NULL" json:"employerZh"`    // 工作单位 中文
	EmployerEn    string `gorm:"DEFAULT:NULL" json:"employerEn"`    // 工作单位 英文
	DescriptionZh string `gorm:"DEFAULT:NULL" json:"descriptionZh"` // 个人介绍 中文
	DescriptionEn string `gorm:"DEFAULT:NULL" json:"descriptionEn"` // 个人介绍 英文
}
