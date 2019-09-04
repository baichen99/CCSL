package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"
	"errors"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// UserController is for user CURD
type UserController struct {
	Context     iris.Context
	UserService services.UserInterface
}

// BeforeActivation will register routes for controllers
func (c *UserController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetUsersList", middlewares.CheckJWTToken, middlewares.CheckSuper) // super admin only
	app.Handle("POST", "/", "CreateUser", middlewares.CheckJWTToken, middlewares.CheckSuper)  // super admin only
	app.Handle("GET", "/{id: string}", "GetUser", middlewares.CheckJWTToken)                  // super admin and user
	app.Handle("PUT", "/{id: string}", "UpdateUser", middlewares.CheckJWTToken)               // super admin and user
	app.Handle("DELETE", "/{id: string}", "DeleteUser", middlewares.CheckJWTToken)
	app.Handle("POST", "/login", "UserLogin")
}

// GetUsersList GET /users
func (c *UserController) GetUsersList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "username")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	userType := c.Context.URLParamDefault("type", "")
	searchName := c.Context.URLParamDefault("search", "")
	listParameters := utils.GetUserListParameters{
		GetListParameters: listParams,
		UserType:          userType,
		SearchName:        searchName,
	}
	users, count, err := c.UserService.GetUsersList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "UserService::GetUsersList", err)
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

// CreateUser POST /users
func (c *UserController) CreateUser() {
	defer c.Context.Next()
	var form userCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::ParamsError", err)
		return
	}

	// If user not shu account, must have password set
	if form.Password == "" {
		if !utils.IsShuUser(form.Username) {
			utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::ParamsError", errors.New("PasswordRequired"))
			return
		}
	}

	// PSQL - Create user in database.
	user := form.ConvertToModel()
	if err := c.UserService.CreateUser(user); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "UserService::CreateUser", err)
		return
	}

	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetUser GET /user/{id:string}
func (c *UserController) GetUser() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	userID := c.Context.Params().Get("id")

	// Only super admin and user
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)

	if !(tokenRole == "super" || tokenUser == userID) {
		utils.SetResponseError(c.Context, iris.StatusForbidden, "Not super user or self", errors.New("RoleError"))
		return
	}

	// PSQL - Looking for specified user via the ID.
	user, err := c.UserService.GetUser("id", userID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserService::GetUser", err)
		return
	}

	// Returning user information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    user,
	})
}

// UpdateUser Controller: PUT /user/{id: string}
func (c *UserController) UpdateUser() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	userID := c.Context.Params().Get("id")
	var form userUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::ParamsError", err)
		return
	}

	// Only super admin and user
	tokenUser, tokenRole := middlewares.GetJWTParams(c.Context)

	if tokenRole != "super" || tokenUser != userID {
		utils.SetResponseError(c.Context, iris.StatusForbidden, "Not super user or self", errors.New("RoleError"))
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given user ID.
	if err := c.UserService.UpdateUser(userID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserService::UpdateUser", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)

}

// DeleteUser Controller: DELETE /user/{id: string}
func (c *UserController) DeleteUser() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	userID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given user.
	if err := c.UserService.DeleteUser(userID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserService::DeleteUser", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// UserLogin POST /user/login
func (c *UserController) UserLogin() {
	defer c.Context.Next()
	var form userLoginForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "UserController::ParamsError", err)
		return
	}

	// PSQL - Get user from database
	user, err := c.UserService.GetUser("username", form.Username)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnauthorized, "UserService::GetUser", errors.New("AuthFailed"))
		return
	}
	if utils.IsShuUser(form.Username) {
		utils.LogInfo(c.Context, "SHU OAuth Login")
		// Is SHU user, auth by shu oauth
		if result, _ := utils.ShuLogin(form.Username, form.Password); !result {
			utils.SetResponseError(c.Context, iris.StatusUnauthorized, "Password incorrect", errors.New("AuthFailed"))
			return
		}
	} else {
		utils.LogInfo(c.Context, "Password Login")
		// Not shu user, compare password
		if result := utils.ComparePassword(user.Password, form.Password); !result {
			utils.SetResponseError(c.Context, iris.StatusUnauthorized, "Password incorrect", errors.New("AuthFailed"))
			return
		}
	}
	token, _ := middlewares.SignJWTToken(user.ID, user.UserType)
	c.Context.JSON(iris.Map{
		message: success,
		data:    token,
	})
}
