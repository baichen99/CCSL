package models

// User data model
type User struct {
	Base
	Username string `gorm:"NOT NULL;UNIQUE;INDEX:username" json:"username"`
	Name     string `gorm:"DEFAULT:NULL" json:"name"`
	UserType string `gorm:"DEFAULT:'user'" json:"userType"`
	Password string `gorm:"DEFAULT:NULL" json:"-"`
}
