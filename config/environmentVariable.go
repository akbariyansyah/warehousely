package config

import (
	"log"
)

var MAIN_SERVER_HOST,
	MAIN_SERVER_PORT,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_NAME string

func SetEnvironmentVariable() {
	MAIN_SERVER_HOST = "0.0.0.0"
	MAIN_SERVER_PORT = "8080"

	DB_USER = "root"
	DB_PASSWORD = "admin"
	DB_HOST = "127.0.0.1"
	DB_PORT = "5432"
	DB_NAME = "default"

	log.Println("-------------------------------------------")
}