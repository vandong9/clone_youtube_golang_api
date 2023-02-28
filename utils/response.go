package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, status int, err interface{}, data interface{}) {
	ctx.JSON(status, gin.H{"error": err, "data": data})
}

func ReponseBadRequest(ctx *gin.Context, err interface{}) {
	Response(ctx, http.StatusBadRequest, nil, nil)
}

func ResponseForbidden(ctx *gin.Context, err interface{}) {
	Response(ctx, http.StatusForbidden, err, nil)
}

func ReponseSuccess(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, nil, data)
}
