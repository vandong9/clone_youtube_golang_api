package comment

import (
	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/middleware"
	"com.vandong9.clone_youtube_golang_api/modules/comment/controllers/create"
	"com.vandong9.clone_youtube_golang_api/modules/comment/controllers/query"
	"com.vandong9.clone_youtube_golang_api/modules/comment/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InitCommentRoutes(db *gorm.DB, route *gin.Engine) {
	db.AutoMigrate(&models.Comment{})

	group := route.Group("api/v1/comment", func(ctx *gin.Context) {
		ctx.Set(constant.Context_ID, "Comment-context-"+uuid.New().String())
	})
	group.POST("/create", middleware.MiddlewareCheckValidToken(db), create.HandleCreateComment(db))
	group.GET("/", query.HandleQueryComment(db))
}
