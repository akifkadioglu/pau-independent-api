package routes

import (
	"net/http"
)

func (s *Server) Errors() {
	s.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
	})
	s.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
	})
}
