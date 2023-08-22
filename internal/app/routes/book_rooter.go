package routes

import (
	"zatrano/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type BookRouter struct{}

func (bookRouter *BookRouter) SetupRoutes(router fiber.Router) {
	books := router.Group("/books")
	books.Get("/", handlers.GetAllBooks)
	books.Get("/:id", handlers.GetBookByID)
	books.Post("/", handlers.CreateBook)
	books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}
