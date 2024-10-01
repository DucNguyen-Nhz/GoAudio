package socket

import (
	"fmt"
	"os"

	"golang.org/x/net/websocket"
)

type ClientAPI interface {
	CreateWSClient() *Client
}

type Client struct {
	Conn *websocket.Conn
}

func CreateWSClient() *Client {
	return &Client{
		Conn: nil,
	}
}

func (c *Client) Connect(socket string, origin string) {
	ws, err := websocket.Dial(socket, "", origin)
	fmt.Println("Connecting to:", socket)
	if err != nil {
		fmt.Println("Error dialing:", err)
		os.Exit(1)
	}
	c.Conn = ws
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) ReadLoop() {
	go c.Read()
}

func (c *Client) Read() {
	buf := make([]byte, 1024)

	n, err := c.Conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	msg := string(buf[:n])
	fmt.Println("Received:", msg)
}

func (c *Client) Send(msg string) {
	if _, err := c.Conn.Write([]byte(msg)); err != nil {
		fmt.Println("Error sending:", err)
	}
	fmt.Println("Sent:", msg)
}
