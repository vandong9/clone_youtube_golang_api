package query

import (
	"net/http"

	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HandleQueryComment(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input QueryCommentInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		var validate = validator.New()
		if err := validate.Struct(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		repo := CreateQueryCommentRepository(db)
		service := CreateQueryCommentService(&repo)
		comments, err := service.QueryComment(input)
		if err != commonModels.ServiceErrorCode_OK {
			utils.Response(ctx, http.StatusBadRequest, err, []models.Comment{})
			return
		}

		utils.Response(ctx, http.StatusOK, nil, comments)

	}
}
