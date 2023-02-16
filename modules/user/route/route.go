package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"com.vandong9.clone_youtube_golang_api/modules/user/transport"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", transport.CreateUser(db))

}
