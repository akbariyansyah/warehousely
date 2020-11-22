package config

import (
	"log"
	"os"
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
	MAIN_SERVER_PORT = os.Getenv(MAIN_SERVER_PORT)
	if MAIN_SERVER_PORT == "" {
		MAIN_SERVER_PORT = "8000"
	}

	// DB_USER = "postgres"
	// DB_PASSWORD = "P@ssW02d123"
	// DB_HOST = "127.0.0.1"
	// DB_PORT = "5432"
	// DB_NAME = "warehousely"

	DB_USER = "root"
	DB_PASSWORD = "admin"
	DB_HOST = "127.0.0.1"
	DB_PORT = "5432"
	DB_NAME = "default"
	log.Println("-------------------------------------------")
}
