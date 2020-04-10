package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"
	"errors"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// LexicalVideoController is for video CRUD
type LexicalVideoController struct {
	Context             iris.Context
	LexicalVideoService services.LexicalVideoInterface
}

// BeforeActivation will register routes for controllers
func (c *LexicalVideoController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodGet, "/", "GetLexicalVideosList", middlewares.CheckRateLimit(3))
	app.Handle(iris.MethodGet, "/{id: string}", "GetLexicalVideo", middlewares.CheckRateLimit(3))
	app.Router().Use(middlewares.CheckUserRole([]string{configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateLexicalVideo")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateLexicalVideo")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteLexicalVideo")
}

// GetLexicalVideosList GET /lexical/videos
// >>>>> DOCS  <<<<<
// =================
// @Tags Lexical Videos
// @Summary List videos for lexical database
// @Description get videos list
// @Accept  json
// @Produce json
// @Router /lexical/videos [GET]
// @Param page 		        query int    false "select from page" 			mininum(1)
// @Param limit 	        query int    false "limit number" 				mininum(0)
// @Param order		        query string false "order by field"
// @Param orderBy 	        query string false "order by asc or desc" 		enums(asc, desc)
// @Param lexiconID         query string false "filter by lexicon ID"
// @Param initial 	        query string false "filter initial"
// @Param chinese 	        query string false "search by Chinese"
// @Param english 	        query string false "search by English"
// @Param pos	 	        query string false "filter by part of speech"
// @Param regionID	        query int	 false "filter by region ID"
// @Param gender	        query string false "filter by gender" 			enums(M, F)
// @Param leftHandshapeID   query string false "filter by left handshape ID"
// @Param rightHandshapeID  query string false "filter by right handshape ID"
// @Param handshapeID	    query string false "filter by handshape ID"
// @Param morpheme 	        query string false "search by morpheme"
// @Param wordFormation     query string false "filter by word formation" 	enums(simple, compound)
// @Param performerID	    query string false "filter by performer ID"
// @Success 200 {object} controllers.GetLexicalVideosListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *LexicalVideoController) GetLexicalVideosList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "lexicons.initial, lexicons.id, performers.region_id, performers.gender")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "LexicalVideoController::GetLexicalVideosList", errors.New("ParamsError"))
		return
	}
	listParameters := utils.GetLexicalVideoListParameters{
		GetListParameters: listParams,
		LexiconID:         c.Context.URLParamDefault("lexiconID", ""),
		Initial:           c.Context.URLParamDefault("initial", ""),
		Chinese:           c.Context.URLParamDefault("chinese", ""),
		English:           c.Context.URLParamDefault("english", ""),
		Pos:               c.Context.URLParamDefault("pos", ""),
		RegionID:          c.Context.URLParamDefault("regionID", ""),
		Gender:            c.Context.URLParamDefault("gender", ""),
		LeftHandshapeID:   c.Context.URLParamDefault("leftHandshapeID", ""),
		RightHandshapeID:  c.Context.URLParamDefault("rightHandshapeID", ""),
		HandshapeID:       c.Context.URLParamDefault("handshapeID", ""),
		WordFormation:     c.Context.URLParamDefault("wordFormation", ""),
		Morpheme:          c.Context.URLParamDefault("morpheme", ""),
		PerformerID:       c.Context.URLParamDefault("performerID", ""),
	}
	videos, count, err := c.LexicalVideoService.GetLexicalVideosList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "LexicalVideoService::GetLexicalVideosList", errors.New("SqlError"))
		return
	}

	c.Context.JSON(GetLexicalVideosListResponse{
		GetListResponse{
			success,
			listParams.Page,
			listParams.Limit,
			count,
		},
		videos,
	})
}

// CreateLexicalVideo POST /lexical/videos
func (c *LexicalVideoController) CreateLexicalVideo() {
	defer c.Context.Next()
	var form LexicalVideoCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "LexicalVideoController::CreateLexicalVideo", errParams)
		return
	}
	// PSQL - Create video in database.
	video := form.ConvertToModel()

	if err := c.LexicalVideoService.CreateLexicalVideo(video); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "LexicalVideoService::CreateLexicalVideo", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})

}

// GetLexicalVideo GET /lexical/videos/{id:string}
func (c *LexicalVideoController) GetLexicalVideo() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")
	video, err := c.LexicalVideoService.GetLexicalVideo(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "LexicalVideoService::GetLexicalVideo", errSQL)
	}
	c.Context.JSON(GetLexicalVideoResponse{
		success,
		video,
	})
}

// UpdateLexicalVideo PUT /lexical/videos/{id:string}
func (c *LexicalVideoController) UpdateLexicalVideo() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form LexicalVideoUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "LexicalVideoController::UpdateLexicalVideo", errParams)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.LexicalVideoService.UpdateLexicalVideo(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "LexicalVideoService::UpdateLexicalVideo", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteLexicalVideo DELETE /lexical/videos/{id:string}
func (c *LexicalVideoController) DeleteLexicalVideo() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.LexicalVideoService.DeleteLexicalVideo(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "WordService::DeleteWord", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
