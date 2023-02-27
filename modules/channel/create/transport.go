package create

import (
	"fmt"
	"net/http"

	"com.vandong9.clone_youtube_golang_api/utils"
	jwt_handler "com.vandong9.clone_youtube_golang_api/utils/jwt_handler"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateChannelHandler(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		token, err := jwt_handler.ExtractClaims(ctx.Writer, ctx.Request)
		fmt.Println("receive toke:  " + token.UserFullName)
		if err != nil || token == nil {
			utils.Response(ctx, http.StatusForbidden, err, nil)
			return
		}

		var input CreateChannelInput
		err = ctx.ShouldBindJSON(&input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		validate := validator.New()
		err = validate.Struct(input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		repo := CreateRepository(db)
		service := CreateChannelServiceInstance(&repo)
		service.CreateChannel(input)
	}
}
