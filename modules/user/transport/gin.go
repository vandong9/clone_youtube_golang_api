package transport

import (
	"fmt"
	"net/http"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var input models.LoginInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var validate *validator.Validate
		validate = validator.New()
		err := validate.Struct(input)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "request params error: " + err.Error()})
			return
		}

		uc := usecase.CreateUserUsecase(db)
		user, err := uc.Login(ctx, &input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "login error:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"body": &user})
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
