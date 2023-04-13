package routes

import (
	departmentcontroller "pamukkale_university/controller/department_controller"
	indexcontroller "pamukkale_university/controller/index_controller"

	"github.com/go-chi/chi"
)

func (s *Server) Public() {
	s.Router.Group(func(r chi.Router) {
		PublicMiddleware(r)
		//for all people
		r.Route("/public", func(p chi.Router) {
			p.Get("/index", indexcontroller.Index)

			//departments
			p.Route("/department", func(d chi.Router) {
				d.Get("/all", departmentcontroller.All)
				d.Get("/fetch-by-id/{id}", departmentcontroller.FetchById)
			})
		})
	})
}

func (s *Server) Admin() {
	s.Router.Group(func(r chi.Router) {
		AdminMiddleware(r)
		//just debug mode
		r.Route("/admin", func(a chi.Router) {

			//departments
			a.Route("/department", func(d chi.Router) {
				d.Post("/create", departmentcontroller.Create)
				d.Delete("/delete-by-id/{id}", departmentcontroller.DeleteById)
			})
		})
	})
}
