package controllers

import (
    "net/http"
    "strconv"
    "tp/internal/models"
    "tp/internal/services"

    "github.com/gin-gonic/gin"
)

type ReviewController struct {
    ReviewService services.ReviewService
}

// Criar nova avaliação (Review)
func (rc *ReviewController) CreateReview(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdReview, err := rc.ReviewService.CreateReview(review)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdReview)
}

// Listar avaliações por ID do livro
func (rc *ReviewController) GetReviewsByBookID(c *gin.Context) {
    bookID, err := strconv.Atoi(c.Param("book_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    reviews, err := rc.ReviewService.GetReviewsByBookID(bookID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reviews)
}

