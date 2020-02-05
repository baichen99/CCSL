package models

// Class model
type Class struct {
	Base
	Name     string   `json:"name"`
	Teachers []User   `gorm:"many2many:class_teachers" json:"teachers"`
	Students []User   `gorm:"many2many:class_students" json:"students"`
	Details  string   `json:"details"`
	Resource string   `json:"resource"`
	Courses  []Course `json:"courses"`
}
