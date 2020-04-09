package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// ReplyController is fro Reply CRUD
type ReplyController struct {
	Context      iris.Context
	ReplyService services.ReplyInterface
}

// BeforeActivation will register routes for controllers
func (c *ReplyController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetReplyList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetReply")
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodPost, "/", "CreateReply")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateReply")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteReply")
}

// GetReplyList GET /replies
// >>>>> DOCS  <<<<<
// =================
// @Tags Replies
// @Summary List Replies
// @Description get replies list
// @Accept json
// @Produce json
// @Router /replies [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param message 	query string false 	"search message of notification"
// @Success 200 {object}    controllers.GetReplyListResponse
// @Failure 400 {object}    controllers.ErrorResponse
// @Failure 422 {object}    controllers.ErrorResponse
// =================
func (c *ReplyController) GetReplyList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "replies.id")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ReplyController::GetReplyList", errParams)
		return
	}
	listParameters := utils.GetReplyListParameters{
		GetListParameters: listParams,
		PostID:            c.Context.URLParamDefault("PostID", ""),
		CreatorID:         c.Context.URLParamDefault("creatorID", ""),
	}
	replies, count, err := c.ReplyService.GetReplyList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ReplyService::GetReplyList", errSQL)
	}
	c.Context.JSON(GetReplyListResponse{
		GetListResponse{
			Message: success,
			Page:    listParameters.Page,
			Limit:   listParameters.Limit,
			Total:   count,
		},
		replies,
	})
}

// GetReply GET /replies/{id: string}
// >>>>> DOCS <<<<<
// =================
// @Tags Replies
// @Summary Get reply
// @Description Get a reply by id
// @Accept  json
// @Produce json
// @Router /replies/{id} [GET]
// @Param   id  path    string  true    "reply id" format(uuid)
// @Success 200 {object}    controllers.GetReplyResponse
// @Failure 422 {object}    controllers.ErrorResponse
// =================
func (c *ReplyController) GetReply() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")

	reply, err := c.ReplyService.GetReply(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ReplyService::GetReply", errSQL)
		return
	}

	c.Context.JSON(GetReplyResponse{
		success,
		reply,
	})

}

// CreateReply POST /replies
// >>>>> DOCS  <<<<<
// =================
// @Tags Replies
// @Summary Create reply
// @Description create a new reply
// @Accept json
// @Produce json
// @Router /replies [POST]
// @Param   reply   body    controllers.ReplyCreateForm true    "create reply"
// @Success 201     {object}    controllers.SuccessResponse
// @Failure 400     {object}    controllers.ErrorResponse
// @Failure 422     {object}    controllers.ErrorResponse
// =================
func (c *ReplyController) CreateReply() {
	defer c.Context.Next()
	var form ReplyCreateForm
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ReplyController:CreateReply", errParams)
		return
	}
	form.Content = middlewares.FilterUserInput(form.Content)
	reply := form.ConvertToModel()
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	reply.CreatorID = tokenID

	if err := c.ReplyService.CreateReply(reply); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ReplyService::CreateReply", errSQL)
		return
	}

	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})
}

// UpdateReply PUT /replies/{id: string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Replies
// @Summary Update reply
// @Description update a reply by id
// @Accept json
// @Produce json
// @Router  /replies/{id} [PUT]
// @Param   id  path    string      true    "reply id" form(uuid)
// @Success 204
// @Failure 400     {object}    controllers.ErrorResponse
// @Failure 401     {object}    controllers.ErrorResponse
// @Failure 422     {object}    controllers.ErrorResponse
// =================
func (c *ReplyController) UpdateReply() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")
	var form ReplyUpdateForm

	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ReplyController::UpdateReply", errParams)
		return
	}
	reply, _ := c.ReplyService.GetReply(id)

	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	if tokenID != reply.CreatorID {
		utils.SetError(c.Context, iris.StatusUnauthorized, "ReplyController::UpdateReply", errAuth)
		return
	}

	updateData := utils.MakeUpdateData(form)
	if err := c.ReplyService.UpdateReply(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ReplyService::UpdateReply", errSQL)
		return
	}
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteReply DELETE /replies/{id: string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Replies
// @Summary Delete reply
// @Description delete a post by id
// @Accept  json
// @Produce json
// @Router  /replies/{id} [DELETE]
// @Param   id      path        string      true "reply id" format(uuid)
// @Success 204
// @Failure 401     {object}    controllers.ErrorResponse
// @Failure 422     {object}    controllers.ErrorResponse
// =================
func (c *ReplyController) DeleteReply() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")

	reply, _ := c.ReplyService.GetReply(id)

	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)

	if !(tokenID == reply.CreatorID || middlewares.HasPermisson(tokenRole, configs.RoleSuperUser)) {
		utils.SetError(c.Context, iris.StatusUnauthorized, "ReplyController::DeleteReply", errAuth)
		return
	}

	if err := c.ReplyService.DeleteReply(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ReplyService::DeleteReply", errSQL)
		return
	}

	c.Context.StatusCode(iris.StatusNoContent)
}
