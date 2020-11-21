package goods

import (
	"github.com/go-pg/pg"
	"log"
)

type GoodsRepositoryInterface interface {
	GetAllGoods(id string) ([]*Goods,error)
	CreateGoods(newGoods *Goods) error
	UpdateGoods(newGoods *Goods) error
	DeleteGoods(id string) error
}
type GoodsRepository struct {
	DB *pg.DB
	goodsRepo GoodsRepositoryInterface
}

func newGoodsRepository(db *pg.DB) *GoodsRepository {
	return &GoodsRepository{DB: db}
}

func (g GoodsRepository) GetAllGoods(id string) ([]*Goods, error) {
	return nil,nil
}

func (g GoodsRepository) CreateGoods(newGoods *Goods) error {
	tx,err := g.DB.Begin()
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : 1",err)
		return err
	}
	//stmt,err := tx.Prepare()
	//if err != nil {
	//	tx.Rollback()
	//	log.Println("ERROR ->",err)
	//	return err
	//}

	_ , err = tx.Exec("INSERT INTO m_goods (name,stock,category_id,user_id,status,image_path) VALUES (?,?,?,?,?,?)",newGoods.Name,newGoods.Stock,newGoods.CategoryID,newGoods.UserID,newGoods.Status,newGoods.ImagePath)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : 3",err)
		return err
	}
	tx.Commit()

	return nil
}

func (g GoodsRepository) UpdateGoods(newGoods *Goods) error {
	panic("implement me")
}

func (g GoodsRepository) DeleteGoods(id string) error {
	panic("implement me")
}

