package channel

import (
	"com.vandong9.clone_youtube_golang_api/modules/channel/create"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&create.Channel{})

	groupRoute := route.Group("/api/v1/channel")
	groupRoute.POST("/create", create.CreateChannelHandler(db))
}
