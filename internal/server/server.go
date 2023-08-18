package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	router "github.com/zatrano/zatrano/internal/app/routes"
)

func Server() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
