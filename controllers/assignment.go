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

// AssignmentController is for assignment CRUD
type AssignmentController struct {
	Context           iris.Context
	AssignmentService services.AssignmentInterface
	CourseService     services.CourseInterface
	ClassService      services.ClassInterface
}

// BeforeActivation will register routes for controllers
func (c *AssignmentController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodPost, "/submitted/", "CreateSubmittedAssignment", middlewares.CheckUserRole([]string{configs.RoleStudent}))
	app.Handle(iris.MethodPut, "/submitted/{id: string}", "UpdateSubmittedAssignment", middlewares.CheckUserRole([]string{configs.RoleStudent, configs.RoleTeacher}))
	app.Handle(iris.MethodDelete, "/submitted/{id: string}", "DeleteSubmittedAssignment", middlewares.CheckUserRole([]string{configs.RoleTeacher}))
	app.Router().Use(middlewares.CheckUserRole([]string{configs.RoleTeacher}))
	app.Handle(iris.MethodGet, "/submitted", "GetSubmittedAssignmentList")
	app.Handle(iris.MethodGet, "/submitted/{id: string}", "GetSubmittedAssignment")
	app.Handle(iris.MethodGet, "/", "GetAssignmentList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetAssignment")
	app.Handle(iris.MethodPost, "/", "CreateAssignment")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateAssignment")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteAssignment")
}

// GetAssignmentList GET /assignments
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary List assignments
// @Description get assignment list
// @Accept  json
// @Produce json
// @Router /assignments [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param courseID  query string false 	"filter course_id of assignment"
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
	listParameters := utils.GetAssignmentListParameters{
		GetListParameters: listParams,
		CourseID:          c.Context.URLParamDefault("courseID", ""),
		Title:             c.Context.URLParamDefault("title", ""),
		Type:              c.Context.URLParamDefault("type", ""),
		Content:           c.Context.URLParamDefault("content", ""),
	}
	assignments, count, err := c.AssignmentService.GetAssignmentsList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::GetAssignmentList", errSQL)
		return
	}
	c.Context.JSON(GetAssignmentListResponse{
		GetListResponse{
			success,
			listParameters.Page,
			listParameters.Limit,
			count,
		},
		assignments,
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
// @Tags Assignments
// @Summary Create assignment
// @Description create a new assignment
// @Accept  json
// @Produce json
// @Router 	/assignments 	[POST]
// @Param 	assignment 	body	controllers.AssignmentCreateForm	  true	"create assignment"
// @Success 201		{object} 	controllers.SuccessResponse
// @Failure 400 	{object} 	controllers.ErrorResponse
// @Failure 401 	{object} 	controllers.ErrorResponse
// @Failure 422 	{object} 	controllers.ErrorResponse
// =================
func (c *AssignmentController) CreateAssignment() {
	defer c.Context.Next()
	var form AssignmentCreateForm
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
	c.Context.JSON(SuccessResponse{success})
}

// GetAssignment GET /assignment/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Get assignment
// @Description Get a assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignments/{id} [GET]
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
	id := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	assignment, err := c.AssignmentService.GetAssignment(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::GetAssignment", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(GetAssignmentResponse{
		success,
		assignment,
	})
}

// GetAssignmentResponse Response for get assignment
type GetAssignmentResponse struct {
	Message string            `json:"message" example:"success"`
	Data    models.Assignment `json:"data"`
}

// UpdateAssignment PUT /assignments/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Update assignment
// @Description update a assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignments/{id} [PUT]
// @Param 	id 		path	 string						 		  true	"assignment id" format(uuid)
// @Param 	user 	body	 controllers.AssignmentUpdateForm	  true	"updated assignment"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) UpdateAssignment() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form AssignmentUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "AssignmentController::UpdateAssignment", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.AssignmentService.UpdateAssignment(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::UpdateAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteAssignment DELETE /assignments/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Delete assignment
// @Description delete a assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignments/{id} [DELETE]
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
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.AssignmentService.DeleteAssignment(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "AssignmentService::DeleteAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// GetSubmittedAssignmentList GET /assignments/submitted
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary List submitted assignments
// @Description get submitted assignments list
// @Accept  json
// @Produce json
// @Router /assignments/submitted [GET]
// @Param page 		    query int    false  "select from page" 			    mininum(1)
// @Param limit 	    query int    false  "limit number" 				    mininum(0)
// @Param order 	    query string false	"order by field"
// @Param orderBy 	    query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param creatorID    query string false 	"filter creatorid of submitted assignment"
// @Param assignmentID query string false 	"filter assignmentid of submitted assignment"
// @Param graderid     query string false 	"filter graderid of submitted assignment"
// @Param answer        query string false 	"search answer of submitted assignment"
// @Param comment       query string false 	"search comment of submitted assignment"
// @Success 200 {object} controllers.GetSubmmitedAssignmentListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *AssignmentController) GetSubmittedAssignmentList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "submitted_assignments.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::GetSubmittedAssignmentList", errParams)
		return
	}
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	listParameters := utils.GetSubmittedAssignmentsListParameters{
		GetListParameters: listParams,
		CreatorID:         tokenUser,
		AssignmentID:      c.Context.URLParamDefault("assignmentID", ""),
		GraderID:          c.Context.URLParamDefault("graderID", ""),
		Answer:            c.Context.URLParamDefault("answer", ""),
		Comment:           c.Context.URLParamDefault("comment", ""),
	}
	submittedAssignment, count, err := c.AssignmentService.GetSubmittedAssignmentsList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::GetSubmittedAssignmentList", errSQL)
		return
	}
	c.Context.JSON(GetSubmmitedAssignmentListResponse{
		GetListResponse{
			success,
			listParams.Page,
			listParams.Limit,
			count,
		},
		submittedAssignment,
	})
}

// GetSubmmitedAssignmentListResponse Response for GetSubmmitedAssignmentList
type GetSubmmitedAssignmentListResponse struct {
	GetListResponse
	Data []models.SubmittedAssignment `json:"data"`
}

// CreateSubmittedAssignment POST /assignments/submitted
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Create submmited assignment
// @Description create a new submmitedAssignment
// @Accept  json
// @Produce json
// @Router 	/assignments/submitted 	[POST]
// @Param 	assignment 	body	 controllers.SubmittedAssignmentCreateForm	  true	"create submmited assignment"
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) CreateSubmittedAssignment() {
	defer c.Context.Next()
	var form SubmittedAssignmentCreateForm
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

	flag := false
	for _, student := range class.Students {
		if student.ID.String() == tokenUser {
			flag = true
		}
	}

	if flag == true {
		submittedAssignment.CreatorID, _ = uuid.FromString(tokenUser)
	} else {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errRole)
		return
	}

	if err := c.AssignmentService.CreateSubmittedAssignment(submittedAssignment); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::CreateSubmittedAssignment", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})
}

// GetSubmittedAssignment GET /assignments/submitted{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Get submitted assignment
// @Description Get a submitted assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignments/submitted{id} [GET]
// @Param 	id 		path	 string	    true	"submitted assignment id" format(uuid)
// @Success 200 	{object} controllers.GetSubmittedAssignmentResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) GetSubmittedAssignment() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	submittedAssignment, err := c.AssignmentService.GetSubmittedAssignment(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::GetSubmittedAssignment", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(GetSubmittedAssignmentResponse{
		success,
		submittedAssignment,
	})
}

// GetSubmittedAssignmentResponse Response for get submitted_assignment
type GetSubmittedAssignmentResponse struct {
	Message string                     `json:"message" example:"success"`
	Data    models.SubmittedAssignment `json:"data"`
}

// UpdateSubmittedAssignment PUT /assignments/submitted/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Update submitted assignment
// @Description update a submitted assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignments/submitted/{id} [PUT]
// @Param 	id 		path	 string						  				  true	"submitted assignment id" format(uuid)
// @Param 	user 	body	 controllers.SubmittedAssignmentUpdateForm	  true	"updated submitted_assignment"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) UpdateSubmittedAssignment() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	updateData := make(map[string]interface{})

	var form SubmittedAssignmentUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "SubmittedAssignmentController::UpdateSubmittedAssignment", errParams)
		return
	}
	// Check if out of date
	submittedAssignment, err := c.AssignmentService.GetSubmittedAssignment(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}
	assginment, err := c.AssignmentService.GetAssignment(submittedAssignment.AssignmentID.String())
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}

	// Check role, only teacher can update grade and comment
	tokenUser, roles := middlewares.GetJWTParams(c.Context)
	if form.Grade != nil || form.Comment != nil {
		if !middlewares.HasPermisson(roles, configs.RoleTeacher) {
			utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errRole)
			return
		}
		updateData["GraderID"] = tokenUser
	}

	// After deadline, only teacher can update data
	if assginment.Deadline != nil {
		if time.Now().After(*assginment.Deadline) && !middlewares.HasPermisson(roles, configs.RoleTeacher) {
			utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errOutdated)
			return
		}
	}

	updateData = utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.AssignmentService.UpdateSubmittedAssignment(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::UpdateSubmittedAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteSubmittedAssignment DELETE /assignments/submitted/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Assignments
// @Summary Delete submitted assignment
// @Description delete a submitted assignment by id
// @Accept  json
// @Produce json
// @Router 	/assignments/submitted/{id} [DELETE]
// @Param 	id 		path	 string						  true	"submitted assignment id" format(uuid)
// @Success 204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *AssignmentController) DeleteSubmittedAssignment() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.AssignmentService.DeleteSubmittedAssignment(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "SubmittedAssignmentService::DeleteSubmittedAssignment", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
