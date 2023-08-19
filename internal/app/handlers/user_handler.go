package handlers

import (
	"strconv"

	"zatrano/internal/app/dtos"
	"zatrano/internal/app/models"
	"zatrano/internal/app/repositories"
	"zatrano/internal/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)

	users, err := service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	userDTOs := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dtos.UserDTO{
			ID:   user.ID,
			Name: user.Name,
		}
	}

	return c.JSON(userDTOs)
}

func GetUserByID(c *fiber.Ctx) error {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := service.GetUserByID(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	userDTO := dtos.UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}
	return c.JSON(userDTO)
}

func CreateUser(c *fiber.Ctx) error {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)

	var input dtos.CreateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	createdUser, err := service.CreateUser(models.User{
		Name: input.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	response := dtos.UserDTO{
		ID:   createdUser.ID,
		Name: createdUser.Name,
	}
	return c.JSON(response)
}

func UpdateUser(c *fiber.Ctx) error {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var input dtos.UpdateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	userToUpdate, err := service.GetUserByID(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	userToUpdate.Name = input.Name

	updatedUser, err := service.UpdateUser(userToUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	response := dtos.UserDTO{
		ID:   updatedUser.ID,
		Name: updatedUser.Name,
	}
	return c.JSON(response)
}

func DeleteUser(c *fiber.Ctx) error {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = service.DeleteUser(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
