package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
)

type IQueryRepository interface {
	GetVideoByID(ID string) *models.Video
	QueryVideo(input QueryInput) ([]VideoItemList, *commonModels.RepositoryError)
}

type QueryVideoService struct {
	repo IQueryRepository
}

func CreateQueryVideoService(repo IQueryRepository) QueryVideoService {
	return QueryVideoService{repo: repo}
}

func (s *QueryVideoService) GetVideoByID(ID string) *models.Video {
	return s.repo.GetVideoByID(ID)
}

func (s *QueryVideoService) QueryVideo(input QueryInput) ([]VideoItemList, commonModels.ServiceErrorCode) {
	videos, err := s.repo.QueryVideo(input)

	if err != nil {
		return []VideoItemList{}, commonModels.ServiceErrorCode_Fail
	}

	return videos, commonModels.ServiceErrorCode_OK
}
