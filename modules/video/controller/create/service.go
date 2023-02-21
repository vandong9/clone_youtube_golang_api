package create

import (
	comonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
)

type ICreateVideoRepository interface {
	CreateVideo(input CreateVideoInput) (*models.Video, *comonModels.RepositoryError)
}

type CreateVideoService struct {
	repo ICreateVideoRepository
}

func InitCreateVideoService(repo ICreateVideoRepository) CreateVideoService {
	return CreateVideoService{repo: repo}
}

func (s *CreateVideoService) CreateVideo(input CreateVideoInput) (*models.Video, error) {
	video, err := s.repo.CreateVideo(input)
	if err != nil {
		return nil, err
	}

	return video, nil
}
