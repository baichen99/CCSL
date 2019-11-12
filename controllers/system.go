package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"fmt"
	"os/exec"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// SystemController is for system monitor and maintenance
type SystemController struct {
	Context       iris.Context
	SystemService services.SystemInterface
}

// BeforeActivation register routes
func (c *SystemController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("POST", "/error", "JsErrorLogger")
	app.Router().Use(middlewares.CheckJWTToken, middlewares.CheckSuper)
	app.Handle("GET", "/error", "GetJsErrorList")
	app.Handle("GET", "/login", "GetLoginHistoryList")
	app.Handle("GET", "/dump", "DumpDatabase")
}

// JsErrorLogger POST /system/error
func (c *SystemController) JsErrorLogger() {
	defer c.Context.Next()
	var jsError models.JsError
	if err := utils.ReadValidateForm(c.Context, &jsError); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "SystemController::JsErrorLogger", err)
		return
	}
	if err := c.SystemService.CreateJsError(jsError); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "SystemService::CreateJsError", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetJsErrorList GET /systems/error
func (c *SystemController) GetJsErrorList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "js_errors.created_at")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	listParameters := utils.GetJsErrorListParameters{
		GetListParameters: listParams,
	}
	errors, count, err := c.SystemService.GetJsErrorList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "SystemService::GetJsErrorList", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    errors,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// GetLoginHistoryList GET /systems/login
func (c *SystemController) GetLoginHistoryList() {
	defer c.Context.Next()
}

// DumpDatabase POST /systems/dump
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
