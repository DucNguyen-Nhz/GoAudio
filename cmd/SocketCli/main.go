package main

import (
	"fmt"

	sck "github.com/DucNguyen-Nhz/GoStream/internal/socket"
)

func GetInput(client *sck.Client) {
	var msg string
	switch {
	case msg == "quit":
		client.Close()
		return
	default:
		fmt.Scanln(&msg)
		client.Send(msg)
	}
}

func main() {
	client := sck.CreateWSClient()
	client.Connect("ws://localhost:3000/ws", "http://localhost:3000")

	// client.Send("Hello, World!")
	// client.ReadLoop()
	for {
		go client.Read()
		go GetInput(client)
	}
}
