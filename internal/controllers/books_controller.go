package controllers

import (
  "net/http"
  "tp/config"
  "tp/internal/models"
  "github.com/gin-gonic/gin"
)

// Cria um novo livro
func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

// Obtém todos os livros
func GetAllBooks(c *gin.Context) {
    var books []models.Book
    if err := config.DB.Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

// Obtém um livro pelo ID
func GetBookByID(c *gin.Context) {
    var book models.Book
    if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

// Atualiza um livro pelo ID
func UpdateBook(c *gin.Context) {
    var book models.Book
    if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

// Exclui um livro pelo ID
func DeleteBook(c *gin.Context) {
    if err := config.DB.Delete(&models.Book{}, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
