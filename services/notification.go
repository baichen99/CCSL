package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// NotificationInterface struct
type NotificationInterface interface {
	GetNotificationList(parameters utils.GetNotificationListParameters) (Notifications []models.Notification, count int, err error)
	GetNotification(NotificationID string) (Notification models.Notification, err error)
	CreateNotification(Notification models.Notification) (err error)
	UpdateNotification(NotificationID string, updateData map[string]interface{}) (err error)
	DeleteNotification(NotificationID string) (err error)
}

// NotificationService implements Notification interface
type NotificationService struct {
	PG *gorm.DB
}

// NewNotificationService returns new Notification serivce
func NewNotificationService(pg *gorm.DB) NotificationInterface {
	return &NotificationService{
		PG: pg,
	}
}

// GetNotificationList returns Notifications list
func (s *NotificationService) GetNotificationList(parameters utils.GetNotificationListParameters) (Notifications []models.Notification, count int, err error) {
	db := s.PG.LogMode(false).Scopes(
		utils.FilterByColumn("notifications.user_id", parameters.UserID),
		utils.SearchByColumn("notifications.message", parameters.Message),
	)
	err = db.Model(&models.Notification{}).Count(&count).Error

	if err != nil {
		return
	}
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.
			Order(orderQuery).
			Limit(parameters.Limit).
			Offset(parameters.Limit * (parameters.Page - 1)).
			Find(&Notifications).
			Error
	} else {
		err = db.
			Order(orderQuery).
			Find(&Notifications).
			Error
	}
	return
}

// GetNotification returns Notification with given id
func (s *NotificationService) GetNotification(NotificationID string) (Notification models.Notification, err error) {
	err = s.PG.
		LogMode(false).
		Where("id = ?", NotificationID).
		Take(&Notification).
		Error
	return
}

// CreateNotification creates a new Notification
func (s *NotificationService) CreateNotification(Notification models.Notification) (err error) {
	err = s.PG.
		LogMode(true).
		Create(&Notification).
		Error
	return
}

// UpdateNotification updates Notification with given id
func (s *NotificationService) UpdateNotification(NotificationID string, updateData map[string]interface{}) (err error) {
	var notification models.Notification
	err = s.PG.
		LogMode(true).
		Where("id = ?", NotificationID).
		Take(&notification).
		Update(updateData).
		Error
	return
}

// DeleteNotification soft deletes a Notification with given id
func (s *NotificationService) DeleteNotification(NotificationID string) (err error) {
	var Notification models.Notification
	err = s.PG.
		LogMode(true).
		Where("id = ?", NotificationID).
		Delete(&Notification).
		Error
	return
}
