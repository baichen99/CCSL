package models

// Performer data model
type Performer struct {
	Base
	Name     string   `gorm:"NOT NULL" json:"name"`
	Region   District `gorm:"FOREIGNKEY:RegionID" json:"region"`
	RegionID int      `gorm:"NOT NULL" json:"regionID"`
	Gender   string   `gorm:"NOT NULL" json:"gender"`
}
