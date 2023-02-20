package create

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateChannelHandler(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var input CreateChannelInput
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		var validate *validator.Validate
		validate = validator.New()
		err = validate.Struct(input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		service := CreateChannelServiceInstance(db)
		service.CreateChannel(input)
	}
}
