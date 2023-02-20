package create

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Channel struct {
	ID          string `gorm:"primaryKey;"`
	UserId      string `gorm:"type:varchar(255);"`
	Name        string `gorm:"type:varchar(255);"`
	ThumnailUrl string `gorm:"type:varchar(255);"`
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

type CreateChannelInput struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	ThumnailUrl string `json:"thumnail"`
	Description string `json:"description"`
}
