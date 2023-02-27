package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	"gorm.io/gorm"
)

type QueryCommentRepository struct {
	db *gorm.DB
}

func CreateQueryCommentRepository(db *gorm.DB) QueryCommentRepository {
	return QueryCommentRepository{db: db}
}

func (repo *QueryCommentRepository) QueryComment(input QueryCommentInput) ([]models.Comment, commonModels.RepositoryErrorCode) {
	var comments []models.Comment
	input.PageIndex -= 1
	if input.PageIndex < 0 {
		input.PageIndex = 0
	}
	err := repo.db.Offset(input.PageIndex * input.PageSize).Limit(input.PageSize).Debug().Where("").Find(&comments)
	if err != nil {
		return []models.Comment{}, commonModels.RepositoryErrorCode_Fail
	}

	return comments, commonModels.RepositoryErrorCode_OK
}
