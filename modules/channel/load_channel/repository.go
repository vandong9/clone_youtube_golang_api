package load_channel

import (
	"math"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"gorm.io/gorm"
)

type QueryChannelRepository struct {
	db *gorm.DB
}

func CreateQueryChannelRepo(db *gorm.DB) QueryChannelRepository {

	return QueryChannelRepository{db: db}
}

func (repo *QueryChannelRepository) QueryChannel(input QueryInput) []models.Channel {
	var channels []models.Channel
	input.PageIndex -= 1
	pageIndex := int(math.Max(float64(input.PageIndex), 0))
	repo.db.Offset(pageIndex * input.PageSize).Limit(int(math.Min(constant.Max_Channel_Page_Size, float64(input.PageSize)))).Model(&models.Channel{}).Find(&channels)
	return channels
}
