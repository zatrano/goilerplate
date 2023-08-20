package handlers

import (
	"strconv"

	"zatrano/internal/app/dtos"
	"zatrano/internal/app/models"
	"zatrano/internal/app/repositories"
	"zatrano/internal/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func createAuthorDTOToModel(input dtos.CreateAuthorDTO) models.Author {
	return models.Author{
		Name: input.Name,
	}
}

func validateInputs(input interface{}) error {
	validate := validator.New()
	return validate.Struct(input)
}

func handleError(c *fiber.Ctx, err error) error {
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
}

func GetAllAuthors(c *fiber.Ctx) error {
	repo := repositories.NewAuthorRepository()
	service := services.NewAuthorService(repo)

	authors, err := service.GetAllAuthors()
	if err != nil {
		return handleError(c, err)
	}

	authorDTOs := make([]dtos.AuthorDTO, len(authors))
	for i, author := range authors {
		authorDTOs[i] = dtos.AuthorDTO{
			ID:   author.ID,
			Name: author.Name,
		}
	}

	return c.JSON(authorDTOs)
}

func GetAuthorByID(c *fiber.Ctx) error {
	repo := repositories.NewAuthorRepository()
	service := services.NewAuthorService(repo)

	authorIDParam := c.Params("id")
	authorID, err := strconv.Atoi(authorIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid author ID"})
	}

	author, err := service.GetAuthorByID(uint(authorID))
	if err != nil {
		return handleError(c, err)
	}

	authorDTO := dtos.AuthorDTO{
		ID:   author.ID,
		Name: author.Name,
	}
	return c.JSON(authorDTO)
}

func CreateAuthor(c *fiber.Ctx) error {
	repo := repositories.NewAuthorRepository()
	service := services.NewAuthorService(repo)

	var input dtos.CreateAuthorDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := validateInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	newAuthor := createAuthorDTOToModel(input)

	createdAuthor, err := service.CreateAuthor(newAuthor)
	if err != nil {
		return handleError(c, err)
	}

	response := dtos.AuthorDTO{
		ID:   createdAuthor.ID,
		Name: createdAuthor.Name,
	}
	return c.JSON(response)
}

func UpdateAuthor(c *fiber.Ctx) error {
	repo := repositories.NewAuthorRepository()
	service := services.NewAuthorService(repo)

	authorIDParam := c.Params("id")
	authorID, err := strconv.Atoi(authorIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid author ID"})
	}

	var input dtos.UpdateAuthorDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := validateInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	authorToUpdate, err := service.GetAuthorByID(uint(authorID))
	if err != nil {
		return handleError(c, err)
	}

	authorToUpdate.Name = input.Name

	updatedAuthor, err := service.UpdateAuthor(authorToUpdate)
	if err != nil {
		return handleError(c, err)
	}

	response := dtos.AuthorDTO{
		ID:   updatedAuthor.ID,
		Name: updatedAuthor.Name,
	}
	return c.JSON(response)
}

func DeleteAuthor(c *fiber.Ctx) error {
	repo := repositories.NewAuthorRepository()
	service := services.NewAuthorService(repo)

	authorIDParam := c.Params("id")
	authorID, err := strconv.Atoi(authorIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid author ID"})
	}

	err = service.DeleteAuthor(uint(authorID))
	if err != nil {
		return handleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
