package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/zatrano/internal/app/dtos"
	"github.com/zatrano/zatrano/internal/app/models"
	"github.com/zatrano/zatrano/internal/app/repositories"
	"github.com/zatrano/zatrano/internal/app/services"
	"gorm.io/gorm"
)

func GetAllBooks(c *fiber.Ctx) error {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)

	books, err := service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	bookDTOs := make([]dtos.BookDTO, len(books))
	for i, book := range books {
		bookDTOs[i] = dtos.BookDTO{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
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
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	bookDTO := dtos.BookDTO{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
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

	createdBook, err := service.CreateBook(models.Book{
		Title:  input.Title,
		Author: input.Author,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	response := dtos.BookDTO{
		ID:     createdBook.ID,
		Title:  createdBook.Title,
		Author: createdBook.Author,
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

	bookToUpdate, err := service.GetBookByID(uint(bookID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	bookToUpdate.Title = input.Title
	bookToUpdate.Author = input.Author

	updatedBook, err := service.UpdateBook(bookToUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	response := dtos.BookDTO{
		ID:     updatedBook.ID,
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
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
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
