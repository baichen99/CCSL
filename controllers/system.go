package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"fmt"
	"os/exec"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// SystemController is for system infomation CRUD
type SystemController struct {
	Context       iris.Context
	SystemService services.SystemInterface
	UserService   services.UserInterface
}

// BeforeActivation register routes
func (c *SystemController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodPost, "/error", "JsErrorLogger")
	app.Handle(iris.MethodGet, "/cities", "GetCitiesList")
	app.Handle(iris.MethodGet, "/info/{key: string}", "GetAppInfo")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleAdminUser}))
	app.Handle(iris.MethodPut, "/info/{key: string}", "UpdateAppInfo")
	app.Router().Use(middlewares.CheckUserRole([]string{configs.RoleSuperUser}))
	app.Handle(iris.MethodGet, "/error", "GetJsErrorList")
	app.Handle(iris.MethodDelete, "/error/{id: string}", "DeleteJsError")
	app.Handle(iris.MethodGet, "/login", "GetLoginHistoryList")
	app.Handle(iris.MethodGet, "/dump", "DumpDatabase")
}

// GetAppInfo GET /systems/info/xxx
func (c *SystemController) GetAppInfo() {
	defer c.Context.Next()
	key := c.Context.Params().Get("key")
	info, err := c.SystemService.GetAppInfo(key)
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SystemService::GetAppInfo", errParams)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    info.Data,
	})
}

// UpdateAppInfo PUT /systems/info/xxx
func (c *SystemController) UpdateAppInfo() {
	defer c.Context.Next()
	key := c.Context.Params().Get("key")
	var form InfoUpdateForm
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SystemController::UpdateAppInfo", errParams)
		return
	}

	updateData := utils.MakeUpdateData(form)

	err := c.SystemService.UpdateAppInfo(key, updateData)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SystemService::UpdateAppInfo", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// JsErrorLogger POST /systems/error
func (c *SystemController) JsErrorLogger() {
	defer c.Context.Next()
	var jsError models.JsError
	if err := utils.ReadValidateForm(c.Context, &jsError); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SystemController::JsErrorLogger", errParams)
		return
	}
	if err := c.SystemService.CreateJsError(jsError); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SystemService::CreateJsError", errSQL)
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
		utils.SetError(c.Context, iris.StatusBadRequest, "SystemController::GetJsErrorList", errParams)
		return
	}
	listParameters := utils.GetJsErrorListParameters{
		GetListParameters: listParams,
	}
	errors, count, err := c.SystemService.GetJsErrorList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SystemService::GetJsErrorList", errSQL)
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

// DeleteJsError DELETE /systems/error/{id:string}
func (c *SystemController) DeleteJsError() {
	defer c.Context.Next()

	id := c.Context.Params().Get("id")

	if err := c.SystemService.DeleteJsError(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SystemService::DeleteJsError", errSQL)
		return
	}

	c.Context.StatusCode(iris.StatusNoContent)
}

// GetLoginHistoryList GET /systems/login
func (c *SystemController) GetLoginHistoryList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "login_histories.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SystemController::GetLoginHistoryList", errParams)
		return
	}
	userID := c.Context.URLParamDefault("userID", "")
	status := c.Context.URLParamDefault("status", "")
	ip := c.Context.URLParamDefault("ip", "")
	listParameters := utils.GetUserLoginListParameters{
		GetListParameters: listParams,
		UserID:            userID,
		Status:            status,
		IP:                ip,
	}
	list, count, err := c.UserService.GetLoginHistoryList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::GetLoginHistoryList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    list,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// DumpDatabase POST /systems/dump
func (c *SystemController) DumpDatabase() {
	defer c.Context.Next()
	const PGDumpCmd = "pg_dump"
	options := []string{
		"-Fc",
		fmt.Sprintf(`-d%v`, configs.Conf.Postgresql.Database),
		fmt.Sprintf(`-h%v`, configs.Conf.Postgresql.Server),
		fmt.Sprintf(`-p%v`, configs.Conf.Postgresql.Port),
		fmt.Sprintf(`-U%v`, configs.Conf.Postgresql.User),
	}
	out, err := exec.Command(PGDumpCmd, options...).Output()
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SystemController::DumpDB", errSQL)
		return
	}
	c.Context.Binary(out)
}

// GetCitiesList GET /systems/cities
func (c *SystemController) GetCitiesList() {
	defer c.Context.Next()
	cities, err := c.SystemService.GetCitiesList()
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SystemService::GetCitiesList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    cities,
	})
}
