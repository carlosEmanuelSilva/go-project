package services

import (
  "fmt"
  "tp/internal/models"
  "tp/internal/repositories"
)

type BookService interface {
    CreateBook(book *models.Book) error
    GetAllBooks() ([]models.Book, error)
    GetBookByID(id uint) (*models.Book, error)
    UpdateBook(book *models.Book) error
    DeleteBook(id uint) error
}

type bookService struct {
    repository repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
    return &bookService{repository: repo}
}

func (s *bookService) CreateBook(book *models.Book) error {
    if book.Title == "" || book.Author == "" {
        return fmt.Errorf("Title and author are required")
    }
    return s.repository.Create(book)
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
    return s.repository.FindAll()
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
    return s.repository.FindByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
    if book.Title == "" || book.Author == "" {
        return fmt.Errorf("Title and author are required")
    }
    return s.repository.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
  return s.repository.Delete(id)
}
