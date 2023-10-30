package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	AdminRole Role = "Admin"
	UserRole  Role = "User"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Role      Role           `gorm:"type:enum('Admin', 'User')" json:"role"`
	Name      string
	Email     string
	Password  string
}
