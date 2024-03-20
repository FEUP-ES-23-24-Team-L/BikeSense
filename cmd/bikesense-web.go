package main

import (
	"fmt"
	"os"
	"strconv"

	dbApi "bikesense-web/internal/database"
	server "bikesense-web/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Error parsing port, Trying default port 5432")
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

	server.Run(dbApi.InitDB(config))
}
