package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	"gorm.io/gorm"
)

type CreateCommentRepository struct {
	db *gorm.DB
}

func InitCreateCommentRepository(db *gorm.DB) CreateCommentRepository {
	return CreateCommentRepository{db: db}
}

func (repo CreateCommentRepository) CreateComment(input CreateCommentInput) (*models.Comment, commonModels.RepositoryErrorCode) {
	var comment models.Comment
	comment.UserId = input.UserId
	comment.VideoID = input.VideoID
	comment.CommentID = input.CommentID
	comment.Content = input.Content
	result := repo.db.Debug().Create(&comment)
	if result.Error != nil {
		return nil, commonModels.RepositoryErrorCode_Fail
	}

	return &comment, commonModels.RepositoryErrorCode_OK
}
