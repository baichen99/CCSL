package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// LexiconController is for word CRUD
type LexiconController struct {
	Context        iris.Context
	LexiconService services.LexiconInterface
}

// BeforeActivation will register routes for controllers
func (c *LexiconController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckJWTToken)
	app.Handle("GET", "/", "GetWordsList")
	app.Handle("GET", "/{id: string}", "GetWord")
	app.Router().Use(middlewares.CheckAdmin)
	app.Handle("POST", "/", "CreateWord")
	app.Handle("PUT", "/{id: string}", "UpdateWord")
	app.Handle("DELETE", "/{id: string}", "DeleteWord")
}

// GetWordsList GET /lexicon
func (c *LexiconController) GetWordsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "lexicons.initial")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "LexiconController::GetWordsList", errParams)
		return
	}
	pos := c.Context.URLParamDefault("pos", "")
	initial := c.Context.URLParamDefault("initial", "")
	searchChinese := c.Context.URLParamDefault("chinese", "")
	searchEnglish := c.Context.URLParamDefault("english", "")
	listParameters := utils.GetLexiconListParameters{
		GetListParameters: listParams,
		Pos:               pos,
		Initial:           initial,
		Chinese:           searchChinese,
		English:           searchEnglish,
	}
	words, count, err := c.LexiconService.GetWordsList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "LexiconService::GetWordsList", errSQL)
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

// CreateWord POST /lexicon
func (c *LexiconController) CreateWord() {
	defer c.Context.Next()
	var form lexiconCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "LexiconController::CreateWord", errParams)
		return
	}
	// PSQL - Create word in database.
	word := form.ConvertToModel()
	if err := c.LexiconService.CreateWord(word); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "LexiconService::CreateWord", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetWord GET /lexicon/{id:string}
func (c *LexiconController) GetWord() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	wordID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	word, err := c.LexiconService.GetWord(wordID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "LexiconService::GetWord", errSQL)
		return
	}

	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    word,
	})
}

// UpdateWord PUT /lexicon/{id:string}
func (c *LexiconController) UpdateWord() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	wordID := c.Context.Params().Get("id")
	var form lexiconUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "LexiconController::UpdateWord", errParams)
		return
	}

	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the given ID
	if err := c.LexiconService.UpdateWord(wordID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "LexiconService::UpdateWord", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteWord DELETE /lexical/words/{id:string}
func (c *LexiconController) DeleteWord() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	wordID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.LexiconService.DeleteWord(wordID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "LexiconService::DeleteWord", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
