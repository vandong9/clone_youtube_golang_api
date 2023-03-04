package create

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (repo *Repository) CreateChannel(input CreateChannelInput) (models.Channel, error) {
	var channel models.Channel
	channel.Name = input.Name
	channel.UserId = input.UserId
	channel.Thumbnail = input.Thumbnail
	channel.Description = input.Description
	err := repo.db.Create(&channel).Error
	return channel, err
}
