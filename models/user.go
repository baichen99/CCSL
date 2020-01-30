package models

import "github.com/lib/pq"

// User data model
type User struct {
	Base
	Avatar   string         `gorm:"DEFAULT:'https://ccsl.shu.edu.cn/public/assets/default.png'" json:"avatar" example:"https://ccsl.shu.edu.cn/public/assets/default.png"` // 头像
	Username string         `gorm:"NOT NULL;INDEX:username" json:"username" example:"adrianduan@icloud.com"`                                                               // 账号（一卡通或者邮箱）
	Name     string         `gorm:"DEFAULT:NULL" json:"name" example:"Adrian Duan"`                                                                                        // 姓名                                                                                     // 用户角色
	Password string         `gorm:"DEFAULT:NULL" json:"-"`                                                                                                                 // 密码
	State    string         `gorm:"DEFAULT:'inactive'" json:"state" example:"active"`                                                                                      // 状态
	Roles    pq.StringArray `gorm:"TYPE:text[];NOT NULL;DEFAULT:array[]::text[]" json:"roles" swaggertype:"array,string" example:"super,admin"`
}
