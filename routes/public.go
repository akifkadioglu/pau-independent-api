package routes

import (
	departmentcontroller "pamukkale_university/controller/department_controller"
	indexcontroller "pamukkale_university/controller/index_controller"

	"github.com/go-chi/chi/middleware"
)

const (
	IndexRoute       string = "/"
	DepartmentsRoute string = "/departments"
)

func (s *Server) Public() {
	s.Router.Use(middleware.Logger)
	s.Router.Get(IndexRoute, indexcontroller.Index)
	s.Router.Get(DepartmentsRoute, departmentcontroller.All)
}