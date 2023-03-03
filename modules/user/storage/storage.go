package storage

import (
	"context"
	"time"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"gorm.io/gorm"
)

type UserStorage struct {
	DB *gorm.DB
}

func CreateUserStorage(db *gorm.DB) UserStorage {
	return UserStorage{DB: db}
}

func (s *UserStorage) GetUserIDByGivenToken(ctx context.Context, token string) *string {
	var userToken models.UserToken
	err := s.DB.Debug().Where("token == ?", token).First(&userToken).Error
	if err != nil {
		return nil
	}

	return &userToken.UserID
}

func (s *UserStorage) UpdateUserToken(ctx context.Context, userID string, token string) error {
	userToken := models.UserToken{}
	userToken.UserID = userID
	userToken.Token = token
	err := s.DB.Debug().Save(&userToken).Error
	if err != nil {
		err = s.DB.Debug().Create(&userToken).Error
	}
	return err
}

func (s *UserStorage) GetUserByID(ctx context.Context, id string) (user models.User, err error) {
	err = s.DB.First(&user, "ID = ?", id).Error
	return
}

func (s *UserStorage) GetUserByIDAndPassword(ctx context.Context, username string, password string) (user models.User, err error) {
	err = s.DB.Where("Username = ? AND Password= ?", username, password).First(&user).Error
	return
}

func (s *UserStorage) CreateUser(ctx context.Context, data *models.User) error {
	if err := s.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, data *models.UpdateInput) (*models.User, string) {
	// get user by id
	var user models.User
	s.DB.Model(&user)

	user.ID = data.ID
	userID := s.DB.Debug().Select("*").Where("id = ?", data.ID).Find(&user)
	if userID.RowsAffected < 1 {
		return nil, constant.Storage_record_not_found
	}

	// update field that have value update
	user.Fullname = data.Fullname
	user.Email = data.Email
	user.UpdatedAt = time.Now()
	updateUser := s.DB.Save(&user)

	if updateUser.Error != nil {
		return nil, constant.Storage_update_record_fail
	}

	return &user, ""
}
