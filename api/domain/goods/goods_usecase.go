package goods

import "github.com/go-pg/pg"

type GoodsUsecaseInterface interface {
	GetAllGoods(id string) (*[]Goods, error)
	CreateGoods(newGoods *Goods) error
	UpdateGoods(newGoods *Goods) error
	DeleteGoods(id string) error
}

func (g GoodsUsecase) UpdateGoods(newGoods *Goods) error {
	err := g.goodsRepo.UpdateGoods(newGoods)
	if err != nil {
		return err
	}
	return nil
}

func (g GoodsUsecase) DeleteGoods(id string) error {
	err := g.goodsRepo.DeleteGoods(id)
	if err != nil {
		return err
	}
	return nil
}

type GoodsUsecase struct {
	goodsUsecase GoodsUsecaseInterface
	goodsRepo    GoodsRepositoryInterface
}

func (g GoodsUsecase) GetAllGoods(id string) (*[]Goods, error) {
	result, err := g.goodsRepo.GetAllGoods(id)
	if err != nil {
		return nil, err
	}
	return result, nil
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
