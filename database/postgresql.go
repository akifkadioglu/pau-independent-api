package database

import (
	"context"
	"log"
	"pamukkale_university/ent"
	"pamukkale_university/env"

	_ "github.com/lib/pq"
)

func (d PostgreSQL) main() {
	DBClient, err = ent.Open(env.Getenv(env.DB_CONNECTION), env.Getenv(env.DB_EXTERNAL_URL))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := DBClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Print(DBClient)
}
