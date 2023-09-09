package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")

}
