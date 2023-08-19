package routes

import (
	"zatrano/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

func (userRouter *UserRouter) SetupRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", handlers.GetAllUsers)
	users.Get("/:id", handlers.GetUserByID)
	users.Post("/", handlers.CreateUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)
}
