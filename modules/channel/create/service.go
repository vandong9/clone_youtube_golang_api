package create

import "gorm.io/gorm"

type ICreateChannelRepository interface {
	CreateChannel(input CreateChannelInput)
}

type CreateChannelService struct {
	repo Repository
}

func CreateChannelServiceInstance(db *gorm.DB) CreateChannelService {
	return CreateChannelService{repo: CreateRepository(db)}
}

func (s *CreateChannelService) CreateChannel(input CreateChannelInput) {
	s.repo.CreateChannel(input)
}
