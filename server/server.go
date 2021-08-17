package server

import (
	"io"
	"net/http"

	"github.com/dexterlb/dumbloader/static"

	log "github.com/sirupsen/logrus"
)

type Server struct {
}

func (s *Server) Handler() http.Handler {
	top := http.NewServeMux()

	// while we don't have multiple routes here,
	// using ServeMux is still beneficial since it sanitises URLs and
	// does some other cleanup work
	top.HandleFunc("/", s.handleAny)

	return top
}

func (s *Server) handleAny(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" && req.Method == "GET" {
		s.handleUI(rw, req)
		return
	}

	if req.Method == "POST" {
		s.handleUpload(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	io.WriteString(rw, "404\n")
}

func (s *Server) handleUI(rw http.ResponseWriter, req *http.Request) {
	static.Serve(static.Index, rw, req)
}

func (s *Server) handleUpload(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	log.WithFields(log.Fields{
		"path": req.URL.Path,
	}).Info("Initiated upload")
	io.WriteString(rw, "hello upload\n")
}

func New() *Server {
	return &Server{}
}
