package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// SignController is for word CRUD
type SignController struct {
	Context     iris.Context
	SignService services.SignInterface
}

// BeforeActivation register routes
func (c *SignController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckJWTToken)
	app.Handle("GET", "/", "GetSignsList")
	app.Handle("GET", "/{id: string}", "GetSign")
	app.Router().Use(middlewares.CheckAdmin)
	app.Handle("POST", "/", "CreateSign")
	app.Handle("PUT", "/{id: string}", "UpdateSign")
	app.Handle("DELETE", "/{id: string}", "DeleteSign")
}

// GetSignsList GET /signs
func (c *SignController) GetSignsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "signs.name")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SignController::GetSignsList", errParams)
		return
	}
	name := c.Context.URLParamDefault("name", "")
	listParameters := utils.GetSignListParameters{
		GetListParameters: listParams,
		Name:              name,
	}
	signs, count, err := c.SignService.GetSignList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SignService::GetSignsList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    signs,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateSign POST /signs
func (c *SignController) CreateSign() {
	defer c.Context.Next()
	var form signCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SignController::CreateSign", errParams)
		return
	}
	sign := form.ConvertToModel()
	if err := c.SignService.CreateSign(sign); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "WordService::CreateSign", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetSign GET  /signs/{id:string}
func (c *SignController) GetSign() {
	defer c.Context.Next()
	signID := c.Context.Params().Get("id")
	sign, err := c.SignService.GetSign(signID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SignService::GetSign", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    sign,
	})
}

// UpdateSign PUT /sings/{id:string}
func (c *SignController) UpdateSign() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	signID := c.Context.Params().Get("id")
	var form signUpdateForm
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SignController::UpdateSign", errParams)
	}
	if err := c.SignService.UpdateSign(signID, utils.MakeUpdateData(form)); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SignService::UpdateSign", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteSign DELETE /signs/{id:string}
func (c *SignController) DeleteSign() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	signID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.SignService.DeleteSign(signID); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SignService::DeleteSign", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
