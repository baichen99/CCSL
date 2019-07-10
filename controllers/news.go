package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// NewsController is for news CURD
type NewsController struct {
	Context     iris.Context
	NewsService services.NewsInterface
}

// BeforeActivation will register routes for controllers
func (c *NewsController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetNewsList")
	app.Handle("GET", "/{id: string}", "GetNews")
	app.Handle("POST", "/", "CreateNews", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("PUT", "/{id: string}", "UpdateNews", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeleteNews", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

// GetNewsList returns news list
func (c *NewsController) GetNewsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "created_at")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	column := c.Context.URLParamDefault("column", "")
	title := c.Context.URLParamDefault("title", "")
	newsType := c.Context.URLParamDefault("type", "")
	text := c.Context.URLParamDefault("text", "")
	language := c.Context.URLParamDefault("language", "")
	listParameters := utils.GetNewsListParameters{
		GetListParameters: listParams,
		Column:            column,
		Title:             title,
		Type:              newsType,
		Text:              text,
		Language:          language,
	}
	news, count, err := c.NewsService.GetNewsList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "NewsService::GetNewsList", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data: iris.Map{
			"news":   news,
			pageKey:  listParams.Page,
			limitKey: listParams.Limit,
			totalKey: count,
		},
	})
}

func (c *NewsController) CreateNews() {
	defer c.Context.Next()
	var form newsCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, paramsKey, err)
		return
	}
	// PSQL - Create news in database.
	news := form.ConvertToModel()
	if err := c.NewsService.CreateNews(news); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NewsService::CreateNews", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

func (c *NewsController) GetNews() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	newsID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	news, err := c.NewsService.GetNews(newsID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NewsService::GetNews", err)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    news,
	})
}

func (c *NewsController) UpdateNews() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	newsID := c.Context.Params().Get("id")
	var form newsUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, paramsKey, err)
		return
	}
	updateData := form.ConvertToModel()

	// PSQL - Update of the given ID
	if err := c.NewsService.UpdateNews(newsID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "NewsService::UpdateNews", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

func (c *NewsController) DeleteNews() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	newsID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.NewsService.DeleteNews(newsID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NewsService::DeleteNews", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
