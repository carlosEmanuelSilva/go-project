package controllers

import (
    "net/http"
    "tp/config"
    "tp/internal/models"
    "github.com/gin-gonic/gin"
)

// Cria uma nova avaliação
func CreateReview(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&review).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, review)
}

