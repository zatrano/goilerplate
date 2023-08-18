package services

import (
	"github.com/zatrano/zatrano/internal/app/models"
	"github.com/zatrano/zatrano/internal/app/repositories"
)

type AuthorService struct {
	repo *repositories.AuthorRepository
}

func NewAuthorService(repo *repositories.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
	return s.repo.GetAll()
}

func (s *AuthorService) GetAuthorByID(id uint) (models.Author, error) {
	return s.repo.GetByID(id)
}

func (s *AuthorService) CreateAuthor(book models.Author) (models.Author, error) {
	return s.repo.Create(book)
}

func (s *AuthorService) UpdateAuthor(book models.Author) (models.Author, error) {
	return s.repo.Update(book)
}

func (s *AuthorService) DeleteAuthor(id uint) error {
	return s.repo.Delete(id)
}
