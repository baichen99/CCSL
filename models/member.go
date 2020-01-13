package models

// Member mdoel for team members
type Member struct {
	Base
	Profile       string `json:"profile" example:"members/a.jpg"`          // 照片文件路径
	Type          string `json:"type" example:"researchFellow"`            // 类型：高级顾问、研究员……
	NameZh        string `json:"nameZh" example:"张三"`                      // 姓名 中文
	NameEn        string `json:"nameEn" example:"Zhang San"`               // 姓名 英文
	Degree        string `json:"degree" example:"doctor"`                  // 学位
	EmployerZh    string `json:"employerZh" example:"上海大学"`                // 工作单位 中文
	EmployerEn    string `json:"employerEn" example:"Shanghai University"` // 工作单位 英文
	DescriptionZh string `json:"descriptionZh" example:"个人介绍"`             // 个人介绍 中文
	DescriptionEn string `json:"descriptionEn" example:"Description"`      // 个人介绍 英文
}
