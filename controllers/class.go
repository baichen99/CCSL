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
	UserService  services.UserInterface
}

// BeforeActivation will register routes for controllers
func (c *ClassController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodGet, "/", "GetClassList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetClass")
	app.Router().Use(middlewares.CheckUserRole([]string{configs.RoleSuperUser}))
	app.Handle(iris.MethodPost, "/", "CreateClass")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateClass")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteClass")
	app.Handle(iris.MethodPost, "/{id: string}/teachers/{uid: string}", "CreateTeacher")
	app.Handle(iris.MethodDelete, "/{id: string}/teachers/{uid: string}", "DeleteTeacher")
	app.Handle(iris.MethodPost, "/{id: string}/students", "CreateStudent")
	app.Handle(iris.MethodDelete, "/{id: string}/students/{uid: string}", "DeleteStudent")
}

// GetClassList GET /classes
// >>>>> DOCS  <<<<<
// =================
// @Tags Classes
// @Summary List classes
// @Description get classes list
// @Accept  json
// @Produce json
// @Router /classes [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param name   	query string false 	"search name of class"
// @Success 200 {object} controllers.GetClassListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *ClassController) GetClassList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "classes.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ClassController::GetClassList", errParams)
		return
	}
	listParameters := utils.GetClassListParameters{
		GetListParameters: listParams,
		Name:              c.Context.URLParamDefault("name", ""),
	}
	classes, count, err := c.ClassService.GetClassesList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::GetClassList", errSQL)
		return
	}
	c.Context.JSON(GetClassListResponse{
		GetListResponse{
			success,
			listParams.Page,
			listParams.Limit,
			count,
		}, classes,
	})
}

// GetClass GET /classes/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Classes
// @Summary Get a class
// @Description Get a class by id
// @Accept  json
// @Produce json
// @Router 	/classes/{id} [GET]
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
	id := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	class, err := c.ClassService.GetClass(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::GetClass", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(GetClassResponse{
		success,
		class,
	})
}

// CreateClass POST /classes
// >>>>> DOCS  <<<<<
// =================
// @Tags Classes
// @Summary Create class
// @Description create a new class
// @Accept  json
// @Produce json
// @Router 	/classes 	[POST]
// @Param 	user 	body	 controllers.ClassCreateForm	  true	"create user"
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *ClassController) CreateClass() {
	defer c.Context.Next()
	var form ClassCreateForm
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
	c.Context.JSON(SuccessResponse{success})
}

// UpdateClass PUT /classes/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Classes
// @Summary Update class
// @Description update a class by id
// @Accept  json
// @Produce json
// @Router 	/classes/{id} [PUT]
// @Param 	id 		path	 string							  true	"class id" format(uuid)
// @Param 	user 	body	 controllers.ClassUpdateForm	  true	"updated class"
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *ClassController) UpdateClass() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form ClassUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ClassController::UpdateClass", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.ClassService.UpdateClass(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::UpdateClass", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteClass DELETE /classes/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Classes
// @Summary Delete class
// @Description delete a class by id
// @Accept  json
// @Produce json
// @Router 	/classes/{id} [DELETE]
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
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.ClassService.DeleteClass(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::DeleteClass", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// CreateTeacher POST /classes/{id:string}/teachers/{id:string}
func (c *ClassController) CreateTeacher() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")
	uid := c.Context.Params().Get("uid")
	// 将uid的老师添加到id班级
	if err := c.ClassService.CreateTeacher(id, uid); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::CreateTeacher", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})
}

// DeleteTeacher POST /classes/{id:string}/teachers/{id:string}
func (c *ClassController) DeleteTeacher() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")
	uid := c.Context.Params().Get("uid")
	if err := c.ClassService.DeleteTeacher(id, uid); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::DeleteTeacher", errSQL)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// CreateStudent POST /classes/{id:string}/students
func (c *ClassController) CreateStudent() {
	defer c.Context.Next()
	var form ClassStudentCreateForm
	id := c.Context.Params().Get("id")
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "ClassController::CreateStudent", errParams)
		return
	}
	// Check if user already exists
	user, err := c.UserService.GetUser("username", form.Username)
	if err != nil {
		c.UserService.CreateUser(
			models.User{
				Username: form.Username,
				Name:     form.Name,
				Roles:    []string{configs.RoleStudent},
				State:    "active",
			})
		user, _ = c.UserService.GetUser("username", form.Username)
	}
	c.ClassService.CreateStudent(id, user.ID.String())
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})
}

// DeleteStudent DELETE /classes/{id:string}/students/{id:string}
func (c *ClassController) DeleteStudent() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")
	uid := c.Context.Params().Get("uid")
	// 将uid的学生从id班级中删除
	if err := c.ClassService.DeleteStudent(id, uid); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "ClassService::DeleteStudent", errSQL)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
