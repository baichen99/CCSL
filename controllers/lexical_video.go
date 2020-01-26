package controllers

import (
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"errors"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// VideoController is for video CRUD
type VideoController struct {
	Context      iris.Context
	VideoService services.LexicalVideoInterface
}

// BeforeActivation will register routes for controllers
func (c *VideoController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckJWTToken)
	app.Handle("GET", "/", "GetVideosList", middlewares.CheckRateLimit(3))
	app.Handle("GET", "/{id: string}", "GetVideo", middlewares.CheckRateLimit(3))
	app.Router().Use(middlewares.CheckAdmin)
	app.Handle("POST", "/", "CreateVideo")
	app.Handle("PUT", "/{id: string}", "UpdateVideo")
	app.Handle("DELETE", "/{id: string}", "DeleteVideo")
}

// GetVideosList GET /lexical/videos
// >>>>> DOCS  <<<<<
// =================
// @Tags Lexical Videos
// @Summary List videos for lexical database
// @Description get videos list
// @Accept  json
// @Produce json
// @Router /lexical/videos [GET]
// @Param page 		    query int    false "select from page" 			mininum(1)
// @Param limit 	    query int    false "limit number" 				mininum(0)
// @Param order		    query string false "order by field"
// @Param orderBy 	    query string false "order by asc or desc" 		enums(asc, desc)
// @Param lexiconID     query string false "filter by lexicon ID"
// @Param initial 	    query string false "filter initial"
// @Param chinese 	    query string false "search by Chinese"
// @Param english 	    query string false "search by English"
// @Param pos	 	    query string false "filter by part of speech"
// @Param regionID	    query int	 false "filter by region ID"
// @Param gender	    query string false "filter by gender" 			enums(M, F)
// @Param leftSignID    query string false "filter by left sign ID"
// @Param rightSignID   query string false "filter by right sign ID"
// @Param signID	    query string false "filter by sign ID"
// @Param morpheme 	    query string false "search by morpheme"
// @Param wordFormation query string false "filter by word formation" 	enums(simple, compound)
// @Param performerID	query string false "filter by performer ID"
// @Success 200 {object} controllers.GetVideosListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *VideoController) GetVideosList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "lexicons.initial, lexicons.id, performers.region_id, performers.gender")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "VideoController::GetVideosList", errors.New("ParamsError"))
		return
	}
	lexiconID := c.Context.URLParamDefault("lexiconID", "")
	initial := c.Context.URLParamDefault("initial", "")
	chinese := c.Context.URLParamDefault("chinese", "")
	english := c.Context.URLParamDefault("english", "")
	pos := c.Context.URLParamDefault("pos", "")
	regionID := c.Context.URLParamDefault("regionID", "")
	gender := c.Context.URLParamDefault("gender", "")
	leftSignID := c.Context.URLParamDefault("leftSignID", "")
	rightSignID := c.Context.URLParamDefault("rightSignID", "")
	signID := c.Context.URLParamDefault("signID", "")
	morpheme := c.Context.URLParamDefault("morpheme", "")
	wordFormation := c.Context.URLParamDefault("wordFormation", "")
	performerID := c.Context.URLParamDefault("performerID", "")
	listParameters := utils.GetVideoListParameters{
		GetListParameters: listParams,
		LexiconID:         lexiconID,
		Initial:           initial,
		Chinese:           chinese,
		English:           english,
		Pos:               pos,
		RegionID:          regionID,
		Gender:            gender,
		LeftSignID:        leftSignID,
		RightSignID:       rightSignID,
		SignID:            signID,
		WordFormation:     wordFormation,
		Morpheme:          morpheme,
		PerformerID:       performerID,
	}
	videos, count, err := c.VideoService.GetVideosList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "VideoService::GetVideosList", errors.New("SqlError"))
		return
	}

	c.Context.JSON(iris.Map{
		message: success,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
		data:    videos,
	})
}

// GetVideosListResponse Response for GetVideosList
type GetVideosListResponse struct {
	GetListResponse
	Data []models.LexicalVideo `json:"data"`
}

// CreateVideo POST /lexical/videos
func (c *VideoController) CreateVideo() {
	defer c.Context.Next()
	var form lexicalVideoCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "VideoController::CreateVideo", errParams)
		return
	}
	// PSQL - Create video in database.
	video := form.ConvertToModel()

	if err := c.VideoService.CreateVideo(video); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "VideoService::CreateVideo", errSQL)
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
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "VideoService::GetVideo", errSQL)
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
		utils.SetError(c.Context, iris.StatusBadRequest, "VideoController::UpdateVideo", errParams)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.VideoService.UpdateVideo(videoID, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "VideoService::UpdateVideo", errSQL)
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
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "WordService::DeleteWord", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
