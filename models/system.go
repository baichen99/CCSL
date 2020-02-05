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
