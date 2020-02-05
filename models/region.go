package models

// District model
type District struct {
	Code         int    `gorm:"primary_key" json:"code"`
	Name         string `gorm:"not null" json:"name"`
	CityCode     int    `gorm:"not null" json:"cityCode"`
	ProvinceCode int    `gorm:"not null" json:"provinceCode"`
}

// City model
type City struct {
	Code         int    `gorm:"primary_key" json:"code"`
	Name         string `gorm:"not null" json:"name"`
	ProvinceCode int    `gorm:"not null" json:"provinceCode"`
}

// Province model
type Province struct {
	Code int    `gorm:"primary_key" json:"code"`
	Name string `gorm:"not null" json:"name"`
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
