package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/utils"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	uuid "github.com/satori/go.uuid"
)

type FileController struct {
	Context iris.Context
}

const maxSize = 5 << 20 // 5MB

// BeforeActivation will register routes for controllers
func (c *FileController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("POST", "/", "UploadFile", middlewares.CheckJWTToken, iris.LimitRequestBodySize(maxSize+1<<20))
}

// UploadFile saves file and returns file name
func (c *FileController) UploadFile() {
	defer c.Context.Next()
	file, info, err := c.Context.FormFile("file")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UploadFile::ReadFile", err)
		return
	}
	defer file.Close()
	fileSuffix := path.Ext(info.Filename)
	filePrefix := uuid.NewV4()
	fileName := filePrefix.String() + strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
	filePath := configs.Conf.File.Dir + fileName
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UploadFile::SaveFile", err)
		return
	}
	defer out.Close()
	io.Copy(out, file)
	c.Context.JSON(iris.Map{
		message: success,
		data:    fileName,
	})
}
