package models

// Performer data model
type Performer struct {
	Base
	RegionID int    `gorm:"DEFAULT:NULL" json:"regionID"`
	Name     string `gorm:"DEFAULT:NULL" json:"name"`
	Gender   string `gorm:"DEFAULT:NULL" json:"gender"`
	// Region   District `gorm:"FOREIGNKEY:RegionID" json:"region"`
}
