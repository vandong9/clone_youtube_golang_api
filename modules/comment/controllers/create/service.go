package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	videoQueryService "com.vandong9.clone_youtube_golang_api/modules/video/controller/query"
	"gorm.io/gorm"
)

type ICreateCommentRepository interface {
	GetDB() *gorm.DB
	GetCommentByID(ID string) *models.Comment
	CreateComment(input CreateCommentInput) (*models.Comment, commonModels.RepositoryErrorCode)
}
type CreateCommentService struct {
	repo ICreateCommentRepository
}

func InitCreateCommentService(repo ICreateCommentRepository) CreateCommentService {
	return CreateCommentService{repo: repo}
}

func (s *CreateCommentService) CreateComment(input CreateCommentInput) (*models.Comment, *commonModels.ServiceError) {
	// check if this comment is reply to other comment => priority reply more than comment to video
	if len(input.CommentID) > 0 {
		refComment := s.repo.GetCommentByID(input.CommentID)
		if refComment == nil {
			return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_NotFound}
		}
		comment, err := s.repo.CreateComment(input)
		if err != commonModels.RepositoryErrorCode_OK {
			return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
		}

		return comment, nil
	} else {
		// this is case comment to video
		// verify channel id exist
		videoRepo := videoQueryService.CreateQueryRepository(s.repo.GetDB())
		videoService := videoQueryService.CreateQueryVideoService(&videoRepo)

		video := videoService.GetVideoByID(input.VideoID)

		if video == nil {
			return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_NotFound}
		}

		comment, err := s.repo.CreateComment(input)
		if err != commonModels.RepositoryErrorCode_OK {
			return nil, &commonModels.ServiceError{Code: commonModels.ServiceErrorCode_Fail}
		}

		return comment, nil
	}

}
