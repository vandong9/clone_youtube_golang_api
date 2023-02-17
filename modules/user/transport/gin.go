package transport

import (
	"fmt"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var input models.LoginInput
		ctx.ShouldBindJSON(&input)

		uc := usecase.CreateUserUsecase(db)
		uc.Login(ctx, &input)
	}
}

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		user := models.User{}
		uc := usecase.CreateUserUsecase(db)
		err := uc.CreateUser(ctx, &user)
		if err != nil {
			fmt.Print(err)
		}
	}
}

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

	}
}
