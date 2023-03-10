package update

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
)

type IUpdateVideoRepo interface {
	UpdateVideo(input UpdateVideoInput) *commonModels.RepositoryError
	UpdateVideoCount(id string, count int) *commonModels.RepositoryError
}

type UpdateVideoService struct {
	repo IUpdateVideoRepo
}

func CreateUpdateVideService(repo IUpdateVideoRepo) UpdateVideoService {
	return UpdateVideoService{repo: repo}
}

func (s UpdateVideoService) UpdateVideoByInput(input UpdateVideoInput) *commonModels.ServiceError {
	if err := s.repo.UpdateVideo(input); err != nil {
		return &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}
	return nil
}

func (s UpdateVideoService) UpdateVideoCountByOne(id string) *commonModels.ServiceError {
	if err := s.repo.UpdateVideoCount(id, 1); err != nil {
		return &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}

	return nil
}
