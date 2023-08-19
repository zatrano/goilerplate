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

func GetAllAuthors(c *fiber.Ctx) error {
	repo := repositories.NewAuthorRepository()
	service := services.NewAuthorService(repo)

	authors, err := service.GetAllAuthors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
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
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
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

	createdAuthor, err := service.CreateAuthor(models.Author{
		Name: input.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
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

	authorToUpdate, err := service.GetAuthorByID(uint(authorID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	authorToUpdate.Name = input.Name

	updatedAuthor, err := service.UpdateAuthor(authorToUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
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
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
