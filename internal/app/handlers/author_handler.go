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

func createAuthorDTOToModel(input dtos.CreateAuthorDTO) models.Author {
	return models.Author{
		Name: input.Name,
	}
}

func updateAuthorDTOToModel(input dtos.UpdateAuthorDTO, author models.Author) models.Author {
	author.Name = input.Name
	return author
}

func validateAuthorInputs(input interface{}) error {
	validate := validator.New()
	return validate.Struct(input)
}

func getAuthorService() *services.AuthorService {
	repo := repositories.NewAuthorRepository()
	return services.NewAuthorService(repo)
}

func GetAllAuthors(c *fiber.Ctx) error {
	service := getAuthorService()

	authors, err := service.GetAllAuthors()
	if err != nil {
		return helpers.HandleError(c, err)
	}

	authorDTOs := make([]dtos.AuthorDTO, len(authors))
	for i, author := range authors {
		authorDTOs[i] = dtos.AuthorDTO{
			ID:   author.ID,
			Name: author.Name,
		}
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", authorDTOs)
}

func GetAuthorByID(c *fiber.Ctx) error {
	service := getAuthorService()

	authorIDParam := c.Params("id")
	authorID, err := strconv.Atoi(authorIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid author ID", nil))
	}

	author, err := service.GetAuthorByID(uint(authorID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	authorDTO := dtos.AuthorDTO{
		ID:   author.ID,
		Name: author.Name,
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", authorDTO)
}

func CreateAuthor(c *fiber.Ctx) error {
	service := getAuthorService()

	var input dtos.CreateAuthorDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateAuthorInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	newAuthor := createAuthorDTOToModel(input)

	createdAuthor, err := service.CreateAuthor(newAuthor)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	response := dtos.AuthorDTO{
		ID:   createdAuthor.ID,
		Name: createdAuthor.Name,
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", response)
}

func UpdateAuthor(c *fiber.Ctx) error {
	service := getAuthorService()

	authorIDParam := c.Params("id")
	authorID, err := strconv.Atoi(authorIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid author ID", nil))
	}

	var input dtos.UpdateAuthorDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateAuthorInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	authorToUpdate, err := service.GetAuthorByID(uint(authorID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	authorToUpdate = updateAuthorDTOToModel(input, authorToUpdate)

	updatedAuthor, err := service.UpdateAuthor(authorToUpdate)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	response := dtos.AuthorDTO{
		ID:   updatedAuthor.ID,
		Name: updatedAuthor.Name,
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", response)
}

func DeleteAuthor(c *fiber.Ctx) error {
	service := getAuthorService()

	authorIDParam := c.Params("id")
	authorID, err := strconv.Atoi(authorIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid author ID", nil))
	}

	err = service.DeleteAuthor(uint(authorID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusNoContent, "Success", nil)
}
