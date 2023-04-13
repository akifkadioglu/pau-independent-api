package env

import (
	"os"
	"strconv"
)

const (
	DB_CONNECTION int = iota
	DB_HOST
	DB_PORT
	DB_DATABASE
	DB_USERNAME
	DB_PASSWORD

	PORT
)

func Init() {
	Setenv(DB_CONNECTION, "")
	Setenv(DB_HOST, "")
	Setenv(DB_PORT, "")
	Setenv(DB_DATABASE, "")
	Setenv(DB_USERNAME, "")
	Setenv(DB_PASSWORD, "")

	Setenv(PORT, "")
}

func Setenv(key int, value string) {
	os.Setenv(strconv.Itoa(key), value)
}

func Getenv(key int) string {
	return os.Getenv(strconv.Itoa(key))
}