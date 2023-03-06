package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
)

type IQueryCommentRepository interface {
	QueryComment(input QueryCommentInput) ([]models.Comment, commonModels.RepositoryErrorCode)
}

type QueryCommentService struct {
	repo IQueryCommentRepository
}

func CreateQueryCommentService(repo IQueryCommentRepository) QueryCommentService {
	return QueryCommentService{repo: repo}
}

func (s *QueryCommentService) QueryComment(input QueryCommentInput) ([]models.Comment, *commonModels.ServiceError) {
	comments, err := s.repo.QueryComment(input)
	if err != commonModels.RepositoryErrorCode_OK {
		return []models.Comment{}, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}

	return comments, nil
}
