package controllers_test

import (
	"net/http"
	"pamukkale_university/routes"
	"pamukkale_university/test"
	"testing"
)

func TestIndex(t *testing.T) {
	test.TestMain()
	s := routes.InitTest()

	req, _ := http.NewRequest("GET", "/public/index", nil)

	response := test.ExecuteRequest(req, s)

	test.CheckResponseCode(t, http.StatusOK, response.Code)
}
