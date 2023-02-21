package create

type CreateChannelInput struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}
