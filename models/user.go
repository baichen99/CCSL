package models

import uuid "github.com/satori/go.uuid"

// User data model
type User struct {
	Base
	Avatar   string `gorm:"DEFAULT:'https://ccsl.shu.edu.cn/public/assets/default.png'" json:"avatar"`
	Username string `gorm:"NOT NULL;INDEX:username" json:"username"`
	Name     string `gorm:"DEFAULT:NULL" json:"name"`
	UserType string `gorm:"DEFAULT:'user'" json:"userType"`
	Password string `gorm:"DEFAULT:NULL" json:"-"`
	State    string `gorm:"DEFAULT:'inactive'" json:"state"`
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
