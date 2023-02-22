package create

import (
	comonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"gorm.io/gorm"
)

type CreateVideoRepository struct {
	db *gorm.DB
}

func CreateAddVideoRepository(db *gorm.DB) CreateVideoRepository {
	return CreateVideoRepository{db: db}
}

func (repo *CreateVideoRepository) CreateVideo(input CreateVideoInput) (*models.Video, *comonModels.RepositoryError) {
	var video models.Video
	video.UserId = input.UserId
	video.ChannelID = input.ChannelID
	video.Name = input.Name
	video.Image = input.Image
	video.Description = input.Description
	video.Source = input.Source
	video.SourceType = input.SourceType
	result := repo.db.Debug().Create(&video)

	if result.Error != nil {
		return nil, &comonModels.RepositoryError{Code: comonModels.RepositoryErrorCode_Fail}
	}

	return &video, nil
}
