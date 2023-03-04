package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
)

type QueryInput struct {
	UserID    string `form:"user_id"`
	PageIndex int    `form:"page_index"`
	PageSize  int    `form:"page_size"`
}

type QueryResponse struct {
	Channels []models.Channel `json:"channels"`
}
