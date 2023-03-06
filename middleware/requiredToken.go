package middleware

import (
	"fmt"
	"net/http"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/user/usecase"
	"com.vandong9.clone_youtube_golang_api/utils"
	"com.vandong9.clone_youtube_golang_api/utils/jwt_handler"
	"com.vandong9.clone_youtube_golang_api/utils/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MiddlewareCheckValidToken(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// verify token
		token, err := jwt_handler.ExtractClaims(ctx.Writer, ctx.Request)
		if err != nil || token == nil {
			errMsg := "Header missing Token key"
			if err != nil {
				errMsg = err.Error()
			}
			logger.LogInfo(ctx, "Middle-CheckToken validate error: "+errMsg)
			utils.Response(ctx, http.StatusForbidden, err, nil)
			ctx.Abort()
			return
		}

		// get user from token
		tokenString := ctx.Request.Header[constant.Request_Header_Token_Key][0]
		userService := usecase.CreateUserUsecase(db)
		userID := userService.GetUserIDByGivenToken(ctx, tokenString)
		if userID == nil {
			logger.LogInfo(ctx, "Middle-CheckToken query userid error : "+err.Error())
			utils.Response(ctx, http.StatusForbidden, err, nil)
			ctx.Abort()
			return
		}
		fmt.Println("Stored user id into header " + *userID)
		ctx.Set(constant.Header_User_ID_Key, *userID)
		ctx.Next()
	}
}
