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
	VideoService services.LexicalVideoInterface
}

// BeforeActivation will register routes for controllers
func (c *VideoController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetVideosList", middlewares.CheckJWTToken)
	app.Handle("POST", "/", "CreateVideo", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("GET", "/{id: string}", "GetVideo", middlewares.CheckJWTToken)
	app.Handle("PUT", "/{id: string}", "UpdateVideo", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeleteVideo", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

// GetVideosList GET /lexical/videos
func (c *VideoController) GetVideosList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "lexical_words.initial")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	wordID := c.Context.URLParamDefault("wordID", "")
	initial := c.Context.URLParamDefault("initial", "")
	chinese := c.Context.URLParamDefault("chinese", "")
	english := c.Context.URLParamDefault("english", "")
	pos := c.Context.URLParamDefault("pos", "")
	regionID := c.Context.URLParamDefault("regionID", "")
	gender := c.Context.URLParamDefault("gender", "")
	leftSignID := c.Context.URLParamDefault("leftSignID", "")
	rightSignID := c.Context.URLParamDefault("rightSignID", "")
	signID := c.Context.URLParamDefault("signID", "")
	constructWords := c.Context.URLParamDefault("constructWords", "")
	constructType := c.Context.URLParamDefault("constructType", "")
	performerID := c.Context.URLParamDefault("performerID", "")
	listParameters := utils.GetVideoListParameters{
		GetListParameters: listParams,
		WordID:            wordID,
		Initial:           initial,
		Chinese:           chinese,
		English:           english,
		Pos:               pos,
		RegionID:          regionID,
		Gender:            gender,
		LeftSignID:        leftSignID,
		RightSignID:       rightSignID,
		SignID:            signID,
		ConstructType:     constructType,
		ConstructWords:    constructWords,
		PerformerID:       performerID,
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

// CreateVideo POST /lexical/videos
func (c *VideoController) CreateVideo() {
	defer c.Context.Next()
	var form lexicalVideoCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "VideoController::ParamsError", err)
		return
	}
	// PSQL - Create video in database.
	video := form.ConvertToModel()
	leftSignsIds := form.LeftSignsID
	rightSignsIds := form.RightSignsID
	if err := c.VideoService.CreateVideo(video, leftSignsIds, rightSignsIds); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "VideoService::CreateVideo", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})

}

// GetVideo GET /lexical/videos/{id:string}
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

// UpdateVideo PUT /lexical/videos/{id:string}
func (c *VideoController) UpdateVideo() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	videoID := c.Context.Params().Get("id")
	var form lexicalVideoUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "VideoController::ParamsError", err)
		return
	}

	updateData := utils.MakeUpdateData(form)

	delete(updateData, "LeftSignsID")
	delete(updateData, "RightSignsID")

	// PSQL - Update of the given ID
	if err := c.VideoService.UpdateVideo(videoID, updateData, form.LeftSignsID, form.RightSignsID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "VideoService::UpdateVideo", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteVideo DELETE /lexical/videos/{id:string}
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
