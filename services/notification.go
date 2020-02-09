package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// NotificationInterface struct
type NotificationInterface interface {
	GetNotificationList(parameters utils.GetNotificationListParameters) (notifications []models.Notification, count int, unreadCount int, err error)
	GetNotification(id string) (notification models.Notification, err error)
	CreateNotification(notification models.Notification) (err error)
	UpdateNotification(id string, updateData map[string]interface{}) (err error)
	DeleteNotification(id string) (err error)
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
func (s *NotificationService) GetNotificationList(parameters utils.GetNotificationListParameters) (notifications []models.Notification, count int, unreadCount int, err error) {
	db := s.PG.
		Scopes(
			utils.FilterByColumn("notifications.user_id", parameters.UserID),
			utils.SearchByColumn("notifications.title", parameters.Title),
			utils.SearchByColumn("notifications.message", parameters.Message),
		)

	if err = db.Model(&models.Notification{}).
		Count(&count).
		Error; err != nil {
		return
	}

	if err = db.Model(&models.Notification{}).
		Where("notifications.read_at IS NULL").
		Count(&unreadCount).
		Error; err != nil {
		return
	}

	err = db.
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&notifications).
		Error

	return
}

// GetNotification returns Notification with given id
func (s *NotificationService) GetNotification(id string) (notification models.Notification, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&notification).
		Error
	return
}

// CreateNotification creates a new Notification
func (s *NotificationService) CreateNotification(notification models.Notification) (err error) {
	err = s.PG.
		Create(&notification).
		Error
	return
}

// UpdateNotification updates Notification with given id
func (s *NotificationService) UpdateNotification(id string, updatedData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Notification{}).
		Where("id = ?", id).
		Updates(updatedData).
		Error
	return
}

// DeleteNotification soft deletes a Notification with given id
func (s *NotificationService) DeleteNotification(id string) (err error) {
	err = s.PG.
		Where("id = ?", id).
		Delete(&models.Notification{}).
		Error
	return
}
