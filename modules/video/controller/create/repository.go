package create

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"com.vandong9.clone_youtube_golang_api/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateVideoRepository struct {
	db *gorm.DB
}

func CreateAddVideoRepository(db *gorm.DB) CreateVideoRepository {
	return CreateVideoRepository{db: db}
}

func (repo *CreateVideoRepository) GetDB() *gorm.DB {
	return repo.db
}

func (repo *CreateVideoRepository) CreateVideo(ctx *gin.Context, input CreateVideoInput) (*models.Video, *commonModels.RepositoryError) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_Fail}
	}

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
		logger.LogInfo(ctx, "Repository - error: "+result.Error.Error())
		return nil, &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_Fail}
	}
	logger.LogInfo(ctx, "Repository - Success: ")
	return &video, nil
}
