package repositories

import (
  "tp/internal/models"
  "tp/config"
)
type BookRepository interface {
    Create(book *models.Book) error
    FindAll() ([]models.Book, error)
    FindByID(id uint) (*models.Book, error)
    Update(book *models.Book) error
    Delete(id uint) error
}


type bookRepository struct{}

func NewBookRepository() BookRepository {
    return &bookRepository{}
}

func (r *bookRepository) Create(book *models.Book) error {
    return config.DB.Create(book).Error
}

func (r *bookRepository) FindAll() ([]models.Book, error) {
    var books []models.Book
    err := config.DB.Find(&books).Error
    return books, err
}

func (r *bookRepository) FindByID(id uint) (*models.Book, error) {
    var book models.Book
    err := config.DB.Where("id = ?", id).First(&book).Error
    return &book, err
}

func (r *bookRepository) Update(book *models.Book) error {
    return config.DB.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
    return config.DB.Delete(&models.Book{}, id).Error
}
