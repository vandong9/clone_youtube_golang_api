package video

import (
	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/middleware"
	"com.vandong9.clone_youtube_golang_api/modules/video/controller/create"
	"com.vandong9.clone_youtube_golang_api/modules/video/controller/query"
	"com.vandong9.clone_youtube_golang_api/modules/video/controller/update"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"com.vandong9.clone_youtube_golang_api/utils/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Video{})

	groupRoute := route.Group("api/v1/video", func(ctx *gin.Context) {
		ctx.Set(constant.Context_Logger_Object, logger.CreateContextLogger(ctx))
	})
	groupRoute.POST("/create", middleware.MiddlewareCheckValidToken(db), create.HandleCreateVideo(db))
	groupRoute.POST("/:video_id", middleware.MiddlewareCheckValidToken(db), update.HandleUpdateVideo(db))

	groupRoute.GET("/", query.HandleQueryVideo(db))
}
