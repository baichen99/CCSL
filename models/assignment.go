package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Assignment model
type Assignment struct {
	Base
	CourseID uuid.UUID  `gorm:"type:uuid;not null" json:"courseID"`
	Title    string     `json:"title"`
	Type     string     `json:"type"`
	Content  string     `json:"content"`
	Deadline *time.Time `json:"deadline"`
}

// SubmittedAssignments model
type SubmittedAssignments struct {
	Base
	AssignmentID uuid.UUID `gorm:"type:uuid" json:"assignmentID"`
	Answer       string    `json:"answer"`
	Grade        *int      `json:"grade"`
	Comment      string    `json:"comment"`
	GraderID     uuid.UUID `gorm:"type:uuid" json:"graderID"`
}
