package config

import (
	"log"
	"os"
)

var MAIN_SERVER_HOST,
	PORT,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_NAME string

func SetEnvironmentVariable() {
	MAIN_SERVER_HOST = "127.0.0.1"
	PORT = os.Getenv(PORT)
	if PORT == "" {
		PORT = "8000"
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
