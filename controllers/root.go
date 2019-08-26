package controllers

import (
	"github.com/kataras/iris"
)

// RootController is the root controller for test
type RootController struct {
	Context iris.Context
}

// Get Controller: GET /
func (c *RootController) Get() {
	defer c.Context.Next()
	hello := c.Context.Translate("Hello")
	lang := c.Context.Values().GetString(c.Context.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())
	c.Context.JSON(iris.Map{
		message:  hello,
		language: lang,
	})
}
