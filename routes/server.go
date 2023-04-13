package routes

import "github.com/go-chi/chi"

func CreateServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}
