package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

// Member mdoel for team members
type Member struct {
	Base
	Profile       string `json:"profile" example:"members/a.jpg"`          // 照片文件路径
	Type          string `json:"type" example:"researchFellow"`            // 类型：高级顾问、研究员……
	NameZh        string `json:"nameZh" example:"张三"`                      // 姓名 中文
	NameEn        string `json:"nameEn" example:"Zhang San"`               // 姓名 英文
	Degree        string `json:"degree" example:"doctor"`                  // 学位
	EmployerZh    string `json:"employerZh" example:"上海大学"`                // 工作单位 中文
	EmployerEn    string `json:"employerEn" example:"Shanghai University"` // 工作单位 英文
	DescriptionZh string `json:"descriptionZh" example:"个人介绍"`             // 个人介绍 中文
	DescriptionEn string `json:"descriptionEn" example:"Description"`      // 个人介绍 英文
}

// JsError model
type JsError struct {
	Base
	Err   postgres.Jsonb `json:"err" validate:"omitempty"`
	Store postgres.Jsonb `json:"store" validate:"omitempty"`
	Info  string         `json:"info" validate:"omitempty"`
	URL   string         `json:"url" validate:"omitempty"`
}

// LoginHistory model
type LoginHistory struct {
	Base
	User         User      `json:"user"`
	UserID       uuid.UUID `json:"userID"`
	Status       string    `json:"status"`
	IP           string    `json:"ip"`
	Country      string    `json:"country"`
	CountryCode  string    `json:"countryCode"`
	Region       string    `json:"region"`
	RegionName   string    `json:"regionName"`
	City         string    `json:"city"`
	District     string    `json:"district"`
	Longitude    float64   `json:"lon"`
	Latitude     float64   `json:"lat"`
	Timezone     string    `json:"timezone"`
	ISP          string    `json:"isp"`
	Organization string    `json:"org"`
}

// Info model
type Info struct {
	Key  string         `gorm:"primary_key" json:"key"`
	Data postgres.Jsonb `json:"data"`
}
