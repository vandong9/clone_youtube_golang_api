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
	err := repo.db.Debug().Where("").Find(&comments)
	if err != nil {
		return []models.Comment{}, commonModels.RepositoryErrorCode_Fail
	}

	return comments, commonModels.RepositoryErrorCode_OK
}
