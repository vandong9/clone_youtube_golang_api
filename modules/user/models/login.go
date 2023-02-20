package models

import "encoding/json"

type LoginInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type LoginResponse struct {
	User  User
	Token string
}

type LoginError struct {
	Code    string
	Title   string
	Content string
}

func (le *LoginError) Error() string {
	errModel, err := json.Marshal(le)
	if err != nil {
		return "error"
	}
	return string(errModel)
}
