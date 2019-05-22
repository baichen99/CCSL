package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// PerformerController is for performers CURD
type PerformerController struct {
	Context          iris.Context
	PerformerService services.PerformerInterface
}

// BeforeActivation will register routes for controllers
func (c *PerformerController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetPerformersList", middlewares.CheckJWTToken)
	app.Handle("POST", "/", "CreatePerformer", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("GET", "/{id: string}", "GetPerformer", middlewares.CheckJWTToken)
	app.Handle("PUT", "/{id: string}", "UpdatePerformer", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeletePerformer", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

// GetPerformersList GET /performers
func (c *PerformerController) GetPerformersList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "region")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	name := c.Context.URLParamDefault("name", "")
	region := c.Context.URLParamDefault("region", "")
	gender := c.Context.URLParamDefault("gender", "")
	listParameters := utils.GetPerformerListParameters{
		GetListParameters: listParams,
		Name:              name,
		Region:            region,
		Gender:            gender,
	}
	performers, count, err := c.PerformerService.GetPerformersList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "PerformerService::GetPerformersList", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data: iris.Map{
			"performers": performers,
			pageKey:      listParams.Page,
			limitKey:     listParams.Limit,
			totalKey:     count,
		},
	})
}

// CreatePerformer POST /performers
func (c *PerformerController) CreatePerformer() {
	defer c.Context.Next()
	var form performerCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, paramsKey, err)
		return
	}
	// PSQL - Create performer in database.
	performer := form.ConvertToModel()
	if err := c.PerformerService.CreatePerformer(performer); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "PerformerService::CreatePerformer", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetPerformer GET /performers/{id:string}
func (c *PerformerController) GetPerformer() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	performerID := c.Context.Params().Get("id")

	// PSQL - Looking for specified performer via the ID.
	performer, err := c.PerformerService.GetPerformer(performerID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "PerformerService::GetPerformer", err)
		return
	}

	// Returning information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    performer,
	})
}

// UpdatePerformer PUT /performers/{id:string}
func (c *PerformerController) UpdatePerformer() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	performerID := c.Context.Params().Get("id")
	var form performerUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, paramsKey, err)
		return
	}
	updateData := form.ConvertToModel()
	// PSQL - Looking for specified performer via the ID.

	if err := c.PerformerService.UpdatePerformer(performerID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "PerformerService::UpdatePerformer", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeletePerformer DELETE /performers/{id:string}
func (c *PerformerController) DeletePerformer() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	performerID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.PerformerService.DeletePerformer(performerID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "PerformerService::DeletePerformer", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
