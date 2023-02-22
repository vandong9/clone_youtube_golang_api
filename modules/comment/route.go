package comment

import (
	"com.vandong9.clone_youtube_golang_api/modules/comment/controllers/create"
	"com.vandong9.clone_youtube_golang_api/modules/comment/controllers/query"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Comment{})

	group := route.Group("api/v1/commment")
	group.POST("/create", create.HandleCreateComment(db))
	group.GET("/", query.HandleQueryComment(db))

}
