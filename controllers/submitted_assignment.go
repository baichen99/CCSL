package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// SubmittedAssignmentController is for submitted_assignment CRUD
type SubmittedAssignmentController struct {
	Context                    iris.Context
	SubmittedAssignmentService services.SubmittedAssignmentInterface
}

// BeforeActivation will register routes for controllers
func (c *SubmittedAssignmentController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetSubmittedAssignmentList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetSubmittedAssignment")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleStudent}))
	app.Handle(iris.MethodPost, "/", "CreateSubmittedAssignment")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteSubmittedAssignment")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleTeacher, configs.RoleAdminUser}))
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateSubmittedAssignment")

}

// GetSubmittedAssignmentList returns submitted_assignment list
func (c *SubmittedAssignmentController) GetSubmittedAssignmentList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "submitted_assignments.assignment_id")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::GetSubmittedAssignmentList", errParams)
		return
	}
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	assignmentID := c.Context.URLParamDefault("assignment_id", "")
	graderID := c.Context.URLParamDefault("grader_id", "")
	answer := c.Context.URLParamDefault("answer", "")
	comment := c.Context.URLParamDefault("comment", "")
	listParameters := utils.GetSubmittedAssignmentsListParameters{
		GetListParameters: listParams,
		CreatorID:         tokenUser,
		AssignmentID:      assignmentID,
		GraderID:          graderID,
		Answer:            answer,
		Comment:           comment,
	}
	submittedAssignment, count, err := c.SubmittedAssignmentService.GetSubmittedAssignmentsList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::GetSubmittedAssignmentList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    submittedAssignment,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateSubmittedAssignment POST /submitted_assignment
func (c *SubmittedAssignmentController) CreateSubmittedAssignment() {
	defer c.Context.Next()
	var form submmittedAssignmentCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::CreateSubmittedAssignment", errParams)
		return
	}
	// PSQL - Create submitted_assignment in database.
	submittedAssignment := form.ConvertToModel()

	// Check if current student select the course
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	if tokenUser != submittedAssignment.CreatorID.String() {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errRole)
		return
    }

	if err := c.SubmittedAssignmentService.CreateSubmittedAssignment(submittedAssignment); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetSubmittedAssignment GET /submitted_assignment/{id:string}
func (c *SubmittedAssignmentController) GetSubmittedAssignment() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	submittedAssignmentID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	submittedAssignment, err := c.SubmittedAssignmentService.GetSubmittedAssignment(submittedAssignmentID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::GetSubmittedAssignment", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    submittedAssignment,
	})
}

// UpdateSubmittedAssignment PUT /submitted_assignment/{id:string}
func (c *SubmittedAssignmentController) UpdateSubmittedAssignment() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	submittedAssignmentID := c.Context.Params().Get("id")
	var form submmittedAssignmentUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::UpdateSubmittedAssignment", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.SubmittedAssignmentService.UpdateSubmittedAssignment(submittedAssignmentID, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteSubmittedAssignment DELETE /submitted_assignment/{id:string}
func (c *SubmittedAssignmentController) DeleteSubmittedAssignment() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	submittedAssignmentID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.SubmittedAssignmentService.DeleteSubmittedAssignment(submittedAssignmentID); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::DeleteSubmittedAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
