package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
)

type QueryInput struct {
	commonModels.Paging
	ChannelID string `json:"channel_id"`
}
