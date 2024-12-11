package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Status      string    `gorm:"default:'pending'"`
	ProjectID   uuid.UUID
	AssigneeID  uuid.UUID `gorm:"type:uuid;default:null"` // Trường AssigneeID là khóa ngoại trỏ đến User
	Assignee    User      `gorm:"foreignKey:AssigneeID"` // Quan hệ với User qua AssigneeID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}