package departmentcontroller

import (
	"github.com/go-chi/render"
	"net/http"
	"pamukkale_university/database"
)

func All(w http.ResponseWriter, r *http.Request) {
	db := database.DBClient
	departments, err := db.Departments.Query().All(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, departments)
}