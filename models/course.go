package models

import uuid "github.com/satori/go.uuid"

// Course model
type Course struct {
	Base
	ClassID     uuid.UUID    `gorm:"type:uuid;not null;" json:"classID"`
	Name        string       `json:"name"`
	Content     string       `json:"content"`
	Assignments []Assignment `json:"assignments"`
}
