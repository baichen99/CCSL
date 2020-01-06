package models

import uuid "github.com/satori/go.uuid"

// User data model
type User struct {
	Base
	Avatar   string `gorm:"DEFAULT:'https://ccsl.shu.edu.cn/public/assets/default.png'" json:"avatar" example:"https://ccsl.shu.edu.cn/public/assets/default.png"` // 头像
	Username string `gorm:"NOT NULL;INDEX:username" json:"username" example:"adrianduan@icloud.com"`                                                               // 账号（一卡通或者邮箱）
	Name     string `gorm:"DEFAULT:NULL" json:"name" example:"Adrian Duan"`                                                                                        // 姓名
	UserType string `gorm:"DEFAULT:'user'" json:"userType" example:"admin"`                                                                                        // 用户角色
	Password string `gorm:"DEFAULT:NULL" json:"-"`                                                                                                                 // 密码
	State    string `gorm:"DEFAULT:'inactive'" json:"state" example:"active"`                                                                                      // 状态
}

// LoginHistory model
type LoginHistory struct {
	Base
	User         User      `json:"user"`
	UserID       uuid.UUID `json:"userID"`
	Status       string    `json:"status"`
	IP           string    `json:"ip"`
	Country      string    `gorm:"DEFAULT:NULL" json:"country"`
	CountryCode  string    `gorm:"DEFAULT:NULL" json:"countryCode"`
	Region       string    `gorm:"DEFAULT:NULL" json:"region"`
	RegionName   string    `gorm:"DEFAULT:NULL" json:"regionName"`
	City         string    `gorm:"DEFAULT:NULL" json:"city"`
	District     string    `gorm:"DEFAULT:NULL" json:"district"`
	Longitude    float64   `gorm:"DEFAULT:NULL" json:"lon"`
	Latitude     float64   `gorm:"DEFAULT:NULL" json:"lat"`
	Timezone     string    `gorm:"DEFAULT:NULL" json:"timezone"`
	ISP          string    `gorm:"DEFAULT:NULL" json:"isp"`
	Organization string    `gorm:"DEFAULT:NULL" json:"org"`
}
