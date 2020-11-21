package goods

import "github.com/go-pg/pg"

type GoodsUsecaseInterface interface {
	GetAllGoods(id string) ([]*Goods,error)
	CreateGoods(newGoods *Goods) error
}
type GoodsUsecase struct {
	goodsUsecase GoodsUsecaseInterface
	goodsRepo GoodsRepositoryInterface
}

func (g GoodsUsecase) GetAllGoods(id string) ([]*Goods, error) {
	panic("implement me")
}

func (g GoodsUsecase) CreateGoods(newGoods *Goods) error {
	err := g.goodsRepo.CreateGoods(newGoods)
	if err != nil {
		return err
	}
	return nil
}

func NewGoodsUsecase(db *pg.DB) *GoodsUsecase {
	return &GoodsUsecase{goodsRepo: newGoodsRepository(db)}
}