package create

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/utils/logger"

	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HandleCreateComment(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input CreateCommentInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			logger.LogInfo(ctx, "transport - business return error: "+err.Error())
			utils.ReponseBadRequest(ctx, err)
			return
		}

		var validate = validator.New()
		if err := validate.Struct(&input); err != nil {
			logger.LogInfo(ctx, "transport - validate error: "+err.Error())
			utils.ReponseBadRequest(ctx, err)
			return
		}

		userID := ctx.GetString(constant.Header_User_ID_Key)
		if len(userID) < 10 {
			logger.LogInfo(ctx, "transport - missing userid error: ")
			utils.ResponseForbidden(ctx, nil)
			return
		}

		input.UserId = userID

		repo := InitCreateCommentRepository(db)
		service := InitCreateCommentService(&repo)

		comment, err := service.CreateComment(input)
		if err != commonModels.ServiceErrorCode_OK {
			logger.LogInfo(ctx, "transport - business return error: "+err.String())
			utils.ReponseBadRequest(ctx, err)
			return
		}

		utils.Response(ctx, http.StatusOK, nil, comment)
	}
}
