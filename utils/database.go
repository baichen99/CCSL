package utils

import (
	"ccsl/configs"
	"ccsl/models"
	"strconv"

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
