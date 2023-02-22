package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
)

type IQueryRepository interface {
	QueryVideo(input QueryInput) ([]models.Video, *commonModels.RepositoryError)
}

type QueryVideoService struct {
	repo IQueryRepository
}

func CreateQueryVideoService(repo IQueryRepository) QueryVideoService {
	return QueryVideoService{repo: repo}
}

func (s *QueryVideoService) QueryVideo(input QueryInput) ([]models.Video, commonModels.ServiceErrorCode) {
	videos, err := s.repo.QueryVideo(input)

	if err != nil {
		return []models.Video{}, commonModels.ServiceErrorCode_Fail
	}

	return videos, commonModels.ServiceErrorCode_OK
}