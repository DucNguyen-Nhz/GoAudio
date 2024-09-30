package socket

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type ServerAPI interface {
	CreateWS() *Server
	HandleWS(ws *websocket.Conn)
	ReadLoop(ws *websocket.Conn)
}

type Server struct {
	Conns map[*websocket.Conn]bool
}

func CreateWS() *Server {
	return &Server{
		Conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())
	s.Conns[ws] = true
	s.ReadLoop(ws)
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading:", err)
			continue
		}
		msg := buf[:n]
		s.Broadcast(msg)
	}
}

func (s *Server) Broadcast(b []byte) {
	for ws := range s.Conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error: ", err)
			}
		}(ws)
	}
}
