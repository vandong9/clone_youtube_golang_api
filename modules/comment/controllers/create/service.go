package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
)

type ICreateCommentRepository interface {
	CreateComment(input CreateCommentInput) (*models.Comment, commonModels.RepositoryErrorCode)
}
type CreateCommentService struct {
	repo ICreateCommentRepository
}

func InitCreateCommentService(repo ICreateCommentRepository) CreateCommentService {
	return CreateCommentService{repo: repo}
}

func (s *CreateCommentService) CreateComment(input CreateCommentInput) (*models.Comment, commonModels.ServiceErrorCode) {
	comment, err := s.repo.CreateComment(input)
	if err != commonModels.RepositoryErrorCode_OK {
		return nil, commonModels.ServiceErrorCode_Fail
	}

	return comment, commonModels.ServiceErrorCode_OK
}
