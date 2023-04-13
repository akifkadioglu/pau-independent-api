package test

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"pamukkale_university/database"
	"pamukkale_university/ent"
	"pamukkale_university/env"
	"pamukkale_university/routes"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var err error

func TestMain() {
	env.InitTest()
	PrepareDatabaseForTest()
}

func PrepareDatabaseForTest() {
	database.DBClient, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := database.DBClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func ExecuteRequest(req *http.Request, s *routes.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func MapToBody(mock map[string]any) *io.PipeReader {
	pr, pw := io.Pipe()
	go func() {
		json.NewEncoder(pw).Encode(mock)
		pw.Close()
	}()
	return pr
}
