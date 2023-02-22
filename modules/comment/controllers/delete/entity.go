package delete

type DeleteCommentInput struct {
	CommentID string `json:"comment_id" validate:"required"`
}
