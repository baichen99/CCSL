package tools

import (
    "ccsl/models"
    "github.com/jinzhu/gorm"
    uuid "github.com/satori/go.uuid"
    "strings"
)

type VideoSign struct {
    VideoID uuid.UUID   `gorm:"TYPE:uuid"`
    SignID  uuid.UUID   `gorm:"TYPE:uuid"`
}

// Migrate 将video表中的sign分离出来, 并建立两张(left right) 关联表
func Migrate(db *gorm.DB) {
    // 初始化
	db.DropTableIfExists(&models.Sign{}, "video_leftsign", "video_rightsign")
	db.CreateTable(&models.Sign{})
	db.Table("video_leftsign").CreateTable(&VideoSign{})
	db.Table("video_rightsign").CreateTable(&VideoSign{})
	// 遍历videos表中的数据, 迁移到新表
	var videos []models.Video
	var sign models.Sign
	err := db.Find(&videos).Error
	if err != nil {
        return
    }
	for _, v := range videos {
        if(v.LeftSign != "") {
            nums := strings.Split(v.LeftSign, ",")
            for _, n := range nums {
                if db.Table("signs").Where("name = ?", n).Scan(&sign).RecordNotFound() {
                    // 如果在signs表中没找到, 新建一个sign记录
                    db.Table("signs").Create(&models.Sign{Name: n})
                    //  然后建立关系
                    db.Table("signs").Where("name = ?", n).Scan(&sign)
                    err = db.Table("video_leftsign").Create(&VideoSign{VideoID: v.ID, SignID: sign.ID}).Error
                } else {
                    // 如果找到, 建立关系
                    err = db.Table("video_leftsign").Create(&VideoSign{VideoID: v.ID, SignID: sign.ID}).Error
                }
                if err != nil  {
                    return
                }
            }
        }

        if(v.RightSign != "") {
            nums := strings.Split(v.RightSign, ",")
            for _, n := range nums {
                if db.Table("signs").Where("name = ?", n).Scan(&sign).RecordNotFound() {
                    // 如果在signs表中没找到, 新建一个sign记录
                    db.Table("signs").Create(&models.Sign{Name: n})
                    //  然后建立关系
                    db.Table("signs").Where("name = ?", n).Scan(&sign)
                    err = db.Table("video_rightsign").Create(&VideoSign{VideoID: v.ID, SignID: sign.ID}).Error
                } else {
                    // 如果找到, 建立关系
                    err = db.Table("video_rightsign").Create(&VideoSign{VideoID: v.ID, SignID: sign.ID}).Error
                }
                if err != nil  {
                    return
                }
            }
        }
    }
}
