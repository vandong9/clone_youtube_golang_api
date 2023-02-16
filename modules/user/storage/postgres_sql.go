package storage

import (
	"com.vandong9.clone_youtube_golang_api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateUserStorageInPostgres(config config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Postgres_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
