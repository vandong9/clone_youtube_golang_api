package utils

import (
	"github.com/gin-gonic/gin"
)

func Reponse(ctx *gin.Context, status int, err interface{}, data interface{}) {
	ctx.JSON(status, gin.H{"error": err, "data": data})
}
