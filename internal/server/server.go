package server

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"github.com/Raghav847/tcp_to_http/internal/request"
	"github.com/Raghav847/tcp_to_http/internal/response"
)

type Server struct {
	listener net.Listener
	closed   atomic.Bool
	handler  Handler
}

type HandleError struct {
	StatusCode response.StatusCode
	Message    string
}

// start with fixing this
type Handler func(w *response.Writer, req *request.Request) *HandleError

func Serve(port int) (*Server, error) {
	ln, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%d", port),
	)
	if err != nil {
		return nil, err
	}

	s := &Server{
		listener: ln,
	}

	go s.listen()

	return s, nil
}

func (s *Server) Close() error {
	s.closed.Store(true)
	return s.listener.Close()
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.closed.Load() {
				return
			}
			log.Printf("accept error: %v", err)
			continue
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	body := []byte("{Hello World!}\r\n")
	headers := response.GetDefaultHeaders(len(body))
	response.WriteStatusLine(conn, response.StatusOk)
	response.WriteHeaders(conn, headers)
	conn.Write(body)
}
