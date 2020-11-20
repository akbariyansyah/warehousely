package goods

import (
	"github.com/go-pg/pg"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	// goodsUsecase GoodsUsecaseInterface
}

func NewGoodsController(db *pg.DB) *GoodsController {
	// return &GoodsController{NewGoodsUsecase(db)}
	return &GoodsController{}
}

func (u *GoodsController) HandleGetAllGoods(c *gin.Context) {
	c.JSON(200, "Get All Goods There")
}

func (u *GoodsController) HandlePostNewGoods(c *gin.Context) {
	c.JSON(200, "Post New Goods There")
}
