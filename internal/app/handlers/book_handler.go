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
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
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

	author := models.Author{
		Name: input.Author.Name,
	}

	newBook := models.Book{
		Title:  input.Title,
		Author: author,
	}

	createdBook, err := service.CreateBook(newBook)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
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

	bookToUpdate, err := service.GetBookByID(uint(bookID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	bookToUpdate.Title = input.Title

	author := models.Author{
		Name: input.Author.Name,
	}
	bookToUpdate.Author = author

	updatedBook, err := service.UpdateBook(bookToUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	response := dtos.BookDTO{
		ID:    updatedBook.ID,
		Title: updatedBook.Title,
		Author: dtos.AuthorDTO{
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
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
