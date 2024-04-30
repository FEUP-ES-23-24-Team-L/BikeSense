package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func (c Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.User == "" {
		return fmt.Errorf("user is required")
	}
	if c.Password == "" {
		return fmt.Errorf("password is required")
	}
	if c.DbName == "" {
		return fmt.Errorf("dbName is required")
	}
	if c.SslMode == "" {
		return fmt.Errorf("sslMode is required")
	}
	if c.TimeZone == "" {
		return fmt.Errorf("timeZone is required")
	}
	return nil
}

func OpenAndMigrateDB(config Config) *gorm.DB {
	dsn := config.GetFullDsn()

	log.Println("[DB Startup] Connecting to database name: ", config.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: config.Environment == PROD,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", config.Environment),
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatalf("[DB Startup] Error connecting to database: %v", err)
	}
	log.Println("[DB Startup] Database connection established")

	if err := db.AutoMigrate(&SensorUnit{}, &Bike{}, &Trip{}, &DataPoint{}); err != nil {
		log.Fatalf("[DB Startup] Error migrating database schema: %v", err)
	}
	log.Println("[DB Startup] Database schema migrated")

	return db
}
