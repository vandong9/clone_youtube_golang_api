package delete

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	"gorm.io/gorm"
)

type DeleteCommentRepository struct {
	db *gorm.DB
}

func CreateDeleteCommentRepository(db *gorm.DB) DeleteCommentRepository {
	return DeleteCommentRepository{db: db}
}

func (repo *DeleteCommentRepository) DeleteComment(input DeleteCommentInput) *commonModels.RepositoryError {
	var comment models.Comment
	comment.ID = input.CommentID
	result := repo.db.Debug().Delete(&comment)
	if err := result.Error; err != nil {
		return &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_Fail}
	}
	return nil
}
