package models

// Performer data model
type Performer struct {
	Base
	Name     string   `gorm:"DEFAULT:NULL" json:"name"`
	Region   District `gorm:"FOREIGNKEY:RegionID;DEFAULT:NULL" json:"region"`
	RegionID int      `gorm:"DEFAULT:NULL" json:"regionID"`
	Gender   string   `gorm:"DEFAULT:NULL" json:"gender"`
}
