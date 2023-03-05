package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/channel/load_channel"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"com.vandong9.clone_youtube_golang_api/utils/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICreateVideoRepository interface {
	GetDB() *gorm.DB
	CreateVideo(ctx *gin.Context, input CreateVideoInput) (*models.Video, *commonModels.RepositoryError)
}

type CreateVideoService struct {
	repo ICreateVideoRepository
}

func InitCreateVideoService(repo ICreateVideoRepository) CreateVideoService {
	return CreateVideoService{repo: repo}
}

func (s *CreateVideoService) CreateVideo(ctx *gin.Context, input CreateVideoInput) (*models.Video, *commonModels.ServiceError) {

	// check input userid and channel id is valid
	channelRepo := load_channel.CreateQueryChannelRepo(s.repo.GetDB())
	channelService := load_channel.CreateQueryService(&channelRepo)
	channel := channelService.GetChannelByID(ctx, input.ChannelID)

	if channel == nil {
		log.LogInfo(ctx, "Business - Fail: input channel is not correct: "+input.ChannelID)
		return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}

	video, err := s.repo.CreateVideo(ctx, input)
	if err != nil {
		log.LogInfo(ctx, "Business - Fail"+err.Code.Error())
		return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}

	log.LogInfo(ctx, "Business - Success")
	return video, nil
}
