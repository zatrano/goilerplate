package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zatrano/zatrano/internal/app/handlers"
)

func Server() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/books", handlers.GetAllBooks)
	app.Get("/books/:id", handlers.GetBookByID)
	app.Post("/books", handlers.CreateBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":3000")
}
