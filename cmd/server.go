package main

import (
  "tp/config"
  "tp/internal/controllers"
  "github.com/gin-gonic/gin"
)

func main() {
  //Conexão com o banco de dados
  config.ConnectDatabase()

  r := gin.Default()
  
  //Definir rotas
  r.POST("/user", controllers.CreateUser)
  r.POST("/users", controllers.CreateUser)
  r.POST("/books", controllers.CreateBook)
  r.GET("/books", controllers.GetAllBooks)
  r.GET("/books/:id", controllers.GetBookByID)
  r.PUT("/books/:id", controllers.UpdateBook)
  r.DELETE("/books/:id", controllers.DeleteBook)
  r.POST("/reviews", controllers.CreateReview)

  r.Run(":8080")
}
