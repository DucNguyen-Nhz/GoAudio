package socket

import (
	"fmt"
	"io"
	"net"

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
		fmt.Println("Received from "+ws.Request().RemoteAddr+": ", string(msg))
		go s.Broadcast(msg)
	}
}

func (s *Server) Broadcast(b []byte) {
	for ws := range s.Conns {
		if _, err := ws.Write(b); err != nil {
			if err == io.EOF {
				delete(s.Conns, ws)
				continue
			}

			if _, ok := err.(*net.OpError); ok {
				delete(s.Conns, ws)
				fmt.Println("Connection closed by client:", ws.RemoteAddr())
				continue
			}

			fmt.Println("Write error: ", err)
		}

	}
}
