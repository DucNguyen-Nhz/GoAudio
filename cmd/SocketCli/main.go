package main

import (
	sck "github.com/DucNguyen-Nhz/GoStream/internal/socket"
)

func main() {
	client := sck.CreateWSClient()
	client.Connect("ws://localhost:3000/ws", "http://localhost:3000")

	client.Send("Hello, World!")
	client.Close()
}
