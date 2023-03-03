package create

type ICreateChannelRepository interface {
	CreateChannel(input CreateChannelInput) error
}

type CreateChannelService struct {
	repo ICreateChannelRepository
}

func CreateChannelServiceInstance(repo ICreateChannelRepository) CreateChannelService {
	return CreateChannelService{repo: repo}
}

func (s *CreateChannelService) CreateChannel(input CreateChannelInput) error {
	return s.repo.CreateChannel(input)
}
