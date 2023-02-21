package create

type CreateVideoInput struct {
	UserId      string `json:"user_id" validate:"required"`
	ChannelID   string `json:"channel_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Image       string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	SourceType  string `json:"source_type" validate:"required"`
	Source      string `json:"source" validate:"required"`
}
