package main

import (
	"log"
	"pamukkale_university/database"
	"pamukkale_university/env"
	"pamukkale_university/routes"
)

func main() {
	env.InitProd()
	database.Init()
	routes.InitProd()
	log.Println(env.Getenv(env.DB_CONNECTION))
	log.Println(env.Getenv(env.DB_EXTERNAL_URL))
	log.Println(env.Getenv(env.PORT))
}
