package goods

import (
	"github.com/go-pg/pg"
	"log"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	goodsUsecase GoodsUsecaseInterface
}

func NewGoodsController(db *pg.DB) *GoodsController {
	return &GoodsController{goodsUsecase: NewGoodsUsecase(db)}
}

func (u *GoodsController) HandleGetAllGoods(c *gin.Context) {
	c.JSON(200, "Get All Goods There")
}

func (u *GoodsController) HandlePostNewGoods(c *gin.Context) {
	var goods Goods
	err := c.BindJSON(&goods)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Oops there was an error")
	}
	err = u.goodsUsecase.CreateGoods(&goods)
	if err != nil {
		log.Println(err)
		c.JSON(501, "Oops there was an error")
	} else {
		c.JSON(200, "Create Goods success")
	}
}
