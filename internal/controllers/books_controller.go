package controllers

import (
  "net/http"
  "tp/internal/models"
  "tp/internal/services"
  "github.com/gin-gonic/gin"
  "strconv"
)

type BookController struct {
  service services.BookService
}

func NewBookController(service services.BookService) *BookController {
  return &BookController{service: service}
}

// Cria um novo livro
func (bc *BookController) CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := bc.service.CreateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

// Obtém todos os livros
func (bc *BookController) GetAllBooks(c *gin.Context) {
   books, err := bc.service.GetAllBooks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

// Obtém um livro pelo ID
func (bc *BookController) GetBookByID(c *gin.Context) {
   idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }
    book, err := bc.service.GetBookByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

// Atualiza um livro pelo ID
func (bc *BookController) UpdateBook(c *gin.Context) {
      // Extrai o ID do livro da URL
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    // Busca o livro existente no serviço
    book, err := bc.service.GetBookByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    // Faz o bind do JSON recebido para a estrutura do livro existente
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Chama o serviço para atualizar o livro
    if err := bc.service.UpdateBook(book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Retorna o livro atualizado
    c.JSON(http.StatusOK, book)
}

// Exclui um livro pelo ID
func (bc *BookController) DeleteBook(c *gin.Context) {
  idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }
    if err := bc.service.DeleteBook(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
