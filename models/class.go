package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Class model
type Class struct {
	Base
	Name      string   `json:"name"`
	Teachers  []User   `gorm:"many2many:class_teachers" json:"teachers"`
	Students  []User   `gorm:"many2many:class_students" json:"students"`
	Details   string   `json:"details"`
	Resources string   `json:"resources"`
	Courses   []Course `json:"courses"`
}

// Course model
type Course struct {
	Base
	ClassID     uuid.UUID    `gorm:"type:uuid;not null;" json:"classID"`
	Name        string       `json:"name"`
	Content     string       `json:"content"`
	Assignments []Assignment `json:"assignments"`
}

// Assignment model
type Assignment struct {
	Base
	CourseID uuid.UUID  `gorm:"type:uuid;not null" json:"courseID"`
	Title    string     `json:"title"`
	Type     string     `json:"type"`
	Content  string     `json:"content"`
	Deadline *time.Time `json:"deadline"`
}

// SubmittedAssignment model
type SubmittedAssignment struct {
	Base
	CreatorID    uuid.UUID `gorm:"type:uuid" json:"creatorID"`
	GraderID     uuid.UUID `gorm:"type:uuid" json:"graderID"`
	AssignmentID uuid.UUID `gorm:"type:uuid" json:"assignmentID"`
	Answer       string    `json:"answer"`
	Grade        *int      `json:"grade"`
	Comment      string    `json:"comment"`
}
