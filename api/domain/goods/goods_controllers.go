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

	result, err := u.goodsUsecase.GetAllGoods("1")
	if err != nil {
		c.JSON(400, "Failed")
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"status": "ok",
			"data":   result,
		})
	}

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
func (u *GoodsController) HandleDeleteGoods(c *gin.Context) {
	id := c.Param("id")

	err := u.goodsUsecase.DeleteGoods(id)
	if err != nil {
		log.Println(err)
		c.JSON(501, "Oops there was an error")
	} else {
		c.JSON(200, "Delete Goods success")
	}
}
func (u *GoodsController) HandlePutGoods(c *gin.Context) {
	var goods Goods
	_ = c.BindJSON(&goods)
	err := u.goodsUsecase.UpdateGoods(&goods)
	if err != nil {
		log.Println(err)
		c.JSON(501, "Oops there was an error")
	} else {
		c.JSON(200, "Change Goods success")
	}
}
