package models

// User data model
type User struct {
	Base
	Avatar   string `gorm:"DEFAULT:'https://ccsl.shu.edu.cn/public/assets/default.png'" json:"avatar"`
	Username string `gorm:"NOT NULL;INDEX:username" json:"username"`
	Name     string `gorm:"DEFAULT:NULL" json:"name"`
	UserType string `gorm:"DEFAULT:'user'" json:"userType"`
	Password string `gorm:"DEFAULT:NULL" json:"-"`
	State    string `gorm:"DEFAULT:'inactive'" json:"state"`
}
