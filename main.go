package main

import (
	"pamukkale_university/database"
	"pamukkale_university/env"
	"pamukkale_university/routes"
)

func main() {
	env.InitProd()
	database.Init()
	routes.InitProd()
}
