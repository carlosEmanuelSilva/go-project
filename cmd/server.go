package main

import (
  "fmt"
  "tp/config"
)

func main() {
  config.ConnectDatabase()
  fmt.Println("Conexão OK!!")
}
