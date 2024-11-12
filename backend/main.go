package main

import (
	"fmt"
  	"net/http"

	"github.com/x86-carlos/go-project/pkg/websocket"
)

func serverWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket endpoint hit")
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w,"%+V\n", err)
	}

	client := &websocket.Client {
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}

func main() {
  setupRoutes()
  http.ListenAndServe(":8080", nil)
}
