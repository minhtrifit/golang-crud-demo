package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Tasks       []Task `gorm:"foreignKey:ProjectID"`
	WorkspaceID uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
