package goods

import (
	"github.com/go-pg/pg"
	"log"
)

type GoodsRepositoryInterface interface {
	GetAllGoods(id string) (*[]Goods, error)
	CreateGoods(newGoods *Goods) error
	UpdateGoods(newGoods *Goods) error
	DeleteGoods(id string) error
}
type GoodsRepository struct {
	db        *pg.DB
	goodsRepo GoodsRepositoryInterface
}

func newGoodsRepository(db *pg.DB) *GoodsRepository {
	return &GoodsRepository{db: db}
}

func (g GoodsRepository) GetAllGoods(id string) (*[]Goods, error) {
	var allGoods []Goods

	_, err := g.db.Query(&allGoods, "SELECT * FROM m_goods where status = true and user_id = ? ", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &allGoods, nil
}

func (g GoodsRepository) CreateGoods(newGoods *Goods) error {
	tx, err := g.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : 1", err)
		return err
	}

	_, err = tx.Exec("INSERT INTO m_goods (name,stock,category_id,user_id,status,image_path) VALUES (?,?,?,?,?,?)", newGoods.Name, newGoods.Stock, newGoods.CategoryID, newGoods.UserID, newGoods.Status, newGoods.ImagePath)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : 3", err)
		return err
	}
	tx.Commit()

	return nil
}

func (g GoodsRepository) UpdateGoods(newGoods *Goods) error {
	tx, err := g.db.Begin()

	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE m_goods set name = $1,stock = $2,image_path = $3,updated = CURRENT_DATE where id =$4 ")
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	_, err = stmt.Exec(newGoods.Name, newGoods.Stock, newGoods.ImagePath, newGoods.ID)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	tx.Commit()
	return nil
}

func (g GoodsRepository) DeleteGoods(id string) error {
	tx, err := g.db.Begin()

	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE m_goods set status = false where id =$1 ")
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	tx.Commit()
	return nil

}
