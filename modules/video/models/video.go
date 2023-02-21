package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Video struct {
	ID          string `gorm:"primaryKey;"`
	UserId      string `gorm:"type:varchar(255);"`
	ChannelID   string `gorm:"type:varchar(255);"`
	Name        string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:text;"`
	Type        int
	CreatedAt   time.Time
}

func (m *Video) BeforeSave(_ *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
		m.CreatedAt = time.Now()
	}
	return nil
}
