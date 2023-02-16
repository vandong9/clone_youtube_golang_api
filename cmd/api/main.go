package main

import (
	"fmt"
	"log"

	"com.vandong9.clone_youtube_golang_api/config"
	userRoute "com.vandong9.clone_youtube_golang_api/modules/user/route"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("connect db:" + config.Postgres_URL)
	fmt.Println("server port:" + config.APIPort)
	// db, err := gorm.Open(sqlite.Open(config.Postgres_URL), &gorm.Config{})

	router := SetupRouter()
	log.Fatal(router.Run(":" + "3001"))
}

func SetupRouter() *gin.Engine {
	dns := "host=localhost user=postgres password=123456 dbname=clone_youtube port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	userRoute.InitAuthRoutes(db, router)

	return router

}
