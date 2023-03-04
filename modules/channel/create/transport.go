package create

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/common/constant"
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

		validate := validator.New()
		err = validate.Struct(input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, nil)
			return
		}

		repo := CreateRepository(db)
		service := CreateChannelServiceInstance(&repo)

		userID := ctx.GetString(constant.Header_User_ID_Key)
		if len(userID) == 0 {
			utils.ResponseForbidden(ctx, nil)
			return
		}

		input.UserId = userID
		channel, err := service.CreateChannel(input)
		if err != nil {
			utils.ReponseBadRequest(ctx, err)
			return
		}

		utils.ReponseSuccess(ctx, channel)

	}
}
