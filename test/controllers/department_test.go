package controllers_test

import (
	"net/http"
	"pamukkale_university/routes"
	"pamukkale_university/test"
	"testing"
)

func TestDepartments(t *testing.T) {
	test.TestMain()
	s := routes.InitTest()
	t.Run("create", func(t *testing.T) {

		//prepare body
		var body = test.MapToBody(map[string]any{
			"code":        "6546451",
			"name":        "deneme",
			"degree_type": false,
			"quota":       20,
		})

		req, _ := http.NewRequest("POST", "/admin/department/create", body)

		response := test.ExecuteRequest(req, s)

		test.CheckResponseCode(t, http.StatusCreated, response.Code)
	})

	t.Run("all", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/public/department/all", nil)

		response := test.ExecuteRequest(req, s)

		test.CheckResponseCode(t, http.StatusOK, response.Code)
	})

}
