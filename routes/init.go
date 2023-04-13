package routes

import (
	"log"
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

	err:=http.ListenAndServe(env.Getenv(env.PORT), s.Router)
	if err != nil {
		log.Println(err.Error())
	}
}

func InitTest() *Server {
	s := CreateServer()
	s.Public()
	s.Admin()
	s.Errors()
	return s
}