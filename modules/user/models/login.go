package models

type LoginInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  User
	Token string
}
