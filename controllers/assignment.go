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

// AssignmentController is for assignment CRUD
type AssignmentController struct {
	Context           iris.Context
	AssignmentService services.AssignmentInterface
}

// BeforeActivation will register routes for controllers
func (c *AssignmentController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetAssignmentList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetAssignment")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleSuperUser, configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateAssignment")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateAssignment")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteAssignment")
}

// GetAssignmentList GET /assignment
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignment
// @Summary List assignment
// @Description get assignment list
// @Accept  json
// @Produce json
// @Router /assignment [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param course_id query string false 	"filter course_id of assignment"
// @Param type   	query string false 	"filter type of assignment"
// @Param title     query string false 	"search title of assignment"
// @Param content   query string false 	"search content of assignment"
// @Success 200 {object} controllers.GetAssignmentListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *AssignmentController) GetAssignmentList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "assignments.title")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "AssignmentController::GetAssignmentList", errParams)
		return
	}
	courseID := c.Context.URLParamDefault("courseID", "")
	title := c.Context.URLParamDefault("title", "")
	atype := c.Context.URLParamDefault("type", "")
	content := c.Context.URLParamDefault("content", "")
	listParameters := utils.GetAssignmentListParameters{
		GetListParameters: listParams,
		CourseID:          courseID,
		Title:             title,
		Type:              atype,
		Content:           content,
	}
	assignments, count, err := c.AssignmentService.GetAssignmentsList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::GetAssignmentList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    assignments,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// GetAssignmentListResponse Response for GetAssignmentList
type GetAssignmentListResponse struct {
	GetListResponse
	Data []models.Assignment `json:"data"`
}

// CreateAssignment POST /assignment
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignment
// @Summary Create assignment
// @Description create a new assignment
// @Accept  json
// @Produce json
// @Router 	/assignment 	[POST]
// @Param 	assignment 	body	 controllers.assignmentCreateForm	  true	"create assignment"
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) CreateAssignment() {
	defer c.Context.Next()
	var form assignmentCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "AssignmentController::CreateAssignment", errParams)
		return
	}
	// PSQL - Create assignment in database.
	assignment := form.ConvertToModel()

	if err := c.AssignmentService.CreateAssignment(assignment); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::CreateAssignment", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetAssignment GET /assignment/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignment
// @Summary Get assignment
// @Description Get a assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignment/{id} [GET]
// @Param 	id 		path	 string	    true	"assignment id" format(uuid)
// @Success 200 	{object} controllers.GetAssignmentResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) GetAssignment() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	assignmentID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	assignment, err := c.AssignmentService.GetAssignment(assignmentID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::GetAssignment", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    assignment,
	})
}

// GetAssignmentResponse Response for get assignment
type GetAssignmentResponse struct {
	SuccessResponse
	Data models.Assignment `json:"data"`
}

// UpdateAssignment PUT /assignment/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignment
// @Summary Update assignment
// @Description update a assignment by id
// @Accept  json
// @Produce json
// @Router 	/users/{id} [PUT]
// @Param 	id 		path	 string						  true	"assignment id" format(uuid)
// @Param 	user 	body	 controllers.assignmentUpdateForm	  true	"updated assignment"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) UpdateAssignment() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	assignmentID := c.Context.Params().Get("id")
	var form assignmentUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "AssignmentController::UpdateAssignment", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.AssignmentService.UpdateAssignment(assignmentID, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::UpdateAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteAssignment DELETE /assignment/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignment
// @Summary Delete assignment
// @Description delete a assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignment/{id} [DELETE]
// @Param 	id 		path	 string						  true	"assignment id" format(uuid)
// @Success 204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) DeleteAssignment() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	assignmentID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.AssignmentService.DeleteAssignment(assignmentID); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::DeleteAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
