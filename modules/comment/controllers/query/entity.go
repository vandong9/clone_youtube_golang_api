package query

type QueryCommentInput struct {
	VideoID   string `json:"video_id"`
	CommentID string `json:"comment_id"`
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
}
