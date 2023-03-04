package query

import (
	"fmt"
	"strings"

	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"gorm.io/gorm"
)

type QueryRepository struct {
	db *gorm.DB
}

func CreateQueryRepository(db *gorm.DB) QueryRepository {
	return QueryRepository{db: db}
}

func (repo *QueryRepository) QueryVideo(input QueryInput) ([]models.Video, *commonModels.RepositoryError) {
	var videos []models.Video
	pageIndex := input.PageIndex - 1
	condition := map[string]string{}

	if len(input.ChannelID) > 0 {
		condition["ChannelID"] = input.ChannelID
	}

	var conditionPair []string
	for key, val := range condition {
		conditionPair = append(conditionPair, fmt.Sprintf("%s == %s", key, val))
	}

	conditionString := strings.Join(conditionPair, " AND ")
	err := repo.db.Where(conditionString).Offset(int(pageIndex * input.PageSize)).Limit(int(input.PageSize)).Model(&models.Video{}).Find(&videos).Error

	if err != nil {
		return videos, &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_Fail}
	}
	return videos, nil
}
