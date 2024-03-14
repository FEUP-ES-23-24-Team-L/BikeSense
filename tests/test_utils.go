package tests

import (
	dbApi "bikesense-web/internal/database"

	"gorm.io/gorm"
)

func GetTestDB() *gorm.DB {
	config := dbApi.Config{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgrespw",
		DbName:   "bikesense",
		SslMode:  "disable",
		TimeZone: "Europe/Lisbon",
		Port:     5432,
	}

	return dbApi.InitDB(config)
}
