package create

import "com.vandong9.clone_youtube_golang_api/modules/channel/models"

type ICreateChannelRepository interface {
	CreateChannel(input CreateChannelInput) (models.Channel, error)
}

type CreateChannelService struct {
	repo ICreateChannelRepository
}

func CreateChannelServiceInstance(repo ICreateChannelRepository) CreateChannelService {
	return CreateChannelService{repo: repo}
}

func (s *CreateChannelService) CreateChannel(input CreateChannelInput) (models.Channel, error) {
	return s.repo.CreateChannel(input)
}
