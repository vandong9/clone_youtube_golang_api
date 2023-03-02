package video

import (
	"net/http"

	"com.vandong9.clone_youtube_golang_api/modules/video/controller/create"
	"com.vandong9.clone_youtube_golang_api/modules/video/controller/query"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"com.vandong9.clone_youtube_golang_api/utils"
	"com.vandong9.clone_youtube_golang_api/utils/jwt_handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Video{})

	groupRoute := route.Group("api/v1/video")
	groupRoute.POST("/create", CheckToken(), create.HandleCreateVideo(db))
	groupRoute.GET("/", query.HandleQueryVideo(db))
}

func CheckToken() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token, err := jwt_handler.ExtractClaims(ctx.Writer, ctx.Request)
		if err != nil || token == nil {
			utils.Response(ctx, http.StatusForbidden, err, nil)
			return
		}

	}
}
