package main

import (
  "tp/config"
  "tp/internal/controllers"
  "tp/internal/repositories"
  "github.com/gin-gonic/gin"
)

func main() {
  //Conexão com o banco de dados
  config.ConnectDatabase()

  r := gin.Default()
  
  //Instanciando um repositorio e um controller
  bookRepo := repositories.NewBookRepository()
  bookController := controllers.NewBookController(bookRepo)
  
  //Definir rotas
  r.POST("/books", bookController.CreateBook)
  r.GET("/books", bookController.GetAllBooks)
  r.GET("/books/:id", bookController.GetBookByID)
  r.PUT("/books/:id", bookController.UpdateBook)
  r.DELETE("/books/:id", bookController.DeleteBook)

  r.Run(":8080")
}
