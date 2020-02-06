package controllers

import (
    "ccsl/configs"
    "ccsl/middlewares"
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
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleSuperUser, configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateCourse")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateCourse")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteCourse")
}

// GetCourseList returns course list
func (c *CourseController) GetCourseList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "courses.name")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "CourseController::GetCourseList", errParams)
		return
	}
	name := c.Context.URLParamDefault("name", "")
	content := c.Context.URLParamDefault("content", "")
	classID := c.Context.URLParamDefault("class_id", "")
	listParameters := utils.GetCourseListParameters{
		GetListParameters: listParams,
		Name:              name,
		Content:           content,
		ClassID:           classID,
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

// CreateCourse POST /course
func (c *CourseController) CreateCourse() {
	defer c.Context.Next()
	var form courseCreateForm
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

// GetCourse GET /course/{id:string}
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

// UpdateCourse PUT /course/{id:string}
func (c *CourseController) UpdateCourse() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	courseID := c.Context.Params().Get("id")
	var form courseUpdateForm

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
