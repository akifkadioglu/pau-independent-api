package database

import (
	"context"
	"log"
	"pamukkale_university/ent"
	"pamukkale_university/env"

	_ "github.com/go-sql-driver/mysql"
)

func (d MySQL) main() {
	DBClient, err = ent.Open("mysql", env.Getenv(env.DB_USERNAME) + ":" + env.Getenv(env.DB_PASSWORD) + "@tcp(" + env.Getenv(env.DB_HOST) + ":" + env.Getenv(env.DB_PORT) + ")/" + env.Getenv(env.DB_DATABASE) + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := DBClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}