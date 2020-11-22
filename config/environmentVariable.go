package config

import (
	"log"
	"os"
)

var HOST,
	PORT,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_NAME string

func SetEnvironmentVariable() {
	HOST = os.Getenv("HOST")

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8000" // Set up default value
	}

	DB_USER = "root"
	DB_PASSWORD = "admin"
	DB_HOST = "127.0.0.1"
	DB_PORT = "5432"
	DB_NAME = "default"

	log.Println("---------> SET UP ENVIRONMENT VARIABLES <---------")
	log.Println(`HOST =`, HOST)
	log.Println(`PORT =`, PORT)
	log.Println(`DB_USER =`, DB_USER)
	log.Println(`DB_PASSWORD =`, DB_PASSWORD)
	log.Println(`DB_HOST =`, DB_HOST)
	log.Println(`DB_PORT =`, DB_PORT)
	log.Println(`DB_NAME =`, DB_NAME)
	log.Println("--------------------------------------------------")
}
