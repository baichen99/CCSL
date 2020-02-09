package utils

import (
	"ccsl/configs"
	"ccsl/models"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

// ConnectPostgres : connect to postgresql database
func ConnectPostgres(app *iris.Application) *gorm.DB {
	var (
		host     = configs.Conf.Postgresql.Server
		port     = strconv.Itoa(configs.Conf.Postgresql.Port)
		username = configs.Conf.Postgresql.User
		password = configs.Conf.Postgresql.Password
		database = configs.Conf.Postgresql.Database
	)

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	pg, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		app.Logger().Errorf("Open PG Error: " + err.Error())
		panic(err)
	}
	authFileBuffer := fmt.Sprintf("%s:%s:%s:%s:%s\n", host, port, database, username, password)
	authFileName, err := os.UserHomeDir()
	if err != nil {
		app.Logger().Errorf("Get Home Dir Error: " + err.Error())
		panic(err)
	}
	authFileName += "/.pgpass"
	if err := ioutil.WriteFile(authFileName, []byte(authFileBuffer), 0600); err != nil {
		app.Logger().Errorf("Write Auth File Error: " + err.Error())
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
			Name:     "Admin User",
			Password: password,
			Roles:    []string{configs.RoleAdminUser},
			State:    "active",
		},
		models.User{
			Username: "user@ccsl.shu.edu.cn",
			Name:     "User",
			Password: password,
			Roles:    []string{configs.RoleDatabaseUser},
			State:    "active",
		},
		models.User{
			Username: "learner@ccsl.shu.edu.cn",
			Name:     "Learner",
			Password: password,
			Roles:    []string{configs.RoleStudent},
			State:    "active",
		},
		models.User{
			Username: "super@ccsl.shu.edu.cn",
			Name:     "Super User",
			Password: password,
			Roles:    []string{configs.RoleSuperUser},
			State:    "active",
		},
	}
	for _, user := range testUsers {
		pg.Create(&user)
	}
}
