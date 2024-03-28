package tests

import (
	dbApi "bikesense-web/internal/database"
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createTestDB(config dbApi.Config) error {
	dsn := config.GetDsnNoDBName()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// create the database
	// this is a very simple way to create a database
	createDBCommand := fmt.Sprintf("CREATE DATABASE %s", config.DbName)
	if err = db.Exec(createDBCommand).Error; err != nil {
		return err
	}

	return nil
}

func GetTestDB(t *testing.T) *gorm.DB {
	dbName := fmt.Sprintf("bikesense_test_%d", time.Now().UnixMicro())

	// create a new database for testing
	config := dbApi.Config{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgrespw",
		DbName:   dbName,
		SslMode:  "disable",
		TimeZone: "Europe/Lisbon",
		Port:     5432,
	}

	if err := createTestDB(config); err != nil {
		t.Fatalf("Error creating test database: %v", err)
	}

	t.Logf("Test database created: %s", dbName)

	return dbApi.OpenAndMigrateDB(config)
}
