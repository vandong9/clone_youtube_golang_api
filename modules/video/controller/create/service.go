package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
)

type ICreateVideoRepository interface {
	CreateVideo(input CreateVideoInput) (*models.Video, *commonModels.RepositoryError)
}

type CreateVideoService struct {
	repo ICreateVideoRepository
}

func InitCreateVideoService(repo ICreateVideoRepository) CreateVideoService {
	return CreateVideoService{repo: repo}
}

func (s *CreateVideoService) CreateVideo(input CreateVideoInput) (*models.Video, *commonModels.ServiceError) {
	video, err := s.repo.CreateVideo(input)
	if err != nil {
		return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}

	return video, nil
}
