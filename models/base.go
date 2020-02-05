package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base data model
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	CreatedAt time.Time  `json:"createdAt" example:"2019-10-21T09:20:39.943618Z"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2019-10-21T09:20:39.943618Z"`
	DeletedAt *time.Time `json:"-"`
}

// BeforeCreate generates a new UUID and attaches it to the the newly created rows in the database
func (base *Base) BeforeCreate(scope *gorm.Scope) {
	uuidV4 := uuid.NewV4()
	scope.SetColumn("ID", uuidV4)
}
