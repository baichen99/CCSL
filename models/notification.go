package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Notification data model
type Notification struct {
	Base
	ReadAt  *time.Time `json:"readAt" example:"2019-10-21T09:20:39.943618Z"`
	UserID  uuid.UUID  `gorm:"TYPE:uuid;NOT NULL;index:user_id" json:"-"`
	Title   string     `json:"title" example:"notification title"`
	Message string     `json:"message" example:"notification content"`
}
