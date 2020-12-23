package goods

import (
	"github.com/go-pg/pg"

	"github.com/gin-gonic/gin"
)

func InitGoodsRoutes(mainRoute string, g *gin.Engine, db *pg.DB) {
	r := g.Group(mainRoute)
	goodsController := NewGoodsController(db)

	//http://localhost:8080/goods
	r.GET("", goodsController.HandleGetAllGoods)

	//http://localhost:8080/goods
	r.POST("", goodsController.HandlePostNewGoods)
	r.PUT("", goodsController.HandlePutGoods)
	r.DELETE("/:id", goodsController.HandleDeleteGoods)
}
