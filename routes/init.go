package routes

import (
	"net/http"
	"pamukkale_university/env"

	"github.com/go-chi/chi"
)

type Server struct {
	Router *chi.Mux
}

func InitProd() {
	s := CreateServer()
	s.Public()
	//s.Admin()
	s.Errors()

	http.ListenAndServe(env.Getenv(env.PORT), s.Router)
}

func InitTest() *Server {
	s := CreateServer()
	s.Public()
	s.Admin()
	s.Errors()
	return s
}