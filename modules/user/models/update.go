package models

type UpdateInput struct {
	ID       string `json:"id" validate:"required"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UpdateError struct {
	Code    string
	Title   string
	Content string
}
