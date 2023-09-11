package main

import (
	routes "goproj/Routes"
	database "goproj/databse"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	database.Connect() //used the connect.go function
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	} //did this to run app on a port

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port) //port listening 

	routes.Setup(app)
}
