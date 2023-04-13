package departmentcontroller

import (
	"encoding/json"
	"net/http"
	"pamukkale_university/database"
	"pamukkale_university/ent/departments"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// that controller get all departments
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

// that controller return department what you want
func FetchById(w http.ResponseWriter, r *http.Request) {
	db := database.DBClient

	id := chi.URLParam(r, "id")
	parsed_id, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("id is invalid"))
		return
	}

	department, err := db.Departments.
		Query().
		Where(
			departments.ID(parsed_id),
		).
		First(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	} else {
		render.JSON(w, r, department)
		return
	}
}

// that controller create new department
func Create(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Code       string `json:"code" validate:"required"`
		Name       string `json:"name" validate:"required"`
		DegreeType bool   `json:"degree_type" validate:"required"`
		Quota      int    `json:"quota" validate:"required"`
	}
	var input Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	db := database.DBClient

	_, err := db.Departments.Create().
		SetName(input.Name).
		SetCode(input.Code).
		SetDegreeType(input.DegreeType).
		SetQuota(input.Quota).
		Save(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// that controller return department what you want
func DeleteById(w http.ResponseWriter, r *http.Request) {
	db := database.DBClient

	id := chi.URLParam(r, "id")
	parsed_id, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("id is invalid"))
		return
	}

	err = db.Departments.
		DeleteOneID(parsed_id).
		Exec(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
