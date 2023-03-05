package route

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"com.vandong9.clone_youtube_golang_api/modules/user/models"
	"com.vandong9.clone_youtube_golang_api/modules/user/transport"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	// Migrate
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserToken{})

	// handle path
	groupRoute := route.Group("/api/v1", func(ctx *gin.Context) {
		ctx.Set(constant.Context_ID, "User-context-"+uuid.New().String())
	})
	groupRoute.POST("/register", transport.CreateUser(db))
	groupRoute.POST("/update", transport.UpdateUser(db))
	groupRoute.POST("/login", transport.Login(db))

}
