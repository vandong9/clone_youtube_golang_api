package channel

import (
	"com.vandong9.clone_youtube_golang_api/middleware"
	"com.vandong9.clone_youtube_golang_api/modules/channel/create"
	"com.vandong9.clone_youtube_golang_api/modules/channel/load_channel"
	"com.vandong9.clone_youtube_golang_api/modules/channel/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Channel{})

	groupRoute := route.Group("/api/v1/channel")
	groupRoute.POST("/create", middleware.MiddlewareCheckValidToken(db), create.CreateChannelHandler(db))
	groupRoute.GET("/", load_channel.QueryChannelHandler(db))
}
