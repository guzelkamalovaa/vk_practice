package main

import (
	"net/http"
)

type HTTPServer struct {
	s *Store
}

func NewHTTPServer(store *Store) *HTTPServer {
	return &HTTPServer{s: store}
}

func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать маршрутизацию эндпоинтов
	w.WriteHeader(200)
	w.Write([]byte("Hello, segmentation API"))
}
