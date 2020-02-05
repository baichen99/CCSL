package models

// Performer data model
type Performer struct {
	Base
	RegionID int    `json:"regionID"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
}
