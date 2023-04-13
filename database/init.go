package database

import "pamukkale_university/ent"

type PostgreSQL struct{}
type MySQL struct{}

type IDatabase interface {
	main()
}

var err error
var DBClient *ent.Client

func Init() {
	var db IDatabase = PostgreSQL{}
	db.main()
}