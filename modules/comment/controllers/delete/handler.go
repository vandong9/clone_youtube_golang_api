package delete

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HandleDeleteComment(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input DeleteCommentInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		var validate = validator.New()
		if err := validate.Struct(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		repo := CreateDeleteCommentRepository(db)
		service := CreateDeleteCommentService(&repo)
		err := service.DeleteComment(input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		utils.Response(ctx, http.StatusOK, nil, nil)

	}
}
