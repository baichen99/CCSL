package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Notification data model
type Notification struct {
	Base
	ReadAt  time.Time `json:"readAt" example:"2019-10-21T09:20:39.943618Z"`
	UserID  uuid.UUID `gorm:"TYPE:uuid;NOT NULL;index:user_id" json:"userID" example:"550e8400-e29b-41d4-a716-446655440000"`
	Message string    `json:"message"`
}
