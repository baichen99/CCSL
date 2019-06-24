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

// InitSuper creates super admin users
func InitSuper(pg *gorm.DB) {
	superUsers := []models.User{
		models.User{
			Username: "16121041",
			UserType: "super",
		},
		models.User{
			Username: "15121709",
			UserType: "super",
		},
	}
	for _, superUser := range superUsers {
		pg.Create(&superUser)
	}
}

// InitTestUser creates test users
func InitTestUser(pg *gorm.DB) {
	password, _ := HashPassword("P@ssw0rd")
	testUsers := []models.User{
		models.User{
			Username: "adrianduan@icloud.com",
			Password: password,
			UserType: "admin",
		},
		models.User{
			Username: "474558417@qq.com",
			Password: password,
			UserType: "user",
		},
		models.User{
			Username: "adrwefian@qq.com",
			Password: password,
			UserType: "user",
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
