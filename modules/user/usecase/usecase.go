package usecase

import (
	"context"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/storage"

	"gorm.io/gorm"
)

type IUserStorage interface {
	CreateUser(ctx context.Context, data *models.User) error
}

type UserUsecase struct {
	storage storage.UserStorage
}

func CreateUserUsecase(db *gorm.DB) *UserUsecase {
	uc := UserUsecase{storage: storage.CreateUserStorage(db)}
	return &uc
}

func (uc *UserUsecase) CreateUser(ctx context.Context, data *models.User) error {
	err := uc.storage.CreateUser(ctx, data)
	return err
}
