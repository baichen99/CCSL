package controllers

import (
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"errors"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// UserController is for user CRUD
type UserController struct {
	Context     iris.Context
	UserService services.UserInterface
}

// BeforeActivation will register routes for controllers
func (c *UserController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("POST", "/login", "UserLogin", middlewares.CheckRateLimit(2))
	app.Router().Use(middlewares.CheckJWTToken)
	app.Handle("GET", "/refresh", "RefreshToken", middlewares.CheckRateLimit(2))
	app.Handle("GET", "/{id: string}", "GetUser")
	app.Handle("PUT", "/{id: string}", "UpdateUser")
	app.Router().Use(middlewares.CheckSuper)
	app.Handle("GET", "/", "GetUsersList")
	app.Handle("POST", "/", "CreateUser")
	app.Handle("DELETE", "/{id: string}", "DeleteUser")
}

// GetUsersList GET /users
// >>>>> DOCS  <<<<<
// =================
// @Tags Users
// @Summary List users
// @Description get users list
// @Accept  json
// @Produce json
// @Router /users [GET]
// @Param page 		query int 	 false	"select from page"
// @Param limit 	query int 	 false	"limit number"
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 	enums(asc, desc)
// @Param userType 	query string false	"filter type of user"	enums(super, admin, user, learner)
// @Param name 		query string false 	"search name of user"
// @Param username 	query string false 	"search username of user"
// @Param state 	query string false 	"filter state of user" 	enums(active, inactive)
// @Success 200 {object} controllers.GetUsersListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *UserController) GetUsersList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "users.username")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::GetUsersList", errParams)
		return
	}
	userType := c.Context.URLParamDefault("userType", "")
	name := c.Context.URLParamDefault("name", "")
	username := c.Context.URLParamDefault("username", "")
	state := c.Context.URLParamDefault("state", "")
	listParameters := utils.GetUserListParameters{
		GetListParameters: listParams,
		UserType:          userType,
		Username:          username,
		Name:              name,
		State:             state,
	}
	users, count, err := c.UserService.GetUsersList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UserService::GetUsersList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    users,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// GetUsersListResponse Response for GetUsersList
type GetUsersListResponse struct {
	GetListResponse
	Data []models.User `json:"data"`
}

// CreateUser POST /users
// >>>>> DOCS  <<<<<
// =================
// @Tags Users
// @Summary Create user
// @Description create a new user
// @Accept  json
// @Produce json
// @Router 	/users 	[POST]
// @Param 	user 	body	 controllers.UserCreateForm	  true	"create user"
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *UserController) CreateUser() {
	defer c.Context.Next()
	var form UserCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::CreateUser", errParams)
		return
	}

	// If user not shu account, must have password set
	if form.Password == "" {
		if !utils.IsShuUser(form.Username) {
			utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::CreateUser", errors.New("PasswordRequired"))
			return
		}
	}

	// PSQL - Create user in database.
	user := form.ConvertToModel()
	if err := c.UserService.CreateUser(user); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UserService::CreateUser", errSQL)
		return
	}

	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetUser GET /users/{id:string}
func (c *UserController) GetUser() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	userID := c.Context.Params().Get("id")

	// Only super admin and user
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)

	if !(tokenRole == "super" || tokenUser == userID) {
		utils.SetResponseError(c.Context, iris.StatusForbidden, "UserController::GetUser", errRole)
		return
	}

	// PSQL - Looking for specified user via the ID.
	user, err := c.UserService.GetUser("id", userID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UserService::GetUser", errSQL)
		return
	}

	// Returning user information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    user,
	})
}

// UpdateUser Controller: PUT /users/{id: string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Users
// @Summary Update user
// @Description update a user by id
// @Accept  json
// @Produce json
// @Router 	/users/{id} [PUT]
// @Param 	id 		path	 string						  true	"user id" format(uuid)
// @Param 	user 	body	 controllers.UserUpdateForm	  true	"updated user"
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *UserController) UpdateUser() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	userID := c.Context.Params().Get("id")
	var form UserUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::UpdateUser", errParams)
		return
	}

	// Only super admin and user
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)
	if tokenRole != "super" && tokenUser != userID {
		utils.SetResponseError(c.Context, iris.StatusForbidden, "UserController::UpdateUser", errRole)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given user ID.
	if err := c.UserService.UpdateUser(userID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UserService::UpdateUser", errSQL)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteUser Controller: DELETE /users/{id: string}
func (c *UserController) DeleteUser() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	userID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given user.
	if err := c.UserService.DeleteUser(userID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserService::DeleteUser", errParams)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// UserLogin POST /users/login
func (c *UserController) UserLogin() {
	defer c.Context.Next()
	var form userLoginForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::UserLogin", errParams)
		return
	}

	// PSQL - Get user from database
	user, err := c.UserService.GetUser("username", form.Username)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserService::GetUser", errAuth)
		return
	}

	// Only active user can login
	if user.State != "active" {
		utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin", errors.New("InactiveAccount"))
		return
	}

	if utils.IsShuUser(form.Username) {
		utils.LogInfo(c.Context, "SHU OAuth Login")
		// Is SHU user, auth by shu oauth
		result, err := utils.ShuLogin(form.Username, form.Password)
		if err != nil {
			utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::OauthLogin", errors.New("OauthUnavailable"))
			return
		}
		if !result {
			utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::OauthLogin", errAuth)
			return
		}
	} else {
		utils.LogInfo(c.Context, "Password Login")
		// Not shu user, compare password
		if result := utils.ComparePassword(user.Password, form.Password); !result {
			utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::PasswordLogin", errAuth)
			return
		}
	}
	token, err := middlewares.SignJWTToken(user.ID, user.UserType)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::SignToken", errAuth)
		return
	}
	ipAddress := c.Context.GetHeader("X-Real-IP")
	c.UpdateUserLoginHistory(ipAddress, user.ID)
	c.Context.JSON(iris.Map{
		message: success,
		data:    token,
	})
}

// RefreshToken POST /users/refresh
func (c *UserController) RefreshToken() {
	defer c.Context.Next()
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)
	userID, _ := uuid.FromString(tokenUser)
	token, err := middlewares.SignJWTToken(userID, tokenRole)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserController::RefreshToken::SignToken", errAuth)
		return
	}
	ipAddress := c.Context.GetHeader("X-Real-IP")
	c.UpdateUserLoginHistory(ipAddress, userID)
	c.Context.JSON(iris.Map{
		message: success,
		data:    token,
	})
}

// UpdateUserLoginHistory nsters log to login history
func (c *UserController) UpdateUserLoginHistory(ipAddress string, userID uuid.UUID) {
	defer c.Context.Next()
	info, _ := utils.GetIPInfo(ipAddress)
	info.UserID = userID
	if err := c.UserService.CreateLoginHistory(info); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UserService::CreateLoginHistory", errSQL)
		return
	}
}
