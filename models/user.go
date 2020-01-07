package models

// User data model
type User struct {
	Base
	Avatar   string `gorm:"DEFAULT:'https://ccsl.shu.edu.cn/public/assets/default.png'" json:"avatar" example:"https://ccsl.shu.edu.cn/public/assets/default.png"` // 头像
	Username string `gorm:"NOT NULL;INDEX:username" json:"username" example:"adrianduan@icloud.com"`                                                               // 账号（一卡通或者邮箱）
	Name     string `gorm:"DEFAULT:NULL" json:"name" example:"Adrian Duan"`                                                                                        // 姓名
	UserType string `gorm:"DEFAULT:'user'" json:"userType" example:"admin"`                                                                                        // 用户角色
	Password string `gorm:"DEFAULT:NULL" json:"-"`                                                                                                                 // 密码
	State    string `gorm:"DEFAULT:'inactive'" json:"state" example:"active"`                                                                                      // 状态
}
