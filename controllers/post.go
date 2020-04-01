package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// NotificationController is for Notification CRUD
type PostController struct {
	Context     iris.Context
	PostService services.PostInterface
}

// BeforeActivation will register routes for controllers
func (c *PostController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/{id: string}", "GetPost")
	app.Handle(iris.MethodGet, "/", "GetPostList")
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodPost, "/", "CreatePost")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdatePost")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeletePost")
}

// GetPostList GET /posts
// >>>>> DOCS  <<<<<
// =================
// @Tags Posts
// @Summary List posts
// @Description get notifications list
// @Accept json
// @Produce json
// @Router /posts [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param message 	query string false 	"search message of notification"
// @Success 200 {object} controllers.GetPostListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *PostController) GetPostList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "posts.id")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "PostController::GetPostList", errParams)
		return
	}
	listParameters := utils.GetPostListParameters{
		GetListParameters: listParams,
		Title:             c.Context.URLParamDefault("title", ""),
		CreatorID:         c.Context.URLParamDefault("creatorID", ""),
	}
	posts, count, err := c.PostService.GetPostList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "PostService::GetPostList", errSQL)
		return
	}
	c.Context.JSON(GetPostListResponse{
		GetListResponse{
			Message: success,
			Page:    listParameters.Page,
			Limit:   listParameters.Limit,
			Total:   count,
		},
		posts,
	})
}

// GetPost GET /posts/{id: string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Posts
// @Summary Get post
// @Description Get a post by id
// @Accept json
// @Produce json
// @Router /posts/{id} [GET]
// @Param   id  path        string  true    "post id" format(uuid)
// @Success 200 {object}    controllers.GetPostResponse
// @Failure 422 {object}    controllers.ErrorResponse
// =================
func (c *PostController) GetPost() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Looking for specified word wia the ID.
	post, err := c.PostService.GetPost(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "PostService::GetPost", errSQL)
		return
	}

	c.Context.JSON(GetPostResponse{
		success,
		post,
	})
}

// CreatePost POST /posts
// >>>>> DOCS  <<<<<
// =================
// @Tags Posts
// @summary Create post
// @Description create a new post
// @Accept json
// @Produce json
// @Router /posts   [POST]
// @Param   post    body        controllers.PostCreateForm  true    "create post"
// @Success 201     {object}    controllers.SuccessResponse
// @Failure 400     {object}    controllers.ErrorResponse
// @Failure 422     {object}    controllers.ErrorResponse
// =================
func (c *PostController) CreatePost() {
	defer c.Context.Next()
	var form PostCreateForm
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "PostController::CreatPost", errParams)
		return
	}

	post := form.ConvertToModel()
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	post.CreatorID = tokenID

	if err := c.PostService.CreatePost(post); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "PostService::CreatePost", errSQL)
		return
	}

	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(SuccessResponse{success})
}

// UpdatePosts PUT /posts/{id: string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Posts
// @Summary Update post
// @Description update a post by id
// @Accept  json
// @Produce json
// @Router  /posts{id}  [PUT]
// @Param   id      path        string      true    "post id" format(uuid)
// @Success 204
// @Failure 400     {object}    controllers.ErrorResponse
// @Failure 401     {object}    controllers.ErrorResponse
// @Failure 422     {object}    controllers.ErrorResponse
// =================
func (c *PostController) UpdatePost() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form PostUpdateForm

	// Read Json from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "PostController::UpdatePost", errParams)
		return
	}
	post, _ := c.PostService.GetPost(id)

	// Get userID from token
	// and check if user is post's creator
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	if tokenID != post.CreatorID {
		// Returns 401 Unauthorized
		utils.SetError(c.Context, iris.StatusUnauthorized, "PostController::UpdatePost", errAuth)
		return
	}

	// Make UpdateData into object & update in database
	updateData := utils.MakeUpdateData(form)
	if err := c.PostService.UpdatePost(id, updateData); err != nil {
		// Returns 422 UnprocessableEntity
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "PostService::UpdatePost", errSQL)
		return
	}

	// Returns with 204 NoContent status
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeletePost DELETE /posts/{id: string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Posts
// @Summary Delete post
// @Description delete a post by id
// @Accept  json
// @Produce json
// @Router  /posts/{id} [DELETE]
// @Param   id      path        string      true "post id" format(uuid)
// @Success 204
// @Failure 401     {object}    controllers.ErrorResponse
// @Failure 422     {object}    controllers.ErrorResponse
// =================
func (c *PostController) DeletePost() {
	defer c.Context.Next()
	id := c.Context.Params().Get("id")

	post, _ := c.PostService.GetPost(id)

	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	if tokenID != post.CreatorID {
		utils.SetError(c.Context, iris.StatusUnauthorized, "PostController::DeleteReply", errAuth)
		return
	}

	if err := c.PostService.DeletePost(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "PostService::DeletePost", errSQL)
		return
	}

	c.Context.StatusCode(iris.StatusNoContent)
}
