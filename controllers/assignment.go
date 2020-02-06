package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
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

// GetAssignmentList returns assignment list
func (c *AssignmentController) GetAssignmentList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "assignments.title")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "AssignmentController::GetAssignmentList", errParams)
		return
	}
	courseID := c.Context.URLParamDefault("course_id", "")
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

// CreateAssignment POST /assignment
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

// UpdateAssignment PUT /assignment/{id:string}
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
