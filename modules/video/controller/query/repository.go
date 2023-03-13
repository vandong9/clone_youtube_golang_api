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

func (repo *QueryRepository) GetVideoByID(ID string) *models.Video {
	var video models.Video
	err := repo.db.Debug().Where("ID = ?", ID).Find(&video).Error
	if err != nil {
		return nil
	}
	return &video
}
func (repo *QueryRepository) QueryVideo(input QueryInput) ([]VideoItemList, *commonModels.RepositoryError) {
	var videos []VideoItemList
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
	err := repo.db.Debug().Select("videos.ID, videos.User_id, videos.Channel_ID, videos.Name, videos.Image, videos.Description, videos.Source_Type, videos.Source, users.Fullname as UserFullname, users.Thumbnail as UserAvatar, channels.Name as ChannelName, channels.Thumbnail as ChannelThumbnail").Where(conditionString).Offset(int(pageIndex * input.PageSize)).Limit(int(input.PageSize)).Joins("left join users on videos.user_id = users.id").Joins("left join channels on videos.channel_id = channels.id").Model(&models.Video{}).Find(&videos).Error

	if err != nil {
		return videos, &commonModels.RepositoryError{Code: commonModels.RepositoryErrorCode_Fail}
	}
	return videos, nil
}

// type VideoItemList struct {
// 	models.Video
// 	UserFullname     string
// 	UserAvatar       string
// 	ChannelName      string
// 	ChannelThumbnail string
// }

// ID           string `gorm:"primaryKey;"`
// UserId       string `gorm:"type:varchar(255);"`
// ChannelID    string `gorm:"type:varchar(255);"`
// Name         string `gorm:"type:varchar(255);"`
// Image        string `gorm:"type:varchar(255);"`
// Description  string `gorm:"type:text;"`
// SourceType   string `gorm:"type:varchar(100);"`
// Source       string `gorm:"type:varchar(255);"`
