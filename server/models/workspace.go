package models

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Projects    []Project `gorm:"foreignKey:WorkspaceID"`
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}