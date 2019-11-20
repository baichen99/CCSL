package models

import "github.com/jinzhu/gorm/dialects/postgres"

// JsError model
type JsError struct {
	Base
	Err   postgres.Jsonb `json:"err" validate:"omitempty"`
	Store postgres.Jsonb `json:"store" validate:"omitempty"`
	Info  string         `json:"info" validate:"omitempty"`
	URL   string         `json:"url" validate:"omitempty"`
}

// Info model
type Info struct {
	Key  string         `gorm:"PRIMARY_KEY" json:"key"`
	Data postgres.Jsonb `gorm:"DEFAULT:NULL" json:"data"`
}
