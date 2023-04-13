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

	// Create a New Request
	req, _ := http.NewRequest("GET", routes.IndexRoute, nil)

	// Execute Request
	response := test.ExecuteRequest(req, s)

	// Check the response code
	test.CheckResponseCode(t, http.StatusNoContent, response.Code)
}
