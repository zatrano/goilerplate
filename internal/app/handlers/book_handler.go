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

func createBookDTOToModel(input dtos.CreateBookDTO) models.Book {
	return models.Book{
		Title: input.Title,
		Author: models.Author{
			Name: input.Author.Name,
		},
	}
}

func updateBookDTOToModel(input dtos.UpdateBookDTO, book models.Book) models.Book {
	book.Title = input.Title
	book.Author.Name = input.Author.Name
	return book
}

func validateBookInputs(input interface{}) error {
	validate := validator.New()
	return validate.Struct(input)
}

func getBookService() *services.BookService {
	repo := repositories.NewBookRepository()
	return services.NewBookService(repo)
}

func GetAllBooks(c *fiber.Ctx) error {
	service := getBookService()

	books, err := service.GetAllBooks()
	if err != nil {
		return helpers.HandleError(c, err)
	}

	bookDTOs := make([]dtos.BookDTO, len(books))
	for i, book := range books {
		bookDTOs[i] = dtos.BookDTO{
			ID:    book.ID,
			Title: book.Title,
			Author: dtos.AuthorDTO{
				ID:   book.Author.ID,
				Name: book.Author.Name,
			},
		}
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", bookDTOs)
}

func GetBookByID(c *fiber.Ctx) error {
	service := getBookService()

	bookIDParam := c.Params("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid book ID", nil))
	}

	book, err := service.GetBookByID(uint(bookID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	bookDTO := dtos.BookDTO{
		ID:    book.ID,
		Title: book.Title,
		Author: dtos.AuthorDTO{
			ID:   book.Author.ID,
			Name: book.Author.Name,
		},
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", bookDTO)
}

func CreateBook(c *fiber.Ctx) error {
	service := getBookService()

	var input dtos.CreateBookDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateBookInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	newBook := createBookDTOToModel(input)

	createdBook, err := service.CreateBook(newBook)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	response := dtos.BookDTO{
		ID:    createdBook.ID,
		Title: createdBook.Title,
		Author: dtos.AuthorDTO{
			ID:   createdBook.Author.ID,
			Name: createdBook.Author.Name,
		},
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", response)
}

func UpdateBook(c *fiber.Ctx) error {
	service := getBookService()

	bookIDParam := c.Params("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid book ID", nil))
	}

	var input dtos.UpdateBookDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateBookInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	bookToUpdate, err := service.GetBookByID(uint(bookID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	bookToUpdate = updateBookDTOToModel(input, bookToUpdate)

	updatedBook, err := service.UpdateBook(bookToUpdate)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	response := dtos.BookDTO{
		ID:    updatedBook.ID,
		Title: updatedBook.Title,
		Author: dtos.AuthorDTO{
			ID:   updatedBook.Author.ID,
			Name: updatedBook.Author.Name,
		},
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", response)
}

func DeleteBook(c *fiber.Ctx) error {
	service := getBookService()

	bookIDParam := c.Params("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid book ID", nil))
	}

	err = service.DeleteBook(uint(bookID))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusNoContent, "Success", nil)
}
