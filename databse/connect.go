package database

import (
	"goproj/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //pointer to the gorm object for db connection

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN") // env usage
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to DB")
	} else {
		log.Println("Connection Succesful")
	} //this entire block is for db connection

	DB = database
	//imported models to database
	database.AutoMigrate(
		&models.User{},
	)
}
