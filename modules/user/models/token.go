package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	UserID    string `gorm:"type:varchar(100);unique;not null"`
	Token     string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}

func (m *UserToken) BeforeSave(_ *gorm.DB) (err error) {
	m.CreatedAt = time.Now()

	return nil
}
