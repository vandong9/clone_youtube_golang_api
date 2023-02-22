package query

import (
	comonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"gorm.io/gorm"
)

type QueryRepository struct {
	db *gorm.DB
}

func CreateQueryRepository(db *gorm.DB) QueryRepository {
	return QueryRepository{db: db}
}

func (repo *QueryRepository) QueryVideo(input QueryInput) ([]models.Video, *comonModels.RepositoryError) {
	var videos []models.Video
	err := repo.db.Model(&models.Video{}).Find(&videos).Error

	if err != nil {
		return videos, &comonModels.RepositoryError{Code: comonModels.RepositoryErrorCode_Fail}
	}
	return videos, nil
}
