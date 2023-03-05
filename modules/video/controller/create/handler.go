package create

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func HandleCreateVideo(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input CreateVideoInput
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		validator := validator.New()
		err = validator.Struct(&input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		repo := CreateAddVideoRepository(db)
		service := InitCreateVideoService(&repo)
		video, er := service.CreateVideo(ctx, input)
		if er != nil {
			utils.ReponseBadRequest(ctx, er)
			return
		}

		utils.ReponseSuccess(ctx, video)
	}
}
