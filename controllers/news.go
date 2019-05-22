package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// NewsController is for news CURD
type NewsController struct {
	Context iris.Context
}

// BeforeActivation will register routes for controllers
func (c *NewsController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetNewsList")
}

// GetNewsList returns news list
func (c *NewsController) GetNewsList() {
	defer c.Context.Next()
}
