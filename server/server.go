package server

import (
	"io"
	"net/http"
)

type Server struct {
}

func (s *Server) Handler() http.Handler {
	top := http.NewServeMux()

	top.HandleFunc("/", s.handleHello)
	return top
}

func (s *Server) handleHello(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, "hello world\n")
}

func New() *Server {
	return &Server{}
}
