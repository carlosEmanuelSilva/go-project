package main

import (
  "tp/config"
  "tp/internal/controllers"
  "tp/internal/repositories"
  "tp/internal/services"
  "github.com/gin-gonic/gin"
)

func main() {
  //Conexão com o banco de dados
  config.ConnectDatabase()

  r := gin.Default()
  
  bookRepo := repositories.NewBookRepository()
  bookService := services.NewBookService(bookRepo)
  bookController := controllers.NewBookController(bookService)
  
  //Definir rotas
  r.POST("/books", bookController.CreateBook)
  r.GET("/books", bookController.GetAllBooks)
  r.GET("/books/:id", bookController.GetBookByID)
  r.PUT("/books/:id", bookController.UpdateBook)
  r.DELETE("/books/:id", bookController.DeleteBook)

  r.Run(":8080")
}
