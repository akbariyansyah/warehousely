package postgres

import (
	"context"
	"log"
	"warehousely/config"

	"github.com/go-pg/pg"
)

// https://pg.uptrace.dev/ <- dokumentasinya
// untuk init database postgree
func InitPostgresDatabase() *pg.DB {
	db := new(pg.DB)

	// Ada dua cara untuk konek ke database, dengan URL atau config
	// kalau DATABASE_URL tidak kosong berarti konekin nya pake URL
	if config.DATABASE_URL != "" {
		opt, err := pg.ParseURL(config.DATABASE_URL)
		if err != nil {
			panic(err)
		}
		db = pg.Connect(opt)
	} else {
		db = pg.Connect(&pg.Options{
			Addr:     config.DB_HOST + ":" + config.DB_PORT,
			User:     config.DB_USER,
			Password: config.DB_PASSWORD,
			Database: config.DB_NAME,
		})
	}

	// untuk cek database bisa konek atau tidak
	_, err := db.ExecContext(context.Background(), "SELECT 1")
	if err != nil {
		log.Println(`Cannot Connect to Database ->`, err.Error())
	} else {
		log.Println(`Successful connect to Database`)
	}

	return db
}
