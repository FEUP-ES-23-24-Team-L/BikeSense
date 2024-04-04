package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	dbApi "bikesense-web/internal/database"
	server "bikesense-web/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GIN_MODE") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("[BKS MAIN - WARN]Error loading .env file")
		}
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("[BKS MAIN - WARN] Error parsing db port, Trying default port 5432")
		port = 5432
	}

	config := dbApi.Config{
		Host:        os.Getenv("DB_HOST"),
		User:        os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		DbName:      os.Getenv("DB_NAME"),
		SslMode:     os.Getenv("DB_SSLMODE"),
		TimeZone:    os.Getenv("DB_TIMEZONE"),
		Port:        uint(port),
		Environment: dbApi.PROD,
	}

	if err := config.Validate(); err != nil {
		log.Fatalf("[BKS MAIN - ERROR] Invalid database configuration: %v", err)
	}

	server.Run(dbApi.OpenAndMigrateDB(config))
}
