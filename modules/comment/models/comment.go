package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID           string `gorm:"primaryKey;"`
	UserId       string `gorm:"type:varchar(100); not null"`
	VideoID      string `gorm:"type:varchar(100);"`
	CommentID    string `gorm:"type:varchar(100);"`
	Content      string `gorm:"type:text;"`
	ReplyCount   int
	LikeCount    int
	DislikeCount int
	UpdatedAt    time.Time
}

func (m *Comment) BeforeSave(_ *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.UpdatedAt = time.Now()

	return nil
}
