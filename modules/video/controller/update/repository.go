package update

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"gorm.io/gorm"
)

type UpdateVideoRepository struct {
	db *gorm.DB
}

func CreateUpdateVideRepository(db *gorm.DB) UpdateVideoRepository {
	return UpdateVideoRepository{db: db}
}

func (repo UpdateVideoRepository) UpdateVideo(input UpdateVideoInput) *commonModels.RepositoryError {
	return nil
}

func (repo UpdateVideoRepository) UpdateVideoCount(id string, count int) *commonModels.RepositoryError {
	var video models.Video
	if err := repo.db.Debug().Where("ID = ?", id).Find(&video).Error; err != nil {
		return &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_NotFound}
	}

	video.ViewCount += int64(count)
	if err := repo.db.Save(&video).Error; err != nil {
		return &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_UpdateFail}
	}
	return nil
}
