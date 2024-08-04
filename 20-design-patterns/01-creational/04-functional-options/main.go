package main

import (
	"fmt"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(options ...func(*Server)) *Server {
	// default server
	svr := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
		maxConn: 1000,
	}

	// apply options
	for _, opt := range options {
		opt(svr)
	}
	return svr
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func main() {
	srv1 := New(
		WithHost("127.0.0.1"),
		WithPort(8081),
	)

	fmt.Println(srv1)
}
