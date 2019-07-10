package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type CarouselController struct {
	Context         iris.Context
	CarouselService services.CarouselInterface
}

func (c *CarouselController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetCarouselsList")
	app.Handle("POST", "/", "CreateCarousel", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("GET", "/{id: string}", "GetCarousel")
	app.Handle("PUT", "/{id: string}", "UpdateCarousel", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id: string}", "DeleteCarousel", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

func (c *CarouselController) GetCarouselsList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "create_at")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
		return
	}
	title := c.Context.URLParamDefault("title", "")
	listParameters := utils.GetCarouselListParameters{
		GetListParameters: listParams,
		Title:             title,
	}
	carousels, count, err := c.CarouselService.GetCarouselList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusInternalServerError, "CarouselService::GetCarouselsList", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data: iris.Map{
			"carousels": carousels,
			pageKey:     listParams.Page,
			limitKey:    listParams.Limit,
			totalKey:    count,
		},
	})
}

func (c *CarouselController) CreateCarousel() {
	defer c.Context.Next()
	var form carouselCreateForm
	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, paramsKey, err)
		return
	}
	// PSQL - Create carousel in database.
	carousel := form.ConvertToModel()
	if err := c.CarouselService.CreateCarousel(carousel); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "Carouselervice::CreateCarousel", err)
		return
	}
	// Return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(iris.Map{
		message: success,
	})
}

func (c *CarouselController) GetCarousel() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	carouselID := c.Context.Params().Get("id")
	carousel, err := c.CarouselService.GetCarousel(carouselID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "WordService::GetWord", err)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    carousel,
	})
}

func (c *CarouselController) UpdateCarousel() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	carouselID := c.Context.Params().Get("id")
	var form carouselUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, paramsKey, err)
		return
	}
	updateData := form.ConvertToModel()
	// PSQL - Update of the given ID
	if err := c.CarouselService.UpdateCarousel(carouselID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "CarouselService::UpdateCarousel", err)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}

func (c *CarouselController) DeleteCarousel() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	carouselID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.CarouselService.DeleteCarousel(carouselID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "CarouselService::DeleteCarousel", err)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
