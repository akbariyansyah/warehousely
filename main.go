package main

import (
	"fmt"
	"warehousely/config"
	"warehousely/config/database/postgres"
	"warehousely/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// set up environment variable
	config.SetEnvironmentVariable()

	db := postgres.InitPostgresDatabase()
	g := gin.Default()

	router.InitRouter(g, db)

	source := fmt.Sprintf("%s:%s", config.MAIN_SERVER_HOST, config.MAIN_SERVER_PORT)
	g.Run(source)
}
