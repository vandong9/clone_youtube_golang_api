package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
)

type QueryInput struct {
	commonModels.PagingQuery
	ChannelID string `form:"channel_id"`
}

type QueryResponse struct {
	models.Video
	UserAvatar       string
	ChannelThumbnail string
}

type VideoItemList struct {
	models.Video
	UserFullname     string
	UserAvatar       string
	ChannelName      string
	ChannelThumbnail string
}
