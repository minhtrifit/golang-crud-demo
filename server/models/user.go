package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       string      `gorm:"not null"`
	Email      string      `gorm:"unique;not null"`
	Password   string      `gorm:"not null"`                        // Mật khẩu (hashed)
	Provider   string      `gorm:"not null;default:'local'"`        // Nhà cung cấp (local, google, etc.)
	Workspaces []Workspace `gorm:"foreignKey:UserID"`
	Tasks      []Task      `gorm:"foreignKey:AssigneeID"`           // Quan hệ với Task qua AssigneeID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateNewUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}