package controllers_test

import (
    "bytes"
    "strconv"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "tp/config"
    "tp/internal/controllers"
    "tp/internal/models"
    "tp/internal/repositories"
    "tp/internal/services"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

// Setup para criar um servidor e controlador para os testes
func setupRouter() *gin.Engine {
    config.ConnectDatabase() // Banco de dados de teste
    bookRepo := repositories.NewBookRepository()
    bookService := services.NewBookService(bookRepo)
    bookController := controllers.NewBookController(bookService)

    router := gin.Default()
    router.POST("/books", bookController.CreateBook)
    router.GET("/books/:id", bookController.GetBookByID)
    router.GET("/books", bookController.GetAllBooks)

    return router
}

// Teste de integração para criar um novo livro
func TestCreateBookEndpoint(t *testing.T) {
    router := setupRouter()

    book := models.Book{Title: "Go Programming", Author: "John Doe"}
    jsonValue, _ := json.Marshal(book)

    req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
}

// Teste de integração para buscar um livro pelo ID
func TestGetBookByIDEndpoint(t *testing.T) {
    router := setupRouter()

    // Primeiro, criamos um livro para testar
    book := models.Book{Title: "Go Programming", Author: "John Doe"}
    config.DB.Create(&book)

    req, _ := http.NewRequest("GET", "/books/"+strconv.Itoa(int(book.ID)), nil)
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
}

// Teste de integração para buscar todos os livros
func TestGetAllBooksEndpoint(t *testing.T) {
    router := setupRouter()

    req, _ := http.NewRequest("GET", "/books", nil)
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
}
