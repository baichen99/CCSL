package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// CarouselController is for carousels CRUD
type CarouselController struct {
	Context         iris.Context
	CarouselService services.CarouselInterface
}

// BeforeActivation will register routes for controllers
func (c *CarouselController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetCarouselsList")
	app.Handle("POST", "/", "CreateCarousel", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("GET", "/{id: string}", "GetCarousel")
	app.Handle("PUT", "/{id: string}", "UpdateCarousel", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeleteCarousel", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

// GetCarouselsList GET
func (c *CarouselController) GetCarouselsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "created_at")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "CarouselController::GetCarouselsList", errParams)
		return
	}
	titleZh := c.Context.URLParamDefault("titleZh", "")
	titleEn := c.Context.URLParamDefault("titleEn", "")
	state := c.Context.URLParamDefault("state", "")
	listParameters := utils.GetCarouselListParameters{
		GetListParameters: listParams,
		TitleZh:           titleZh,
		TitleEn:           titleEn,
		State:             state,
	}
	carousels, count, err := c.CarouselService.GetCarouselList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "CarouselService::GetCarouselsList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    carousels,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateCarousel POST /carousels
func (c *CarouselController) CreateCarousel() {
	defer c.Context.Next()
	var form carouselCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "CarouselCroller::CreateCarousel", errParams)
		return
	}
	// PSQL - Create carousel in database.
	carousel := form.ConvertToModel()
	// Set createor ID
	tokenUser, _ := middlewares.GetJWTParams(c.Context)
	tokenID, _ := uuid.FromString(tokenUser)
	carousel.CreatorID = tokenID
	if err := c.CarouselService.CreateCarousel(carousel); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "Carouselervice::CreateCarousel", errSQL)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

// GetCarousel GET /carousels
func (c *CarouselController) GetCarousel() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	carouselID := c.Context.Params().Get("id")
	carousel, err := c.CarouselService.GetCarousel(carouselID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "CarouselCroller::GetCarousel", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    carousel,
	})
}

// UpdateCarousel POST /carousels/{id:string}
func (c *CarouselController) UpdateCarousel() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	carouselID := c.Context.Params().Get("id")
	var form carouselUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "CarouselCroller::UpdateCarousel", errParams)
		return
	}
	updateData := utils.MakeUpdateData(form)
	// PSQL - Update of the given ID
	if err := c.CarouselService.UpdateCarousel(carouselID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "CarouselService::UpdateCarousel", errSQL)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteCarousel DELETE /carousels/{id:string}
func (c *CarouselController) DeleteCarousel() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	carouselID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.CarouselService.DeleteCarousel(carouselID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "CarouselService::DeleteCarousel", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
