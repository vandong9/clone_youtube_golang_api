package transport
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"com.vandong9.clone_youtube_golang_api/modules/user/usecase"
	"com.vandong9.clone_youtube_golang_api/modules/user/storage"
	"com.vandong9.clone_youtube_golang_api/config"
	"com.vandong9.clone_youtube_golang_api/modules/user/models"

)

func CreateUser() (func(*gin.Context)) {
	return func(ctx *gin.Context) {
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Fatal(err)
			return
		}

		user := models.User{}
		storage := storage.CreateUserStorageInPostgres(config)
		uc := usecase.CreateUserUsecase(storage)
		err := uc.CreateUser(ctx, user)
		if err != nil {
			fmt.Print(err)
		}
	}
}