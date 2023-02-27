package create

type CreateChannelInput struct {
	UserId      string `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Description string `json:"description" validate:"required"`
}
