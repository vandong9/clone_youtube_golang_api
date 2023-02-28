package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"com.vandong9.clone_youtube_golang_api/utils/math_util"
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
	pageIndex := uint(math_util.MaxOfUInt(input.PageIndex, 0))
	repo.db.Offset(int(pageIndex * input.PageSize)).Limit(int(math_util.MaxOfUInt(constant.Max_Channel_Page_Size, input.PageSize))).Model(&models.Channel{}).Find(&channels)
	return channels
}
