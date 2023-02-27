package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
)

type IQueryChannelRepository interface {
	QueryChannel(input QueryInput) []models.Channel
}

type QueryChannelService struct {
	repo IQueryChannelRepository
}

func CreateQueryService(repo IQueryChannelRepository) QueryChannelService {
	return QueryChannelService{repo: repo}
}

func (s *QueryChannelService) QueryChannel(input QueryInput) QueryResponse {
	channels := s.repo.QueryChannel(input)
	return QueryResponse{Channels: channels}
}
