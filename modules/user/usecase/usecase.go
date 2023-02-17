package usecase

import (
	"context"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/storage"

	"gorm.io/gorm"
)

type IUserStorage interface {
	CreateUser(ctx context.Context, data *models.User) error
	UpdateUser(ctx context.Context, data *models.UpdateUserRequest) error
	GetUserByIDAndPassword(ctx context.Context, id string, password string) (user models.User, err error)
}

type UserUsecase struct {
	storage storage.UserStorage
}

func CreateUserUsecase(db *gorm.DB) *UserUsecase {
	uc := UserUsecase{storage: storage.CreateUserStorage(db)}
	return &uc
}

func (uc *UserUsecase) Login(ctx context.Context, data *models.LoginInput) {
	user, err := uc.storage.GetUserByIDAndPassword(ctx, data.Name, data.Password)
}

func (uc *UserUsecase) CreateUser(ctx context.Context, data *models.User) error {
	err := uc.storage.CreateUser(ctx, data)
	return err
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, data *models.UpdateUserRequest) error {
	err := uc.storage.UpdateUser(ctx, data)
	return err

}
