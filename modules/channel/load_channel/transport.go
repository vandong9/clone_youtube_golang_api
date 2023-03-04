package load_channel

import (
	"fmt"
	"net/http"

	"com.vandong9.clone_youtube_golang_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func QueryChannelHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input QueryInput

		err := ctx.ShouldBindQuery(&input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, []string{})
			return
		}

		validate := validator.New()
		err = validate.Struct(input)
		if err != nil {
			utils.Response(ctx, http.StatusBadRequest, err, []string{})
			return
		}

		fmt.Println("Input query")
		fmt.Println(input)
		repo := CreateQueryChannelRepo(db)
		service := CreateQueryService(&repo)
		response := service.QueryChannel(input)
		utils.ReponseSuccess(ctx, response)

	}
}
