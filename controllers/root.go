package controllers

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/kataras/iris/v12"
)

// RootController is the root controller for test
type RootController struct {
	Context iris.Context
}

// Get Controller: GET /
func (c *RootController) Get() {
	defer c.Context.Next()
	hello := c.Context.Tr("Hello")
	lang := c.Context.GetLocale().Language()
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	versionFile := path.Join(workPath, "configs", ".version")
	ver, _ := ioutil.ReadFile(versionFile)
	ip := c.Context.GetHeader("X-Real-IP")
	c.Context.JSON(iris.Map{
		message:   hello,
		language:  lang,
		version:   strings.Replace(string(ver), "\n", "", -1),
		ipAddress: ip,
	})
}
