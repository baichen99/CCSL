package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

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

// Info model
type Info struct {
	Key  string         `gorm:"PRIMARY_KEY" json:"key"`
	Data postgres.Jsonb `gorm:"DEFAULT:NULL" json:"data"`
}
