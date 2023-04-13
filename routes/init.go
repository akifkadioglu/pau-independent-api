package routes

import (
	"log"
	"net/http"

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

	err:=http.ListenAndServe(":10000", s.Router)
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