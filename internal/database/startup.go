package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB_ENV string

const (
	DEV  DB_ENV = "dev"
	PROD DB_ENV = "prod"
)

type Config struct {
	Host        string
	User        string
	Password    string
	DbName      string
	SslMode     string
	TimeZone    string
	Environment DB_ENV
	Port        uint
}

func InitDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.DbName,
		config.Port,
		config.SslMode,
		config.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: config.Environment == PROD,
	})
	if err != nil {
		panic(err)
	}

	// TODO: Add logging
	fmt.Println("Database connection established")
	db.AutoMigrate(&SensorUnit{}, &Bike{}, &Trip{}, &DataPoint{})
	fmt.Println("Database schema migrated")

	return db
}
