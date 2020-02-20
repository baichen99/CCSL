package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// NewsController is for news CRUD
type NewsController struct {
	Context     iris.Context
	NewsService services.NewsInterface
}

// BeforeActivation will register routes for controllers
func (c *NewsController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetNewsList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetNews")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateNews")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateNews")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteNews")
}

// GetNewsList returns news list
func (c *NewsController) GetNewsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "news.date")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "NewsController::GetNewsList", errParams)
		return
	}
	listParameters := utils.GetNewsListParameters{
		GetListParameters: listParams,
		Column:            c.Context.URLParamDefault("column", ""),
		Title:             c.Context.URLParamDefault("title", ""),
		Type:              c.Context.URLParamDefault("type", ""),
		Text:              c.Context.URLParamDefault("text", ""),
		Language:          c.Context.URLParamDefault("language", ""),
		State:             c.Context.URLParamDefault("state", ""),
	}
	news, count, err := c.NewsService.GetNewsList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NewsService::GetNewsList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    news,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateNews POST /news
func (c *NewsController) CreateNews() {
	defer c.Context.Next()
	var form NewsCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "NewsController::CreateNews", errParams)
		return
	}
	// PSQL - Create news in database.
	news := form.ConvertToModel()
	// Set createor ID
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	news.CreatorID = tokenID
	if err := c.NewsService.CreateNews(news); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NewsService::CreateNews", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetNews GET /news/{id:string}
func (c *NewsController) GetNews() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	newsID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	news, err := c.NewsService.GetNews(newsID)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NewsService::GetNews", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    news,
	})
}

// UpdateNews PUT /news/{id:string}
func (c *NewsController) UpdateNews() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form NewsUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "NewsController::UpdateNews", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.NewsService.UpdateNews(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NewsService::UpdateNews", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteNews DELETE /news/{id:string}
func (c *NewsController) DeleteNews() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.NewsService.DeleteNews(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NewsService::DeleteNews", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
