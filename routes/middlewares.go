package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func PublicMiddleware(r chi.Router) {
	r.Use(middleware.Logger)
}

func AdminMiddleware(r chi.Router) {
	r.Use(middleware.Logger)
}
