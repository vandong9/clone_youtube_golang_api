package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"com.vandong9.clone_youtube_golang_api/utils/log"
	"github.com/gin-gonic/gin"
)

type ICreateVideoRepository interface {
	CreateVideo(ctx *gin.Context, input CreateVideoInput) (*models.Video, *commonModels.RepositoryError)
}

type CreateVideoService struct {
	repo ICreateVideoRepository
}

func InitCreateVideoService(repo ICreateVideoRepository) CreateVideoService {
	return CreateVideoService{repo: repo}
}

func (s *CreateVideoService) CreateVideo(ctx *gin.Context, input CreateVideoInput) (*models.Video, *commonModels.ServiceError) {
	video, err := s.repo.CreateVideo(ctx, input)
	if err != nil {
		log.LogInfo(ctx, "Business - Fail"+err.Code.String())
		return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}
	log.LogInfo(ctx, "Business - Success")
	return video, nil
}
