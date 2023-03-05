package channel

import (
	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/middleware"
	"com.vandong9.clone_youtube_golang_api/modules/channel/create"
	"com.vandong9.clone_youtube_golang_api/modules/channel/load_channel"
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Channel{})

	groupRoute := route.Group("/api/v1/channel", func(ctx *gin.Context) {
		ctx.Set(constant.Context_ID, "Comment-context-"+uuid.New().String())
	})
	groupRoute.POST("/create", middleware.MiddlewareCheckValidToken(db), create.CreateChannelHandler(db))
	groupRoute.GET("/", load_channel.QueryChannelHandler(db))
}
