package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey;"`
	Fullname  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserRequest struct {
	Fullname  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UpdateUserRequest struct {
	ID        string
	Fullname  string
	Email     string
	Password  string
	Active    bool
	UpdatedAt time.Time
}
