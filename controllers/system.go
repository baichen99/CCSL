package controllers

import (
	"ccsl/middlewares"

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
}

// GetLoginHistoryList GET /system/login
func (c *SystemController) GetLoginHistoryList() {
	defer c.Context.Next()
}
