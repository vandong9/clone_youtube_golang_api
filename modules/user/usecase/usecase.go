package usecase

import (
	"context"
	"fmt"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/storage"

	"gorm.io/gorm"
)

type IUserStorage interface {
	CreateUser(ctx context.Context, data *models.User) error
	UpdateUser(ctx context.Context, data *models.UpdateInput) (user *models.User, err_code string)
	GetUserByIDAndPassword(ctx context.Context, id string, password string) (user models.User, err error)
}

type UserUsecase struct {
	storage IUserStorage
}

func CreateUserUsecase(db *gorm.DB) *UserUsecase {
	store := storage.CreateUserStorage(db)
	uc := UserUsecase{storage: &store}
	return &uc
}

func (uc *UserUsecase) Login(ctx context.Context, data *models.LoginInput) (user models.User, err error) {
	user, err = uc.storage.GetUserByIDAndPassword(ctx, data.Name, data.Password)
	if err != nil {
		err = &models.LoginError{
			Code:    "Login_Error_1",
			Title:   "title 1",
			Content: "content 1"}
	}
	return
}

func (uc *UserUsecase) CreateUser(ctx context.Context, data *models.RegisterInput) error {
	var user models.User
	user.Username = data.Username
	user.Fullname = data.Fullname
	user.Password = data.Password
	user.Email = data.Email
	fmt.Println(data)
	err := uc.storage.CreateUser(ctx, &user)
	return err
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, data *models.UpdateInput) (*models.User, *models.UpdateError) {
	user, err_code := uc.storage.UpdateUser(ctx, data)
	if len(err_code) > 0 {
		return nil, &models.UpdateError{Code: err_code, Title: "", Content: ""}
	}
	return user, nil
}
