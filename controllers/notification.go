package controllers

import (
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	uuid "github.com/satori/go.uuid"
)

// NotificationController is for Notification CRUD
type NotificationController struct {
	Context             iris.Context
	NotificationService services.NotificationInterface
}

// BeforeActivation will register routes for controllers
func (c *NotificationController) BeforeActivation(app mvc.BeforeActivation) {
	app.Router().Use(middlewares.CheckJWTToken)
	app.Handle("GET", "/", "GetNotificationList")
	app.Handle("GET", "/{id: string}", "GetNotification")
	app.Handle("DELETE", "/{id: string}", "DeleteNotification")
}

// GetNotificationsList GET /notifications
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
// @Param userID 	query string false 	"filter by userID"
// @Success 200 {object} controllers.GetNotificationsListResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 401 {object} controllers.ErrorResponse
// @Failure 422 {object} controllers.ErrorResponse
// =================
func (c *NotificationController) GetNotificationList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "notifications.read_at")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "NotificationController::GetNotificationList", errParams)
		return
	}
	userID, _ := middlewares.GetJWTParams(c.Context)

	listParameters := utils.GetNotificationListParameters{
		GetListParameters: listParams,
		UserID:            userID,
	}
	notifications, count, err := c.NotificationService.GetNotificationList(listParameters)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::GetNotificationList", errSQL)
		return
	}
	c.Context.JSON(iris.Map{
		message: success,
		data:    notifications,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// GetNotificationsListResponse Response for GetUsersList
type GetNotificationsListResponse struct {
	GetListResponse
	Data []models.Notification `json:"data"`
}

// GetNotification GET /notifications/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Notifications
// @Summary GET Notification
// @Description GET a notification by id
// @Accept  json
// @Produce json
// @Router 	/notifications/{id} [GET]
// @Param 	id 		path	 string						  true	"user id" format(uuid)
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *NotificationController) GetNotification() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	notificationID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word via the ID.
	notification, err := c.NotificationService.GetNotification(notificationID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::GetNotification", errSQL)
		return
	}

	if notification.ReadAt.IsZero() {
		notification.ReadAt = time.Now()
		updateData := make(map[string]interface{})
		updateData["ReadAt"] = notification.ReadAt
		if err := c.NotificationService.UpdateNotification(notificationID, updateData); err != nil {
			utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::UpdateNotification", errSQL)
			return
		}
	}

	// Returning word information in data key.
	c.Context.JSON(iris.Map{
		message: success,
		data:    notification,
	})
}

// DeleteNotification DELETE /Notification/{id:string}
// >>>>> DOCS  <<<<<
// =================
// @Tags Notifications
// @Summary DELETE Notification
// @Description DELETE a notification by id
// @Accept  json
// @Produce json
// @Router 	/notifications/{id} [DELETE]
// @Param 	id 		path	 string						  true	"user id" format(uuid)
// @Failure 400 	{object} controllers.ErrorResponse
// @Failure 401 	{object} controllers.ErrorResponse
// @Failure 403 	{object} controllers.ErrorResponse
// @Failure 422 	{object} controllers.ErrorResponse
// =================
func (c *NotificationController) DeleteNotification() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	notificationID := c.Context.Params().Get("id")

	// PSQL - Soft delete of the given ID
	if err := c.NotificationService.DeleteNotification(notificationID); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::DeleteNotification", errSQL)
		return
	}

	// Check notification.userID == UserID
	notfication, err := c.NotificationService.GetNotification(notificationID)
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "NotificationService::GetNotification", errSQL)
		return
	}
	TokenUser, _ := middlewares.GetJWTParams(c.Context)
	userID, _ := uuid.FromString(TokenUser)
	if notfication.UserID != userID {
		utils.SetResponseError(c.Context, iris.StatusForbidden, "NotificationService::GetNotification", errRole)
		return
	}
	// Returns with 204 No Content status.
	c.Context.StatusCode(iris.StatusNoContent)
}
