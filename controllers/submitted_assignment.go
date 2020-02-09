package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// SubmittedAssignmentController is for submitted_assignment CRUD
type SubmittedAssignmentController struct {
	Context                    iris.Context
	SubmittedAssignmentService services.SubmittedAssignmentInterface
	AssignmentService          services.AssignmentInterface
	CourseService              services.CourseInterface
	ClassService               services.ClassInterface
}

// BeforeActivation will register routes for controllers
func (c *SubmittedAssignmentController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetSubmittedAssignmentList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetSubmittedAssignment")
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodPost, "/", "CreateSubmittedAssignment", middlewares.CheckUserRole([]string{configs.RoleStudent}))
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteSubmittedAssignment", middlewares.CheckUserRole([]string{configs.RoleTeacher}))
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateSubmittedAssignment", middlewares.CheckUserRole([]string{configs.RoleStudent, configs.RoleTeacher}))
}

// GetSubmittedAssignmentList GET /submitted_assignment
// >>>>> DOCS  <<<<<
// =================
// @Tags SubmittedAssignment
// @Summary List submitted_assignment
// @Description get submitted_assignment list
// @Accept  json
// @Produce json
// @Router /submitted_assignment [GET]
// @Param page 		    query int    false  "select from page" 			    mininum(1)
// @Param limit 	    query int    false  "limit number" 				    mininum(0)
// @Param order 	    query string false	"order by field"
// @Param orderBy 	    query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param creatorID    query string false 	"filter creatorid of submitted_assignment"
// @Param assignmentid query string false 	"filter assignmentid of submitted_assignment"
// @Param graderid     query string false 	"filter graderid of submitted_assignment"
// @Param answer        query string false 	"search answer of submitted_assignment"
// @Param comment       query string false 	"search comment of submitted_assignment"
// @Success 200 {object} controllers.GetSubmmitedAssignmentListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *SubmittedAssignmentController) GetSubmittedAssignmentList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "submitted_assignments.assignment_id")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::GetSubmittedAssignmentList", errParams)
		return
	}
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	assignmentID := c.Context.URLParamDefault("assignmentID", "")
	graderID := c.Context.URLParamDefault("graderID", "")
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

// GetSubmmitedAssignmentListResponse Response for GetSubmmitedAssignmentList
type GetSubmmitedAssignmentListResponse struct {
	GetListResponse
	Data []models.Assignment `json:"data"`
}

// CreateSubmittedAssignment POST /submitted_assignment
// >>>>> DOCS  <<<<<
// =================
// @Tags SubmmitedAssignment
// @Summary Create submmitedAssignment
// @Description create a new submmitedAssignment
// @Accept  json
// @Produce json
// @Router 	/submmited_assignment 	[POST]
// @Param 	assignment 	body	 controllers.submmittedAssignmentCreateForm	  true	"create submmited_assignment"
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *SubmittedAssignmentController) CreateSubmittedAssignment() {
	defer c.Context.Next()
	var form submmittedAssignmentCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::CreateSubmittedAssignment", errParams)
		return
	}

	// Check if out of date
	assginment, err := c.AssignmentService.GetAssignment(form.AssignmentID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}
	if time.Now().After(*assginment.Deadline) {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errOutdated)
		return
	}

	// PSQL - Create submitted_assignment in database.
	submittedAssignment := form.ConvertToModel()

	// Check if current student select the course
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	assignment, err := c.AssignmentService.GetAssignment(submittedAssignment.AssignmentID.String())
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errSQL)
		return
	}
	course, err := c.CourseService.GetCourse(assignment.CourseID.String())
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errSQL)
		return
	}
	class, err := c.ClassService.GetClass(course.ClassID.String())
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errSQL)
		return
	}
	var flag int
	for _, student := range class.Students {
		if student.ID.String() == tokenUser {
			flag = 1
		}
	}

	if flag == 1 {
		submittedAssignment.CreatorID, _ = uuid.FromString(tokenUser)
	} else {
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
// >>>>> DOCS  <<<<<
// =================
// @Tags SubmittedAssignment
// @Summary Get submitted_assignment
// @Description Get a submitted_assignment by id
// @Accept  json
// @Produce json
// @Router 	/submitted_assignment/{id} [GET]
// @Param 	id 		path	 string	    true	"submitted_assignment id" format(uuid)
// @Success 200 	{object} controllers.GetSubmittedAssignmentResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
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

// GetSubmittedAssignmentResponse Response for get submitted_assignment
type GetSubmittedAssignmentResponse struct {
	SuccessResponse
	Data models.SubmittedAssignments `json:"data"`
}

// UpdateSubmittedAssignment PUT /submitted_assignment/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags SubmittedAssignment
// @Summary Update submitted_assignment
// @Description update a submitted_assignment by id
// @Accept  json
// @Produce json
// @Router 	/users/{id} [PUT]
// @Param 	id 		path	 string						  true	"submitted_assignment id" format(uuid)
// @Param 	user 	body	 controllers.submmittedAssignmentUpdateForm	  true	"updated submitted_assignment"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
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
	// Check if out of date
	submittedAssignment, err := c.SubmittedAssignmentService.GetSubmittedAssignment(submittedAssignmentID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}
	assginment, err := c.AssignmentService.GetAssignment(submittedAssignment.AssignmentID.String())
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}

	// Check role, only teacher can update grade
	tokenUser, roles := middlewares.GetJWTParams(c.Context)
	if form.Grade != nil {
		if utils.StringsContains(roles, configs.RoleStudent) != -1 {
			utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errRole)
			return
		}
	}
	// After deadline, only teacher can update data
	if time.Now().After(*assginment.Deadline) && utils.StringsContains(roles, configs.RoleStudent) != -1 {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errOutdated)
		return
	}

	updateData := utils.MakeUpdateData(form)
	updateData["grader_id"] = tokenUser

	// PSQL - Update of the given ID
	if err := c.SubmittedAssignmentService.UpdateSubmittedAssignment(submittedAssignmentID, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteSubmittedAssignment DELETE /submitted_assignment/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags SubmittedAssignment
// @Summary Delete submitted_assignment
// @Description delete a submitted_assignment by id
// @Accept  json
// @Produce json
// @Router 	/submitted_assignment/{id} [DELETE]
// @Param 	id 		path	 string						  true	"submitted_assignment id" format(uuid)
// @Success 204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
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
