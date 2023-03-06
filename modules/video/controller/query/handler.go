package query

import (
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/utils"
	"com.vandong9.clone_youtube_golang_api/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func HandleQueryVideo(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input QueryInput
		if err := ctx.ShouldBindQuery(&input); err != nil {
			logger.LogInfo(ctx, "transport - parse error: "+err.Error())
			utils.ReponseBadRequest(ctx, err)
			return
		}

		var validate = validator.New()
		if err := validate.Struct(&input); err != nil {
			logger.LogInfo(ctx, "transport - validate error: "+err.Error())
			utils.ReponseBadRequest(ctx, err)
			return
		}

		repo := CreateQueryRepository(db)
		service := CreateQueryVideoService(&repo)
		videos, err := service.QueryVideo(input)
		if err != commonModels.ServiceErrorCode_OK {
			logger.LogInfo(ctx, "transport - service reponse error: "+err.String())
			utils.ReponseBadRequest(ctx, err)
			return
		}
		utils.ReponseSuccess(ctx, videos)

	}
}
