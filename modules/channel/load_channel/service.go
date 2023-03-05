package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IQueryChannelRepository interface {
	GetDB() *gorm.DB
	QueryChannel(input QueryInput) []models.Channel
	GetChannelByID(ctx *gin.Context, channelID string) *models.Channel
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

func (s *QueryChannelService) GetChannelByID(ctx *gin.Context, channelID string) *models.Channel {
	return s.repo.GetChannelByID(ctx, channelID)
}
