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
	gin.SetMode(gin.ReleaseMode)

	db := postgres.InitPostgresDatabase()
	g := gin.New()

	router.InitRouter(g, db)

	source := fmt.Sprintf("%s:%s", config.HOST, config.PORT)
	g.Run(source)
}
