package models

// Performer data model
type Performer struct {
	Base
	Name   string `gorm:"NOT NULL" json:"name"`
	Region string `gorm:"NOT NULL" json:"region"`
	Gender string `gorm:"NOT NULL" json:"gender"`
}
