package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateUserStorageInPostgres(config Config) (*UserStorage, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=123456 dbname=CloneYoutube port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sql := UserStorage{DB: db}
	return &sql, nil
}
