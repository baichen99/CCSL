package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/utils"
	"errors"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// FileController is for uploading files
type FileController struct {
	Context iris.Context
}

const maxSize = 1024 << 20 // 1GB

// BeforeActivation will register routes for controllers
func (c *FileController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckToken, iris.LimitRequestBodySize(maxSize+1<<20))
	app.Handle(iris.MethodPost, "/", "UploadFile")
	// app.Handle("POST", "/replace", "ResizeAndReplaceScreen", middlewares.CheckJWTToken, iris.LimitRequestBodySize(maxSize+1<<20))
}

// UploadFile saves file and returns file name
func (c *FileController) UploadFile() {
	defer c.Context.Next()
	// dir is the subdirectory under public files directory
	dir := c.Context.URLParamDefault("dir", "")
	fileDirName := configs.Conf.File.Dir + dir
	fileDir, err := os.Stat(fileDirName)
	if err != nil || !fileDir.IsDir() {
		utils.SetError(c.Context, iris.StatusBadRequest, "UploadFile::DirNotExsist", errors.New("DirNotExsist"))
		return
	}
	// Read file from request form
	file, info, err := c.Context.FormFile("file")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "UploadFile::ReadFile", errParams)
		return
	}
	defer file.Close()
	fileSuffix := strings.ToLower(path.Ext(info.Filename))
	filePrefix := uuid.NewV4()
	// Rename file, avoid file name conflict
	// fileName = UUID + current timestamp + filename extension
	fileName := filePrefix.String() + strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
	filePath := path.Join(fileDirName, fileName)
	// write file to directory
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UploadFile::SaveFile", errParams)
		return
	}
	data := path.Join(dir, fileName)
	defer out.Close()
	io.Copy(out, file)
	c.Context.JSON(UploadFileResponse{
		success,
		data,
	})
}

type UploadFileResponse struct {
	Message string `json:"message" example:"success"`
	Data    string `json:"data" example:"dir/filename.mp4"`
}

// // ResizeAndReplaceScreen process video and return file name
// func (c *FileController) ResizeAndReplaceScreen() {
// 	// upload - save - resize - save - replace - save
// 	defer c.Context.Next()
// 	file, info, err := c.Context.FormFile("file")
// 	if err != nil {
// 		utils.SetResponseError(c.Context, iris.StatusBadRequest, "ResizeAndReplaceScreen::ReadFile", err)
// 		return
// 	}
// 	defer file.Close()
// 	fileSuffix := path.Ext(info.Filename)
// 	filePrefix := uuid.NewV4()

// 	srcFileName := filePrefix.String() + strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
// 	srcFilePath := os.TempDir() + srcFileName
// 	// save upload file to temp dir
// 	out, err := os.OpenFile(srcFilePath, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "ResizeAndReplaceScreen::SaveFile", err)
// 		return
// 	}
// 	defer out.Close()
// 	io.Copy(out, file)
// 	// resize
// 	resizeFileName := filePrefix.String() + strconv.FormatInt(time.Now().Unix(), 10) + ".avi"
// 	resizeFilePath := os.TempDir() + resizeFileName
// 	err = utils.ResizeVideo(srcFilePath, resizeFilePath, 640, 360)
// 	if err != nil {
// 		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "ResizeAndReplaceScreen::ResizeVideo", err)
// 		return
// 	}
// 	// replace background
// 	dstFileName := filePrefix.String() + strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
// 	dstFilePath := configs.Conf.File.Dir + dstFileName
// 	err = utils.Convert(resizeFilePath, dstFilePath, 230, 230, 230)
// 	if err != nil {
// 		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "ResizeAndReplaceScreen::ReplaceBG", err)
// 		return
// 	}
// 	c.Context.JSON(iris.Map{
// 		message: success,
// 		data:    dstFileName,
// 	})
// }
