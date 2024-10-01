package main

import (
	"fmt"
	"net/http"

	"os"

	sck "github.com/DucNguyen-Nhz/GoStream/internal/socket"
	"golang.org/x/net/websocket"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := sck.CreateWS()
	fmt.Println("Server started on port " + port)
	http.Handle("/ws", websocket.Handler(server.HandleWS))
	http.ListenAndServe(":"+port, nil)
}
