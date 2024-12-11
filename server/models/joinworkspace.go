package models

import (
	"time"

	"github.com/google/uuid"
)

type JoinWorkSpace struct {
	WorkspaceID uuid.UUID `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"primaryKey"`
	Workspace   Workspace `gorm:"foreignKey:WorkspaceID"`
	User        User      `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}