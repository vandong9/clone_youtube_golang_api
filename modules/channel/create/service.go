package create

import "com.vandong9.clone_youtube_golang_api/modules/channel/models"

type ICreateChannelRepository interface {
	CreateChannel(input CreateChannelInput) error
	GetUserIDByGivenToken(token string) (*models.UserToken, error)
}

type CreateChannelService struct {
	repo ICreateChannelRepository
}

func CreateChannelServiceInstance(repo ICreateChannelRepository) CreateChannelService {
	return CreateChannelService{repo: repo}
}

func (s *CreateChannelService) CreateChannel(input CreateChannelInput) {
	s.repo.CreateChannel(input)
}

func (s *CreateChannelService) CheckValidToken(token string) (*models.UserToken, error) {
	userToken, err := s.repo.GetUserIDByGivenToken(token)
	return userToken, err
}
