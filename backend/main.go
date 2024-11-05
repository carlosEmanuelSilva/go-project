package main

import (
  "fmt"
  "net/http"
  "log"

  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
  for {
    messageType, p, err := conn.ReadMessage()

    if err != nil {
      log.Println(err)
      return
    }
  
    fmt.Println(p)

    if err := conn.WriteMessage(messageType, p); err != nil {
      log.Println(err)
      return
    }

  }
}

func serverWs(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.Host)

  ws, err := upgrader.Upgrade(w, r, nil)

  if err != nil {
    log.Println(err)
  }

  reader(ws)
}

func setupRoutes() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Setup do servidor")
  })

  http.HandleFunc("/ws", serverWs)
}

func main() {
  setupRoutes()
  http.ListenAndServe(":8080", nil)
}
