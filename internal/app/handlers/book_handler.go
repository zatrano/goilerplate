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

func handleBookError(c *fiber.Ctx, err error) error {
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
}

func GetAllBooks(c *fiber.Ctx) error {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)

	books, err := service.GetAllBooks()
	if err != nil {
		return handleBookError(c, err)
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

	return c.JSON(bookDTOs)
}

func GetBookByID(c *fiber.Ctx) error {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)

	bookIDParam := c.Params("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	book, err := service.GetBookByID(uint(bookID))
	if err != nil {
		return handleBookError(c, err)
	}

	bookDTO := dtos.BookDTO{
		ID:    book.ID,
		Title: book.Title,
		Author: dtos.AuthorDTO{
			ID:   book.Author.ID,
			Name: book.Author.Name,
		},
	}
	return c.JSON(bookDTO)
}

func CreateBook(c *fiber.Ctx) error {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)

	var input dtos.CreateBookDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := validateBookInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	newBook := createBookDTOToModel(input)

	createdBook, err := service.CreateBook(newBook)
	if err != nil {
		return handleBookError(c, err)
	}

	response := dtos.BookDTO{
		ID:    createdBook.ID,
		Title: createdBook.Title,
		Author: dtos.AuthorDTO{
			ID:   createdBook.Author.ID,
			Name: createdBook.Author.Name,
		},
	}
	return c.JSON(response)
}

func UpdateBook(c *fiber.Ctx) error {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)

	bookIDParam := c.Params("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	var input dtos.UpdateBookDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := validateBookInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	bookToUpdate, err := service.GetBookByID(uint(bookID))
	if err != nil {
		return handleBookError(c, err)
	}

	bookToUpdate = updateBookDTOToModel(input, bookToUpdate)

	updatedBook, err := service.UpdateBook(bookToUpdate)
	if err != nil {
		return handleBookError(c, err)
	}

	response := dtos.BookDTO{
		ID:    updatedBook.ID,
		Title: updatedBook.Title,
		Author: dtos.AuthorDTO{
			ID:   updatedBook.Author.ID,
			Name: updatedBook.Author.Name,
		},
	}
	return c.JSON(response)
}

func DeleteBook(c *fiber.Ctx) error {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)

	bookIDParam := c.Params("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	err = service.DeleteBook(uint(bookID))
	if err != nil {
		return handleBookError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
