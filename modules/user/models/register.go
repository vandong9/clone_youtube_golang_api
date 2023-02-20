package models

type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=12"`
	Fullname string `json:"fullname"`
	Email    string `json:"email" validate:"required"`
}
