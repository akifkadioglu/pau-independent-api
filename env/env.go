package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DB_CONNECTION   = "DB_CONNECTION"
	DB_HOST         = "DB_HOSTF"
	DB_PORT         = "DB_PORT"
	DB_DATABASE     = "DB_DATABASE"
	DB_USERNAME     = "DB_USERNAME"
	DB_PASSWORD     = "DB_PASSWORD"
	DB_EXTERNAL_URL = "DB_EXTERNAL_URL"

	PORT = "PORT"
)

func InitProd() {
	//err := godotenv.Load("./etc/secrets/.env")
//
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
}

func InitTest() {
	err := godotenv.Load("./../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Setenv(key string, value string) {
	os.Setenv(key, value)
}

func Getenv(key string) string {
	return os.Getenv(key)
}
