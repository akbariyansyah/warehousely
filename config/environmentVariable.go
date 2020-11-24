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
	DB_NAME,
	DATABASE_URL string

func SetEnvironmentVariable() {
	HOST = os.Getenv("HOST")

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8000" // Set up default value kalau tidak di set
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DATABASE_URL = os.Getenv("DATABASE_URL")

	// hardcode dlu hehehe
	DATABASE_URL = "postgres://dhopkyohgphwba:c1635497d128cb0658789be8531315041c2cfbbd4b5dfa426e46342140636636@ec2-52-3-4-232.compute-1.amazonaws.com:5432/d6hjiv7dbbhoci"

	log.Println("---------> SET UP ENVIRONMENT VARIABLES <---------")
	log.Println(`HOST =`, HOST)
	log.Println(`PORT =`, PORT)
	log.Println(`DB_USER =`, DB_USER)
	log.Println(`DB_PASSWORD =`, DB_PASSWORD)
	log.Println(`DB_HOST =`, DB_HOST)
	log.Println(`DB_PORT =`, DB_PORT)
	log.Println(`DB_NAME =`, DB_NAME)
	log.Println(`DATABASE_URL =`, DATABASE_URL)
	log.Println("--------------------------------------------------")
}
