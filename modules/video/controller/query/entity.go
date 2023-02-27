package query

type QueryInput struct {
	ChannelID string `json:"channel_id"`
	PageIndex uint   `json:"page_index"`
	PageSize  uint   `json:"page_size"`
}
