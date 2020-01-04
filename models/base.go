package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base data model
type Base struct {
	ID        uuid.UUID  `gorm:"TYPE:uuid;PRIMARY_KEY" json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	CreatedAt time.Time  `json:"createdAt" example:"2000-12-30T00:00:00Z"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2000-12-30T00:00:00Z"`
	DeletedAt *time.Time `json:"-"`
}

// BeforeCreate generates a new UUID and attaches it to the the newly created rows in the database
func (base *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}
