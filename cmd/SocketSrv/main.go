package main

import (
	"fmt"
	"net/http"

	sck "github.com/DucNguyen-Nhz/GoStream/internal/socket"

	"golang.org/x/net/websocket"
)

var (
	port = fmt.Sprint(3000)
)

func main() {
	server := sck.CreateWS()
	fmt.Println("Server started on port " + port)
	http.Handle("/ws", websocket.Handler(server.HandleWS))
	http.ListenAndServe(":"+port, nil)
}
