package storage

import (
	"context"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"gorm.io/gorm"
)

type UserStorage struct {
	DB *gorm.DB
}

func CreateUserStorage(db *gorm.DB) UserStorage {
	return UserStorage{DB: db}
}

func (s *UserStorage) GetUserByID(ctx context.Context, id string) (user models.User, err error) {
	err = s.DB.First(&user, "ID = ?", id).Error
	return
}

func (s *UserStorage) GetUserByIDAndPassword(ctx context.Context, id string, password string) (user models.User, err error) {
	err = s.DB.Where("ID = ? AND Password= ?", id, password).Find(&user).Error
	return
}

func (s *UserStorage) CreateUser(ctx context.Context, data *models.User) error {
	if err := s.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, data *models.UpdateUserRequest) error {
	if err := s.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
