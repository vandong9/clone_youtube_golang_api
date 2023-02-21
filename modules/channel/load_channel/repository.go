package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"gorm.io/gorm"
)

type QueryChannelRepository struct {
	db *gorm.DB
}

func CreateQueryChannelRepo(db *gorm.DB) QueryChannelRepository {
	return QueryChannelRepository{db: db}
}

func (repo *QueryChannelRepository) QueryChannel() []models.Channel {
	var channels []models.Channel
	repo.db.Model(&models.Channel{}).Find(&channels)
	return channels
}
