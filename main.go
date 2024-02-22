package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mysterybee07/blogbackend/database"
	"github.com/mysterybee07/blogbackend/routes"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
