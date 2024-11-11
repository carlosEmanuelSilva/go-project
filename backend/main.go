package main

import (
	"fmt"
  	"net/http"

	"github.com/x86-carlos/go-project/pkg/websocket"
)

func serverWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w,"%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serverWs)
}

func main() {
  setupRoutes()
  http.ListenAndServe(":8080", nil)
}
