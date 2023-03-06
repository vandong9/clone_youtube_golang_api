package create

type CreateCommentInput struct {
	UserId    string `json:"user_id"`
	VideoID   string `json:"video_id"`
	CommentID string `json:"comment_id"`
	Content   string `json:"content" validate:"required"`
}
