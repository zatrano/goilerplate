package handlers

import (
	"strconv"

	"zatrano/internal/app/dtos"
	"zatrano/internal/app/models"
	"zatrano/internal/app/repositories"
	"zatrano/internal/app/services"
	"zatrano/internal/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func createUserDTOToModel(input dtos.CreateUserDTO) models.User {
	return models.User{
		Name: input.Name,
	}
}

func updateUserDTOToModel(input dtos.UpdateUserDTO, user models.User) models.User {
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	return user
}

func validateUserInputs(input interface{}) error {
	validate := validator.New()
	return validate.Struct(input)
}

func getUserService() *services.UserService {
	repo := repositories.NewUserRepository()
	return services.NewUserService(repo)
}

func GetAllUsers(c *fiber.Ctx) error {
	service := getUserService()

	users, err := service.GetAllUsers()
	if err != nil {
		return helpers.HandleError(c, err)
	}

	userDTOs := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dtos.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", userDTOs)
}

func GetUserByID(c *fiber.Ctx) error {
	service := getUserService()

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid user ID", nil))
	}

	user, err := service.GetUserByID(uint(userID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	userDTO := dtos.UserDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", userDTO)
}

func CreateUser(c *fiber.Ctx) error {
	service := getUserService()

	var input dtos.CreateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateUserInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	newUser := createUserDTOToModel(input)

	createdUser, err := service.CreateUser(newUser)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	response := dtos.UserDTO{
		ID:       createdUser.ID,
		Name:     createdUser.Name,
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", response)
}

func UpdateUser(c *fiber.Ctx) error {
	service := getUserService()

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid user ID", nil))
	}

	var input dtos.UpdateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateUserInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	userToUpdate, err := service.GetUserByID(uint(userID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	userToUpdate = updateUserDTOToModel(input, userToUpdate)

	updatedUser, err := service.UpdateUser(userToUpdate)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	response := dtos.UserDTO{
		ID:       updatedUser.ID,
		Name:     updatedUser.Name,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", response)
}

func DeleteUser(c *fiber.Ctx) error {
	service := getUserService()

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid user ID", nil))
	}

	err = service.DeleteUser(uint(userID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusNoContent, "Success", nil)
}
