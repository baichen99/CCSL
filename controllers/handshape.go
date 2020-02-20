package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// HandshapeController is for word CRUD
type HandshapeController struct {
	Context          iris.Context
	HandshapeService services.HandshapeInterface
}

// BeforeActivation register routes
func (c *HandshapeController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodGet, "/", "GetHandshapesList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetHandshape")
	app.Router().Use(middlewares.CheckUserRole([]string{configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateHandshape")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateHandshape")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteHandshape")
}

// GetHandshapesList GET /handshapes
func (c *HandshapeController) GetHandshapesList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "handshapes.name")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "HandshapeController::GetHandshapesList", errParams)
		return
	}
	listParameters := utils.GetHandshapeListParameters{
		GetListParameters: listParams,
		Name:              c.Context.URLParamDefault("name", ""),
	}
	handshapes, count, err := c.HandshapeService.GetHandshapesList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "HandshapeService::GetHandshapesList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    handshapes,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateHandshape POST /handshapes
func (c *HandshapeController) CreateHandshape() {
	defer c.Context.Next()
	var form HandshapeCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "HandshapeController::CreateHandshape", errParams)
		return
	}
	handshape := form.ConvertToModel()
	if err := c.HandshapeService.CreateHandshape(handshape); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "WordService::CreateHandshape", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetHandshape GET  /handshapes/{id:string}
func (c *HandshapeController) GetHandshape() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")
	handshape, err := c.HandshapeService.GetHandshape(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "HandshapeService::GetHandshape", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    handshape,
	})
}

// UpdateHandshape PUT /handshapes/{id:string}
func (c *HandshapeController) UpdateHandshape() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form HandshapeUpdateForm
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "HandshapeController::UpdateHandshape", errParams)
	}
	if err := c.HandshapeService.UpdateHandshape(id, utils.MakeUpdateData(form)); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "HandshapeService::UpdateHandshape", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteHandshape DELETE /handshapes/{id:string}
func (c *HandshapeController) DeleteHandshape() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.HandshapeService.DeleteHandshape(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "HandshapeService::DeleteHandshape", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
