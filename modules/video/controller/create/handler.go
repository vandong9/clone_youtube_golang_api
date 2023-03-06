package create

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/utils"
	"com.vandong9.clone_youtube_golang_api/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func HandleCreateVideo(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input CreateVideoInput
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			logger.LogInfo(ctx, "transport - parse input error: "+err.Error())
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		validator := validator.New()
		err = validator.Struct(&input)
		if err != nil {
			logger.LogInfo(ctx, "transport - validate input error: "+err.Error())
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		userID := ctx.GetString(constant.Header_User_ID_Key)
		if len(userID) < 10 {
			logger.LogInfo(ctx, "transport - missing userid error: "+err.Error())
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		input.UserId = userID

		repo := CreateAddVideoRepository(db)
		service := InitCreateVideoService(&repo)
		video, er := service.CreateVideo(ctx, input)
		if er != nil {
			logger.LogInfo(ctx, "transport - business return error: "+er.Error())
			utils.ReponseBadRequest(ctx, er)
			return
		}

		utils.ReponseSuccess(ctx, video)
	}
}
