package video

import (
	"com.vandong9.clone_youtube_golang_api/modules/video/controller/create"
	"com.vandong9.clone_youtube_golang_api/modules/video/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Video{})

	groupRoute := route.Group("api/v1/video")
	groupRoute.POST("/create", create.HandleCreateVideo(db))
}
