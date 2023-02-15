package usecase

import (
	"context"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
)

type IUserStorage interface {
	CreateUser(ctx context.Context, data *models.User) error
}

type UserUsecase struct {
	storage IUserStorage
}

func CreateUserUsecase(storage IUserStorage) *UserUsecase {
	uc := UserUsecase{storage: storage}
	return &uc
}

func (uc *UserUsecase) createUser(ctx context.Context, data *models.User) error {
	err := uc.storage.CreateUser(ctx, data)
	return err
}
