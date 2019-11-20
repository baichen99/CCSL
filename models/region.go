package models

// District model
type District struct {
	Code         int    `gorm:"PRIMARY_KEY" json:"code"`
	Name         string `gorm:"NOT NULL" json:"name"`
	CityCode     int    `gorm:"NOT NULL" json:"cityCode"`
	ProvinceCode int    `gorm:"NOT NULL" json:"provinceCode"`
}

// City model
type City struct {
	Code         int    `gorm:"PRIMARY_KEY" json:"code"`
	Name         string `gorm:"NOT NULL" json:"name"`
	ProvinceCode int    `gorm:"NOT NULL" json:"provinceCode"`
}

// Province model
type Province struct {
	Code int    `gorm:"PRIMARY_KEY" json:"code"`
	Name string `gorm:"NOT NULL" json:"name"`
}

type JSONDistrict struct {
	LabelZh string `json:"label"`
	Value   int    `json:"value"`
}

type JSONCity struct {
	LabelZh  string         `json:"label"`
	Value    int            `json:"value"`
	Children []JSONDistrict `json:"children"`
}

type JSONProvince struct {
	LabelZh  string     `json:"label"`
	Value    int        `json:"value"`
	Children []JSONCity `json:"children"`
}
