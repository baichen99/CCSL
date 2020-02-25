package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// UserController is for user CRUD
type UserController struct {
	Context             iris.Context
	UserService         services.UserInterface
	NotificationService services.NotificationInterface
}

// BeforeActivation will register routes for controllers
func (c *UserController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodPost, "/login", "UserLogin", middlewares.CheckRateLimit(2))
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodGet, "/refresh", "RefreshToken", middlewares.CheckRateLimit(2))
	app.Handle(iris.MethodGet, "/{id: string}", "GetUser")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateUser")
	app.Router().Use(middlewares.CheckUserRole([]string{configs.RoleSuperUser}))
	app.Handle(iris.MethodGet, "/", "GetUsersList")
	app.Handle(iris.MethodPost, "/", "CreateUser")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteUser")
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
// @Param page 		query int    false  "select from page" 			mininum(1)
// @Param limit 	query int    false  "limit number" 				mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		enums(asc, desc)
// @Param role	 	query string false	"filter role of user"		 enums(super, admin, user, learner)
// @Param name 		query string false 	"search name of user"
// @Param username 	query string false 	"search username of user"
// @Param state 	query string false 	"filter state of user" 		 enums(active, inactive)
// @Success 200 {object} controllers.GetUsersListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *UserController) GetUsersList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "users.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "UserController::GetUsersList", errParams)
		return
	}
	listParameters := utils.GetUserListParameters{
		GetListParameters: listParams,
		Username:          c.Context.URLParamDefault("username", ""),
		Name:              c.Context.URLParamDefault("name", ""),
		State:             c.Context.URLParamDefault("state", ""),
		Roles:             c.Context.URLParamDefault("roles", ""),
	}
	users, count, err := c.UserService.GetUsersList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::GetUsersList", errSQL)
		return
	}
	c.Context.JSON(GetUsersListResponse{
		GetListResponse{
			success,
			listParams.Page,
			listParams.Limit,
			count,
		}, users,
	})
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
// @Success 201		{object} controllers.SuccessResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *UserController) CreateUser() {
	defer c.Context.Next()
	var form UserCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "UserController::CreateUser", errParams)
		return
	}

	// If user not shu account, must have password set
	if form.Password == "" {
		if !utils.IsShuUser(form.Username) {
			utils.SetError(c.Context, iris.StatusBadRequest, "UserController::CreateUser", errors.New("PasswordRequired"))
			return
		}
	}

	// PSQL - Create user in database.
	user := form.ConvertToModel()
	if err := c.UserService.CreateUser(user); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::CreateUser", errSQL)
		return
	}

	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})
}

// GetUser GET /users/{id:string}
func (c *UserController) GetUser() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// Only super admin and user
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)

	if !(middlewares.HasPermisson(tokenRole, configs.RoleSuperUser) || tokenUser == id) {
		utils.SetError(c.Context, iris.StatusForbidden, "UserController::GetUser", errRole)
		return
	}

	// PSQL - Looking for specified user via the ID.
	user, err := c.UserService.GetUser("id", id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::GetUser", errSQL)
		return
	}

	// Returning user information in data key.
	c.Context.JSON(GetUserResponse{
		success,
		user,
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
// @Success	204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *UserController) UpdateUser() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form UserUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "UserController::UpdateUser", errParams)
		return
	}

	// Only super admin and user
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)
	if !(middlewares.HasPermisson(tokenRole, configs.RoleSuperUser) || tokenUser == id) {
		utils.SetError(c.Context, iris.StatusForbidden, "UserController::UpdateUser", errRole)
		return
	}

	// Only super user can change user role
	if form.Roles != nil && !middlewares.HasPermisson(tokenRole, configs.RoleSuperUser) {
		utils.SetError(c.Context, iris.StatusForbidden, "UserController::UpdateUser", errRole)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given user ID.
	if err := c.UserService.UpdateUser(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::UpdateUser", errSQL)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteUser Controller: DELETE /users/{id: string}
func (c *UserController) DeleteUser() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given user.
	if err := c.UserService.DeleteUser(id); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "UserService::DeleteUser", errParams)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// UserLogin POST /users/login
func (c *UserController) UserLogin() {
	defer c.Context.Next()
	var form UserLoginForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "UserController::UserLogin", errParams)
		return
	}

	// PSQL - Get user from database
	user, err := c.UserService.GetUser("username", form.Username)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnauthorized, "UserService::GetUser", errAuth)
		return
	}

	// Only active user can login
	if user.State != configs.UserStateActive {
		utils.SetError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin", errors.New("InactiveAccount"))
		return
	}

	if utils.IsShuUser(form.Username) {
		utils.LogInfo(c.Context, "SHU OAuth Login")
		// Is SHU user, auth by shu oauth
		result, err := utils.ShuLogin(form.Username, form.Password)
		if err != nil {
			utils.SetError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::OauthLogin", errors.New("OauthUnavailable"))
			return
		}
		if !result {
			utils.SetError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::OauthLogin", errAuth)
			return
		}
	} else {
		utils.LogInfo(c.Context, "Password Login")
		// Not shu user, compare password
		if result := utils.ComparePassword(user.Password, form.Password); !result {
			utils.SetError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::PasswordLogin", errAuth)
			return
		}
	}
	token, err := middlewares.SignJWTToken(user.ID, user.Roles)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnauthorized, "UserController::UserLogin::SignToken", errAuth)
		return
	}
	c.UpdateLoginHistory(user.ID)
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
		utils.SetError(c.Context, iris.StatusUnauthorized, "UserController::RefreshToken::SignToken", errAuth)
		return
	}
	c.UpdateLoginHistory(userID)
	c.Context.JSON(iris.Map{
		message: success,
		data:    token,
	})
}

// UpdateLoginHistory nsters log to login history
func (c *UserController) UpdateLoginHistory(userID uuid.UUID) {
	defer c.Context.Next()
	ipAddress := c.Context.GetHeader("X-Real-IP")
	info, _ := utils.GetIPInfo(ipAddress)
	info.UserID = userID
	loginHistories, _, err := c.UserService.GetLoginHistoryList(utils.GetUserLoginListParameters{
		GetListParameters: utils.GetListParameters{
			Order:   "desc",
			OrderBy: "created_at",
		},
		UserID: userID.String(),
	})
	if len(loginHistories) > 0 {
		lastLogin := loginHistories[0]
		lastLoginTime := lastLogin.CreatedAt.Local().Format("2006-01-02 15:04")
		lastLoginAddress := fmt.Sprintf("%s - %s - %s", lastLogin.Country, lastLogin.RegionName, lastLogin.City)
		notice := fmt.Sprintf("您的账号上次于 %s 在 %s 登录，IP地址为 %s。如果不是您本人登录，请立即修改您的登录账号及密码。", lastLoginTime, lastLoginAddress, info.IP)
		message := models.Notification{
			UserID:  userID,
			Title:   "登录通知",
			Message: notice,
		}
		c.NotificationService.CreateNotification(message)
	}
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::GetLoginHistory", errSQL)
		return
	}
	if err := c.UserService.CreateLoginHistory(info); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "UserService::CreateLoginHistory", errSQL)
		return
	}
}
