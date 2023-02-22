package delete

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
)

type IDeleteCommentRepository interface {
	DeleteComment(input DeleteCommentInput) *commonModels.RepositoryError
}
type DeleteCommentService struct {
	repo IDeleteCommentRepository
}

func CreateDeleteCommentService(repo IDeleteCommentRepository) DeleteCommentService {
	return DeleteCommentService{repo: repo}
}

func (s *DeleteCommentService) DeleteComment(input DeleteCommentInput) *commonModels.ServiceError {
	err := s.repo.DeleteComment(input)
	if err != nil {
		return &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
	}

	return nil
}
