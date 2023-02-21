package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;"`
	Username  string `gorm:"type:varchar(100);unique;not null"`
	Fullname  string `gorm:"type:varchar(255);"`
	Thumbnail string `gorm:"type:varchar(255);"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) BeforeSave(_ *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
		m.CreatedAt = time.Now()
	}
	return nil
}

type CreateUserRequest struct {
	Fullname  string
	Email     string
	Thumbnail string
	Password  string
	CreatedAt time.Time
}

type UpdateUserRequest struct {
	ID        string
	Fullname  string
	Email     string
	Thumbnail string
	Password  string
	Active    bool
	UpdatedAt time.Time
}
