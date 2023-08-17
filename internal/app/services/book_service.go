package services

import (
	"github.com/zatrano/zatrano/internal/app/models"
	"github.com/zatrano/zatrano/internal/app/repositories"
)

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id uint) (models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) CreateBook(book models.Book) (models.Book, error) {
	return s.repo.Create(book)
}

func (s *BookService) UpdateBook(book models.Book) (models.Book, error) {
	return s.repo.Update(book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
