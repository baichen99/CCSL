package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// NotificationController is for Notification CRUD
type NotificationController struct {
	Context             iris.Context
	NotificationService services.NotificationInterface
}

// BeforeActivation will register routes for controllers
func (c *NotificationController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckToken)
	app.Handle(iris.MethodGet, "/", "GetNotificationList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetNotification")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteNotification")
}

// GetNotificationList GET /notifications
// >>>>> DOCS  <<<<<
// =================
// @Tags Notifications
// @Summary List notifications
// @Description get notifications list
// @Accept  json
// @Produce json
// @Router /notifications [GET]
// @Param page 		query int    false  "select from page" 			    mininum(1)
// @Param limit 	query int    false  "limit number" 				    mininum(0)
// @Param order 	query string false	"order by field"
// @Param orderBy 	query string false	"order by asc or desc" 		    enums(asc, desc)
// @Param message 	query string false 	"search message of notification"
// @Success 200 {object} controllers.GetNotificationsListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *NotificationController) GetNotificationList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "notifications.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "NotificationController::GetNotificationList", errParams)
		return
	}
	userID, _ := middlewares.GetJWTParams(c.Context)

	listParameters := utils.GetNotificationListParameters{
		GetListParameters: listParams,
		UserID:            userID,
	}
	notifications, count, unreadCount, err := c.NotificationService.GetNotificationList(listParameters)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::GetNotificationList", errSQL)
		return
	}
	c.Context.JSON(GetNotificationsListResponse{
		GetListResponse{
			success,
			listParams.Page,
			listParams.Limit,
			count,
		},
		unreadCount,
		notifications,
	})
}

// GetNotification GET /notifications/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Notifications
// @Summary Get notification
// @Description Get a notification by id
// @Accept  json
// @Produce json
// @Router 	/notifications/{id} [GET]
// @Param 	id 		path	 string						  true	"notification id" format(uuid)
// @Success 200 	{object} controllers.GetNotificationResponse
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *NotificationController) GetNotification() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	notification, err := c.NotificationService.GetNotification(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::GetNotification", errSQL)
		return
	}

	if notification.ReadAt == nil {
		currentTime := time.Now()
		notification.ReadAt = &currentTime
		updateData := make(map[string]interface{})
		updateData["ReadAt"] = notification.ReadAt
		if err := c.NotificationService.UpdateNotification(id, updateData); err != nil {
			utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::UpdateNotification", errSQL)
			return
		}
	}

	c.Context.JSON(GetNotificationResponse{
		success,
		notification,
	})
}

// DeleteNotification DELETE /notifications/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Notifications
// @Summary Delete notification
// @Description delete a notification by id
// @Accept  json
// @Produce json
// @Router 	/notifications/{id} [DELETE]
// @Param 	id 		path	 string						  true	"notification id" format(uuid)
// @Success 204
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *NotificationController) DeleteNotification() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// Check notification.UserID == userID
	notfication, err := c.NotificationService.GetNotification(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::GetNotification", errSQL)
		return
	}

	tokenUser, _ := middlewares.GetJWTParams(c.Context)

	if notfication.UserID.String() != tokenUser {
		utils.SetError(c.Context, iris.StatusForbidden, "NotificationService::GetNotification", errRole)
		return
	}

	// PSQL - Soft delete of the given ID
	if err := c.NotificationService.DeleteNotification(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::DeleteNotification", errSQL)
		return
	}

	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
