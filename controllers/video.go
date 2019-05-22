package controllers

import (
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// VideoController is for video CURD
type VideoController struct {
	Context      iris.Context
	VideoService services.VideoInterface
}

// BeforeActivation will register routes for controllers
func (c *VideoController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetVideosList")
}

// GetVideosList returns videos list with given parameters
func (c *VideoController) GetVideosList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "initial, region")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	initial := c.Context.URLParamDefault("initial", "")
	chinese := c.Context.URLParamDefault("chinese", "")
	english := c.Context.URLParamDefault("english", "")
	wordType := c.Context.URLParamDefault("type", "")
	name := c.Context.URLParamDefault("name", "")
	region := c.Context.URLParamDefault("region", "")
	gender := c.Context.URLParamDefault("gender", "")
	leftSign := c.Context.URLParamDefault("left", "")
	rightSign := c.Context.URLParamDefault("right", "")
	constructWords := c.Context.URLParamDefault("constructWords", "")
	constructType := c.Context.URLParamDefault("constructType", "")
	performerID := c.Context.URLParamDefault("performer", "")
	wordID := c.Context.URLParamDefault("word", "")
	listParameters := utils.GetVideoListParameters{
		GetListParameters: listParams,
		Initial:           initial,
		Chinese:           chinese,
		English:           english,
		Type:              wordType,
		Name:              name,
		Region:            region,
		Gender:            gender,
		LeftSign:          leftSign,
		RightSign:         rightSign,
		ConstructType:     constructType,
		ConstructWords:    constructWords,
		PerformerID:       performerID,
		WordID:            wordID,
	}
	videos, count, err := c.VideoService.GetVideosList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "VideoService::GetVideosList", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data: iris.Map{
			"videos": videos,
			pageKey:  listParams.Page,
			limitKey: listParams.Limit,
			totalKey: count,
		},
	})
}
