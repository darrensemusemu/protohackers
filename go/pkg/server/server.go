package server

import (
	"fmt"
	"net"
)

type Service interface {
	HandleConnection(conn net.Conn)
}

func New(svc Service) *server {
	return &server{
		svc: svc,
	}
}

type server struct {
	svc Service
}

func (s *server) Run(port int) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("server run listen err: %w", err)
	}
	defer ln.Close()
	fmt.Printf("Server listening on port: %d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("server run accept err: %w", err)
		}
		go s.svc.HandleConnection(conn)
	}
}
