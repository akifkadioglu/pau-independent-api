package database

import (
	"context"
	"log"
	"pamukkale_university/ent"

	_ "github.com/lib/pq"
)

func (d Postgre) main() {
	DBClient, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := DBClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Print(DBClient)
}