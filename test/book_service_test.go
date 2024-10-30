package services_test

import (
    "tp/internal/models"
    "tp/internal/services"
    "tp/internal/repositories/mocks" 
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateBook_Success(t *testing.T) {
    repo := new(mocks.BookRepository)
    service := services.NewBookService(repo)

    book := &models.Book{Title: "Go Programming", Author: "John Doe"}
    repo.On("Create", book).Return(nil)

    err := service.CreateBook(book)
    assert.NoError(t, err)
    repo.AssertExpectations(t)
}

func TestCreateBook_MissingFields(t *testing.T) {
    repo := new(mocks.BookRepository)
    service := services.NewBookService(repo)

    book := &models.Book{Title: "", Author: ""}
    err := service.CreateBook(book)

    assert.Error(t, err)
    assert.Equal(t, "Title and author are required", err.Error())
}

func TestGetBookByID_Success(t *testing.T) {
    repo := new(mocks.BookRepository)
    service := services.NewBookService(repo)

    book := &models.Book{ID: 1, Title: "Go Programming", Author: "John Doe"}
    repo.On("FindByID", uint(1)).Return(book, nil)

    result, err := service.GetBookByID(1)
    assert.NoError(t, err)
    assert.Equal(t, book, result)
    repo.AssertExpectations(t)
}

