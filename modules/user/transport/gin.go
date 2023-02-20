package transport

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/usecase"
	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var input models.LoginInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, "")
			return
		}
		var validate *validator.Validate
		validate = validator.New()
		err := validate.Struct(input)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		user := models.RegisterInput{}
		uc := usecase.CreateUserUsecase(db)
		err := uc.CreateUser(ctx, &user)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}
		utils.Response(ctx, http.StatusOK, nil, user)
	}
}

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		//Serialize input
		var input models.UpdateInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		// validate input
		var validate *validator.Validate
		validate = validator.New()
		err := validate.Struct(input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		// Process update
		uc := usecase.CreateUserUsecase(db)
		updateUser, errUpdate := uc.UpdateUser(ctx, &input)
		if errUpdate != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
		} else {
			utils.Response(ctx, http.StatusOK, nil, updateUser)
		}
	}
}
