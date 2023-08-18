package repositories

import (
	"github.com/zatrano/zatrano/internal/app/models"
	"github.com/zatrano/zatrano/internal/database"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository() *BookRepository {
	return &BookRepository{db: database.SetupDatabase()}
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *BookRepository) GetByID(id uint) (models.Book, error) {
	var book models.Book
	result := r.db.First(&book, id)
	return book, result.Error
}

func (r *BookRepository) Create(book models.Book) (models.Book, error) {
	result := r.db.Create(&book)
	return book, result.Error
}

func (r *BookRepository) Update(book models.Book) (models.Book, error) {
	result := r.db.Save(&book)
	return book, result.Error
}

func (r *BookRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Book{}, id)
	return result.Error
}
