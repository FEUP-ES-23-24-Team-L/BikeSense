package database

import (
	"fmt"
	"log"

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

func (c Config) GetFullDsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host,
		c.User,
		c.Password,
		c.DbName,
		c.Port,
		c.SslMode,
		c.TimeZone,
	)
}

func (c Config) GetDsnNoDBName() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s port=%d sslmode=%s TimeZone=%s",
		c.Host,
		c.User,
		c.Password,
		c.Port,
		c.SslMode,
		c.TimeZone,
	)
}

func OpenAndMigrateDB(config Config) *gorm.DB {
	dsn := config.GetFullDsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: config.Environment == PROD,
	})
	if err != nil {
		panic(err)
	}

	// TODO: Add logging
	log.Println("Database connection established")
	db.AutoMigrate(&SensorUnit{}, &Bike{}, &Trip{}, &DataPoint{})
	log.Println("Database schema migrated")

	return db
}
