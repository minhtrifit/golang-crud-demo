package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       string      `gorm:"not null"`
	Email      string      `gorm:"unique;not null"`
	Workspaces []Workspace `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}