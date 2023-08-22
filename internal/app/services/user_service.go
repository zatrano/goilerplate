package services

import (
	"zatrano/internal/app/models"
	"zatrano/internal/app/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
