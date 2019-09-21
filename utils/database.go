package utils

import (
	"ccsl/configs"
	"ccsl/models"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"
)

// ConnectPostgres : connect to postgresql database
func ConnectPostgres(app *iris.Application) *gorm.DB {
	databaseURL := "postgres://" + configs.Conf.Postgresql.User + ":" + configs.Conf.Postgresql.Password + "@" +
		configs.Conf.Postgresql.Server + ":" + strconv.Itoa(configs.Conf.Postgresql.Port) + "/" +
		configs.Conf.Postgresql.Database + "?sslmode=disable"
	pg, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		app.Logger().Errorf("Open PG Error: " + err.Error())
		panic(err)
	}
	return pg
}

// InitTestUser creates test users
func InitTestUser(pg *gorm.DB) {
	password, _ := HashPassword("P@ssw0rd")
	testUsers := []models.User{
		models.User{
			Username: "admin@ccsl.shu.edu.cn",
			Password: password,
			UserType: "admin",
		},
		models.User{
			Username: "user@ccsl.shu.edu.cn",
			Password: password,
			UserType: "user",
		},
		models.User{
			Username: "learner@ccsl.shu.edu.cn",
			Password: password,
			UserType: "learner",
		},
		models.User{
			Username: "super@ccsl.shu.edu.cn",
			Password: password,
			UserType: "super",
		},
	}
	for _, user := range testUsers {
		pg.Create(&user)
	}
}

// InitProdUser inits user on production server
func InitProdUser(pg *gorm.DB) {
	password, _ := HashPassword("learning!CCSL")
	users := []models.User{
		models.User{
			Username: "learning@ccsl.shu.edu.cn",
			Password: password,
			UserType: "user",
			Name:     "学习平台公用账号",
		},
	}
	for _, user := range users {
		pg.Create(&user)
	}
	superUsers := []models.User{
		models.User{
			Username: "16121041",
			UserType: "super",
			Name:     "段靖",
		},
		models.User{
			Username: "10008119",
			UserType: "super",
			Name:     "倪兰",
		},
	}
	for _, superUser := range superUsers {
		pg.Create(&superUser)
	}
}

// UpdateVideoPath updates video_path
func UpdateVideoPath(pg *gorm.DB) {
	targetDir := "/Users/adrian/ErrorSignData/"
	moveDir := "/Users/adrian/Data/"
	dir, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, file := range dir {
		fileName := file.Name()
		word := strings.Split(strings.Split(fileName, "-")[2], ".")[0]
		performer := strings.Split(fileName, "-")[1]
		var wordModel models.Word
		var performerModel models.Performer
		if pg.Where("chinese = ?", word).Take(&wordModel).RecordNotFound() {
			cmd := exec.Command("/bin/bash", "-c", "echo '"+fileName+"' >> error.txt")
			move := exec.Command("/bin/bash", "-c", "mv '"+targetDir+fileName+"' '"+moveDir+fileName+"'")
			if _, err = cmd.Output(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if _, err = move.Output(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			continue
		}
		if pg.Where("name = ?", performer).Take(&performerModel).RecordNotFound() {
			cmd := exec.Command("/bin/bash", "-c", "echo '"+fileName+" PerformerNotFound"+"' >> error.txt")
			move := exec.Command("/bin/bash", "-c", "mv '"+targetDir+fileName+"' '"+moveDir+fileName+"'")
			if _, err = cmd.Output(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if _, err = move.Output(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			continue
		}
		pg.Model(&models.Video{}).Where("word_id = ? AND performer_id = ?", wordModel.ID, performerModel.ID).Update("video_path", fileName)
	}
}

type VideoSign struct {
	VideoID uuid.UUID `gorm:"TYPE:uuid"`
	SignID  uuid.UUID `gorm:"TYPE:uuid"`
}

// Migrate 将video表中的sign分离出来, 并建立两张(left right) 关联表
func Migrate(db *gorm.DB) {
	// 初始化
	// db.DropTableIfExists(&models.Sign{}, "video_leftsign", "video_rightsign")
	// db.CreateTable(&models.Sign{})
	// db.Table("video_leftsign").CreateTable(&VideoSign{})
	// db.Table("video_rightsign").CreateTable(&VideoSign{})
	// 这里不需要初始化 AutoMigrate()会自动初始化表结构并且添加主键
	// 遍历videos表中的数据, 迁移到新表
	var videos []models.Video
	var sign models.Sign
	err := db.Raw("SELECT * FROM videos WHERE left_sign NOTNULL OR right_sign NOTNULL").Scan(&videos).Error
	if err != nil {
		cmd := exec.Command("/bin/bash", "-c", "echo '"+err.Error()+"' >> error.txt")
		cmd.Output()
		return
	}
	for _, v := range videos {
		if v.LeftSign != "" {
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
					if db.Table("video_leftsign").Where("video_id = ?", v.ID).Where("sign_id = ?", sign.ID).Find(&VideoSign{}).RecordNotFound() {
						err = db.Table("video_leftsign").Create(&VideoSign{VideoID: v.ID, SignID: sign.ID}).Error
					}
				}
				if err != nil {
					cmd := exec.Command("/bin/bash", "-c", "echo '"+err.Error()+"' >> error.txt")
					cmd.Output()
					continue
				}
			}
		}

		if v.RightSign != "" {
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
					if db.Table("video_rightsign").Where("video_id = ?", v.ID).Where("sign_id = ?", sign.ID).Find(&VideoSign{}).RecordNotFound() {
						err = db.Table("video_rightsign").Create(&VideoSign{VideoID: v.ID, SignID: sign.ID}).Error
					}
				}
				if err != nil {
					cmd := exec.Command("/bin/bash", "-c", "echo '"+err.Error()+"' >> error.txt")
					cmd.Output()
					continue
				}
			}
		}
	}
}
