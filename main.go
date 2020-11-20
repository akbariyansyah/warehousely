package main

import (
	"fmt"
	"log"
	"warehousely/config"
	"warehousely/config/database/postgres"

	"github.com/gin-gonic/gin"
)

const (
	ServerHost = "localhost"
	ServerPort = "8000"
)

type Goods struct {
	ID    int
	Name  string
	Type  string
	Stock int
}

func main() {
	// set up environment variable
	config.SetEnvironmentVariable()
	db := postgres.InitPostgresDatabase()
	fmt.Println(db)
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello There")
	})

	log.Printf("Starting server at %s:%s", ServerHost, ServerPort)
	source := fmt.Sprintf("%s:%s", config.MAIN_SERVER_HOST, config.MAIN_SERVER_PORT)
	g.Run(source)
}
