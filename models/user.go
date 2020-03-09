package models

import "github.com/lib/pq"

// User data model
type User struct {
	Base
	Avatar       string         `gorm:"default:'https://ccsl.shu.edu.cn/public/assets/default.png'" json:"avatar" example:"https://ccsl.shu.edu.cn/public/assets/default.png"` // 头像
	Username     string         `gorm:"not null;index:username" json:"username" example:"adrianduan@icloud.com"`                                                               // 账号（一卡通或者邮箱）
	Name         string         `json:"name" example:"Adrian"`                                                                                                                 // 姓名
	Password     string         `json:"-"`                                                                                                                                     // 密码
	State        string         `gorm:"default:'inactive'" json:"state" example:"active"`                                                                                      // 状态
	Roles        pq.StringArray `gorm:"type:text[];not null;default:array[]::text[]" json:"roles" swaggertype:"array,string" example:"super, admin"`                           // 用户角色
	StudentClass *[]Class       `json:"-" gorm:"many2many:class_students"`
	TeacherClass *[]Class       `json:"-" gorm:"many2many:class_teachers"`
}
