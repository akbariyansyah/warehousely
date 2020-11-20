package main

import (
	"fmt"
	"warehousely/config"
	"warehousely/config/database/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	// set up environment variable
	config.SetEnvironmentVariable()
	db := postgres.InitPostgresDatabase()
	fmt.Println(db)

	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello There")
	})

	source := fmt.Sprintf("%s:%s", config.MAIN_SERVER_HOST, config.MAIN_SERVER_PORT)
	g.Run(source)
}
