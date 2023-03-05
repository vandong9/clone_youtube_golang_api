package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"com.vandong9.clone_youtube_golang_api/utils/math_util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryChannelRepository struct {
	db *gorm.DB
}

func CreateQueryChannelRepo(db *gorm.DB) QueryChannelRepository {

	return QueryChannelRepository{db: db}
}

func (repo *QueryChannelRepository) GetDB() *gorm.DB {
	return repo.db
}

func (repo *QueryChannelRepository) QueryChannel(input QueryInput) []models.Channel {
	var channels []models.Channel
	input.PageIndex -= 1
	pageIndex := int(math_util.MaxOfInt(input.PageIndex, 0))
	repo.db.Offset(int(pageIndex * input.PageSize)).Limit(int(math_util.MinOfInt(constant.Max_Channel_Page_Size, input.PageSize))).Model(&models.Channel{}).Find(&channels)
	return channels
}

func (repo *QueryChannelRepository) GetChannelByID(ctx *gin.Context, channelID string) *models.Channel {
	var channel models.Channel
	if err := repo.db.Debug().Where("ID = ?", channelID).First(&channel).Error; err != nil {
		return nil
	}

	return &channel
}
