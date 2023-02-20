package create

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (repo *Repository) CreateChannel(input CreateChannelInput) error {
	var channel Channel
	channel.Name = input.Name
	channel.UserId = input.UserId
	channel.ThumnailUrl = input.ThumnailUrl
	err := repo.db.Create(channel).Error
	return err
}
