package controllers

import (
	"ccsl/middlewares"
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
	app.Handle("POST", "/", "CreateVideo", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("GET", "/{id: string}", "GetVideo", middlewares.CheckJWTToken)
	app.Handle("PUT", "/{id: string}", "UpdateVideo", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeleteVideo", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

// GetVideosList returns videos list with given parameters
func (c *VideoController) GetVideosList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "initial")
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
	leftSign := c.Context.URLParamDefault("leftSign", "")
	rightSign := c.Context.URLParamDefault("rightSign", "")
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
		data:    videos,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

func (c *VideoController) CreateVideo() {
	defer c.Context.Next()
	var form videoCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "VideoController::ParamsError", err)
		return
	}
	// PSQL - Create video in database.
	video := form.ConvertToModel()
	if err := c.VideoService.CreateVideo(video); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "VideoService::CreateVideo", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})

}

func (c *VideoController) GetVideo() {
	defer c.Context.Next()
	videoID := c.Context.Params().Get("id")
	video, err := c.VideoService.GetVideo(videoID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "VideoService::GetVideo", err)
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    video,
	})
}

func (c *VideoController) UpdateVideo() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	videoID := c.Context.Params().Get("id")
	var form videoUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "VideoController::ParamsError", err)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.VideoService.UpdateVideo(videoID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "VideoService::UpdateVideo", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

func (c *VideoController) DeleteVideo() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	videoID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.VideoService.DeleteVideo(videoID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "WordService::DeleteWord", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
