package handlers

import (
	"strconv"

	"zatrano/internal/app/dtos"
	"zatrano/internal/app/models"
	"zatrano/internal/app/repositories"
	"zatrano/internal/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func createUserDTOToModel(input dtos.CreateUserDTO) models.User {
	return models.User{
		Name: input.Name,
	}
}

func validateUserInputs(input interface{}) error {
	validate := validator.New()
	return validate.Struct(input)
}

func handleUserError(c *fiber.Ctx, err error) error {
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
}

func GetAllUsers(c *fiber.Ctx) error {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)

	users, err := service.GetAllUsers()
	if err != nil {
		return handleUserError(c, err)
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
		return handleUserError(c, err)
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

	if err := validateUserInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	newUser := createUserDTOToModel(input)

	createdUser, err := service.CreateUser(newUser)
	if err != nil {
		return handleUserError(c, err)
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

	if err := validateUserInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	userToUpdate, err := service.GetUserByID(uint(userID))
	if err != nil {
		return handleUserError(c, err)
	}

	userToUpdate.Name = input.Name

	updatedUser, err := service.UpdateUser(userToUpdate)
	if err != nil {
		return handleUserError(c, err)
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
		return handleUserError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
