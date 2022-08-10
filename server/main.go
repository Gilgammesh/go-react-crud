package main

import (
	"log"
	"os"

	"github.com/Gilgammesh/go-react-crud/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize fiber app
	app := fiber.New()

	// Middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	// Router
	router.Routes(app)

	// Run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
