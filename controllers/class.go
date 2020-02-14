package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
    "ccsl/models"
    "ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// ClassController is for class CRUD
type ClassController struct {
	Context      iris.Context
	ClassService services.ClassInterface
}

// BeforeActivation will register routes for controllers
func (c *ClassController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetClassList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetClass")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleSuperUser, configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateClass")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateClass")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteClass")
}

// GetClassList GET /class
// >>>>> DOCS  <<<<<
// =================
// @Tags Class
// @Summary List classs
// @Description get classs list
// @Accept  json
// @Produce json
// @Router /class [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param name   	query string false 	"search name of class"
// @Param detail   	query string false 	"search detail name of class"
// @Param resource  query string false 	"search resource name of class"
// @Success 200 {object} controllers.GetClassListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *ClassController) GetClassList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "classes.name")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ClassController::GetClassList", errParams)
		return
	}
	name := c.Context.URLParamDefault("name", "")
	detail := c.Context.URLParamDefault("detail", "")
	resource := c.Context.URLParamDefault("resource", "")
	listParameters := utils.GetClassListParameters{
		GetListParameters: listParams,
		Name:              name,
		Details:           detail,
		Resources:         resource,
	}
	class, count, err := c.ClassService.GetClassesList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::GetClassList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    class,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// GetClassListResponse Response for GetClassList
type GetClassListResponse struct {
	GetListResponse
	Data []models.Class `json:"data"`
}

// CreateClass POST /class
// >>>>> DOCS  <<<<<
// =================
// @Tags Class
// @Summary Create class
// @Description create a new class
// @Accept  json
// @Produce json
// @Router 	/class 	[POST]
// @Param 	user 	body	 controllers.classCreateForm	  true	"create user"
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *ClassController) CreateClass() {
	defer c.Context.Next()
	var form classCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ClassController::CreateClass", errParams)
		return
	}
	// PSQL - Create class in database.
	class := form.ConvertToModel()
	if err := c.ClassService.CreateClass(class); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::CreateClass", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetClass GET /classs/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Class
// @Summary Get class
// @Description Get a class by id
// @Accept  json
// @Produce json
// @Router 	/class/{id} [GET]
// @Param 	id 		path	 string	    true	"class id" format(uuid)
// @Success 200 	{object} controllers.GetClassResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *ClassController) GetClass() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	classID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	class, err := c.ClassService.GetClass(classID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::GetClass", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    class,
	})
}

// GetClassResponse Response for get class
type GetClassResponse struct {
	SuccessResponse
	Data models.Class `json:"data"`
}

// UpdateClass PUT /class/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Class
// @Summary Update class
// @Description update a class by id
// @Accept  json
// @Produce json
// @Router 	/users/{id} [PUT]
// @Param 	id 		path	 string						  true	"class id" format(uuid)
// @Param 	user 	body	 controllers.classUpdateForm	  true	"updated class"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *ClassController) UpdateClass() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	classID := c.Context.Params().Get("id")
	var form classUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ClassController::UpdateClass", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.ClassService.UpdateClass(classID, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::UpdateClass", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteClass DELETE /class/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Class
// @Summary Delete class
// @Description delete a class by id
// @Accept  json
// @Produce json
// @Router 	/class/{id} [DELETE]
// @Param 	id 		path	 string						  true	"class id" format(uuid)
// @Success 204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *ClassController) DeleteClass() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	classID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.ClassService.DeleteClass(classID); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::DeleteClass", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
