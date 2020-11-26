package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"warehousely/config"
	"warehousely/config/database/postgres"
	"warehousely/middleware"
	"warehousely/router"
)

func main() {
	// set up environment variable
	config.SetEnvironmentVariable()
	gin.SetMode(gin.ReleaseMode)

	db := postgres.InitPostgresDatabase()
	g := gin.New()
	g.Use(middleware.HandleCORS())
	router.InitRouter(g, db)

	source := fmt.Sprintf("%s:%s", config.HOST, config.PORT)
	g.Run(source)
}
