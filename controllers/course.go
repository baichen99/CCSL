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

// CourseController is for course CRUD
type CourseController struct {
	Context       iris.Context
	CourseService services.CourseInterface
}

// BeforeActivation will register routes for controllers
func (c *CourseController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetCourseList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetCourse")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleTeacher}))
	app.Handle(iris.MethodPost, "/", "CreateCourse")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateCourse")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteCourse")
}

// GetCourseList GET /courses
// >>>>> DOCS  <<<<<
// =================
// @Tags Courses
// @Summary List courses
// @Description get course list
// @Accept  json
// @Produce json
// @Router /courses [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param calssID  	query string false 	"filter classID of course"
// @Param detail   	query string false 	"search detail name of course"
// @Param name      query string false 	"search name of course"
// @Success 200 {object} controllers.GetCourseListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *CourseController) GetCourseList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "courses.name")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "CourseController::GetCourseList", errParams)
		return
	}
	listParameters := utils.GetCourseListParameters{
		GetListParameters: listParams,
		Name:              c.Context.URLParamDefault("name", ""),
		Content:           c.Context.URLParamDefault("content", ""),
		ClassID:           c.Context.URLParamDefault("classID", ""),
	}
	course, count, err := c.CourseService.GetCoursesList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "CourseService::GetCourseList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    course,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// GetCourseListResponse Response for GetCourseList
type GetCourseListResponse struct {
	GetListResponse
	Data []models.Course `json:"data"`
}

// GetCourse GET /courses/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Courses
// @Summary Get course
// @Description Get a course by id
// @Accept  json
// @Produce json
// @Router 	/courses/{id} [GET]
// @Param 	id 		path	 string	    true	"course id" format(uuid)
// @Success 200 	{object} controllers.GetCourseResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *CourseController) GetCourse() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	courseID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	course, err := c.CourseService.GetCourse(courseID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "CourseService::GetCourse", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    course,
	})
}

// GetCourseResponse Response for get course
type GetCourseResponse struct {
	SuccessResponse
	Data models.Class `json:"data"`
}

// CreateCourse POST /courses
// >>>>> DOCS  <<<<<
// =================
// @Tags Courses
// @Summary Create course
// @Description create a new course
// @Accept  json
// @Produce json
// @Router 	/courses 	[POST]
// @Param 	user 	body	 controllers.CourseCreateForm	  true	"create course"
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *CourseController) CreateCourse() {
	defer c.Context.Next()
	var form CourseCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "CourseController::CreateCourse", errParams)
		return
	}
	// PSQL - Create course in database.
	course := form.ConvertToModel()

	if err := c.CourseService.CreateCourse(course); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "CourseService::CreateCourse", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// UpdateCourse PUT /course/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Courses
// @Summary Update course
// @Description update a class by id
// @Accept  json
// @Produce json
// @Router 	/courses/{id} [PUT]
// @Param 	id 		path	 string							  true	"course id" format(uuid)
// @Param 	course 	body	 controllers.CourseUpdateForm	  true	"updated course"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *CourseController) UpdateCourse() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	courseID := c.Context.Params().Get("id")
	var form CourseUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "CourseController::UpdateCourse", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.CourseService.UpdateCourse(courseID, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "CourseService::UpdateCourse", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteCourse DELETE /course/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Courses
// @Summary Delete course
// @Description delete a course by id
// @Accept  json
// @Produce json
// @Router 	/courses/{id} [DELETE]
// @Param 	id 		path	 string		  true	"course id" format(uuid)
// @Success 204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *CourseController) DeleteCourse() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	courseID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.CourseService.DeleteCourse(courseID); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "CourseService::DeleteCourse", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
