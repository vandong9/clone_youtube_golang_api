package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Channel struct {
	ID          string `gorm:"primaryKey;"`
	UserId      string `gorm:"type:varchar(255);"`
	Name        string `gorm:"type:varchar(255);"`
	Thumbnail   string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:text"`
	Active      bool   `gorm:"type:bool;default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Channel) BeforeSave(_ *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
		m.CreatedAt = time.Now()
	}
	return nil
}
