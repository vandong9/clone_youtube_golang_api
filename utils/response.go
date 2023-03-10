package utils

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/utils/logger"
	"github.com/gin-gonic/gin"
)

func GenerateAccessToken() *string {
	return nil
}

/*
***** NOTE all response MUST  use  Response(ctx *gin.Context, status int, err interface{}, data interface{})
***** to log the output
 */
func Response(ctx *gin.Context, status int, err interface{}, data interface{}) {
	logger.PrintLog(ctx)
	ctx.JSON(status, gin.H{"error": err, "data": data})
}

func ReponseBadRequest(ctx *gin.Context, err interface{}) {
	Response(ctx, http.StatusBadRequest, nil, err)
}

func ResponseForbidden(ctx *gin.Context, err interface{}) {
	Response(ctx, http.StatusForbidden, err, nil)
}

func ReponseSuccess(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, nil, data)
}
