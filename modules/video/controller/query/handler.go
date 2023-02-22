package query

import (
	"net/http"

	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func HandleQueryVideo(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input QueryInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		var validate = validator.New()
		if err := validate.Struct(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		repo := CreateQueryRepository(db)
		service := CreateQueryVideoService(&repo)
		videos, err := service.QueryVideo(input)
		if err != commonModels.ServiceErrorCode_OK {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}
		utils.Response(ctx, http.StatusOK, nil, videos)

	}
}
