package middleware

import (
	"fmt"
	"net/http"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/user/usecase"
	"com.vandong9.clone_youtube_golang_api/utils"
	"com.vandong9.clone_youtube_golang_api/utils/jwt_handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MiddlewareCheckValidToken(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// verify token
		token, err := jwt_handler.ExtractClaims(ctx.Writer, ctx.Request)
		if err != nil || token == nil {
			utils.Response(ctx, http.StatusForbidden, err, nil)
			ctx.Abort()
			return
		}

		// get user from token
		tokenString := ctx.Request.Header[constant.Request_Header_Token_Key][0]
		userService := usecase.CreateUserUsecase(db)
		userID := userService.GetUserIDByGivenToken(ctx, tokenString)
		if userID == nil {
			utils.Response(ctx, http.StatusForbidden, err, nil)
			ctx.Abort()
			return
		}
		fmt.Println("Stored user id into header " + *userID)
		ctx.Header(constant.Header_User_ID_Key, *userID)
	}
}
