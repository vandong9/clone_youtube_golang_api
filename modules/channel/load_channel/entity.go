package load_channel

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
)

type QueryInput struct {
	UserID    string `json:"user_id"`
	PageIndex uint   `json:"page_index"`
	PageSize  uint   `json:"page_size"`
}

type QueryResponse struct {
	Channels []models.Channel `json:"channels"`
}
