package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/utils"
	"fmt"
	"os/exec"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// SystemController is for system monitor and maintenance
type SystemController struct {
	Context iris.Context
}

// BeforeActivation register routes
func (c *SystemController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckJWTToken, middlewares.CheckSuper)
	app.Handle("GET", "/login", "GetLoginHistoryList")
	app.Handle("GET", "/dump", "DumpDatabase")
}

// GetLoginHistoryList GET /system/login
func (c *SystemController) GetLoginHistoryList() {
	defer c.Context.Next()
}

// DumpDatabase POST /system/login
func (c *SystemController) DumpDatabase() {
	defer c.Context.Next()
	const PGDumpCmd = "pg_dump"
	options := []string{"-Fc"}
	options = append(options, fmt.Sprintf(`-d%v`, configs.Conf.Postgresql.Database))
	options = append(options, fmt.Sprintf(`-h%v`, configs.Conf.Postgresql.Server))
	options = append(options, fmt.Sprintf(`-p%v`, configs.Conf.Postgresql.Port))
	options = append(options, fmt.Sprintf(`-U%v`, configs.Conf.Postgresql.User))
	out, err := exec.Command(PGDumpCmd, options...).Output()
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "SystemController::DumpDB", err)
		return
	}
	c.Context.Binary(out)
}
