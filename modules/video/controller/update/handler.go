package update

import (
	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func HandleUpdateVideo(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		videoID := ctx.Param("video_id")
		if len(videoID) == 0 {
			utils.ReponseBadRequest(ctx, nil)
			return
		}

		var input UpdateVideoInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.ReponseBadRequest(ctx, err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(input); err != nil {
			utils.ReponseBadRequest(ctx, err)
			return
		}

		repo := CreateUpdateVideRepository(db)
		service := CreateUpdateVideService(&repo)
		if err := service.UpdateVideoByInput(input); err != nil {
			utils.ReponseBadRequest(ctx, err)
			return
		}

		utils.ReponseSuccess(ctx, nil)
	}
}
