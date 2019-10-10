package controllers

import (
	"io/ioutil"
	"os"
	"strings"

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
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	versionFile := workPath + "/configs/.version"
	ver, _ := ioutil.ReadFile(versionFile)
	c.Context.JSON(iris.Map{
		message:  hello,
		language: lang,
		version:  strings.Replace(string(ver), "\n", "", -1),
		"IP":     c.Context.GetHeader("X-Real-IP"),
	})
}
