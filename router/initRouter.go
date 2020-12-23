package router

import (
	"warehousely/api/domain/goods"
	"warehousely/api/domain/users"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func InitRouter(g *gin.Engine, db *pg.DB) {
	g.GET("/", healthCheck)
	users.InitUserRoutes("/users", g, db)
	goods.InitGoodsRoutes("/goods", g, db)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, "Hello There")
}
