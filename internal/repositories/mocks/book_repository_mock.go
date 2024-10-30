package mocks

import (
    "tp/internal/models"

    "github.com/stretchr/testify/mock"
)

// BookRepository é o mock para o repositório de livros
type BookRepository struct {
    mock.Mock
}

// Mock do método Create
func (m *BookRepository) Create(book *models.Book) error {
    args := m.Called(book)
    return args.Error(0)
}

// Mock do método FindByID
func (m *BookRepository) FindByID(id uint) (*models.Book, error) {
    args := m.Called(id)
    if book, ok := args.Get(0).(*models.Book); ok {
        return book, args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *BookRepository) FindAll() ([]models.Book, error) {
    args := m.Called()
    if books, ok := args.Get(0).([]models.Book); ok {
        return books, args.Error(1)
    }
    return nil, args.Error(1)
}

// Mock do método Update
func (m *BookRepository) Update(book *models.Book) error {
    args := m.Called(book)
    return args.Error(0)
}

// Mock do método Delete
func (m *BookRepository) Delete(id uint) error {
    args := m.Called(id)
    return args.Error(0)
}
