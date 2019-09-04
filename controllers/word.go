package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// WordController is for word CURD
type WordController struct {
	Context     iris.Context
	WordService services.WordInterface
}

// BeforeActivation will register routes for controllers
func (c *WordController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetWordsList") //, middlewares.CheckJWTToken)
	app.Handle("POST", "/", "CreateWord", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("GET", "/{id: string}", "GetWord") //, middlewares.CheckJWTToken)
	app.Handle("PUT", "/{id: string}", "UpdateWord", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeleteWord", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

// GetWordsList GET /words
func (c *WordController) GetWordsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "initial")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	wordType := c.Context.URLParamDefault("type", "")
	wordInitial := c.Context.URLParamDefault("initial", "")
	searchChinese := c.Context.URLParamDefault("chinese", "")
	searchEnglish := c.Context.URLParamDefault("english", "")
	listParameters := utils.GetWordListParameters{
		GetListParameters: listParams,
		Type:              wordType,
		Initial:           wordInitial,
		Chinese:           searchChinese,
		English:           searchEnglish,
	}
	words, count, err := c.WordService.GetWordsList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "WordService::GetWordsList", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    words,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateWord POST /words
func (c *WordController) CreateWord() {
	defer c.Context.Next()
	var form wordCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "WordController::ParamsError", err)
		return
	}
	// PSQL - Create word in database.
	word := form.ConvertToModel()
	if err := c.WordService.CreateWord(word); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "WordService::CreateWord", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetWord GET /words/{id:string}
func (c *WordController) GetWord() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	wordID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	word, err := c.WordService.GetWord(wordID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "WordService::GetWord", err)
		return
	}

	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    word,
	})
}

// UpdateWord PUT /words/{id:string}
func (c *WordController) UpdateWord() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	wordID := c.Context.Params().Get("id")
	var form wordUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "WordController::ParamsError", err)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.WordService.UpdateWord(wordID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "WordService::UpdateWord", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteWord DELETE /words/{id:string}
func (c *WordController) DeleteWord() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	wordID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.WordService.DeleteWord(wordID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "WordService::DeleteWord", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
