package server

import (
	"log"

	router "zatrano/internal/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
