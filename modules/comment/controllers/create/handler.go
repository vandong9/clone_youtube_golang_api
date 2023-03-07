package create

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/common/constant"
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

		if len(input.CommentID) == 0 && len(input.VideoID) == 0 {
			logger.LogInfo(ctx, "transport - validate error: ")
			utils.ReponseBadRequest(ctx, "validate error missing field")
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
		if err != nil {
			logger.LogInfo(ctx, "transport - business return error: "+err.Error())
			utils.ReponseBadRequest(ctx, err)
			return
		}

		utils.Response(ctx, http.StatusOK, nil, comment)
	}
}
