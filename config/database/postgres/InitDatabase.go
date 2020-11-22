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
	db := pg.Connect(&pg.Options{
		Addr:     config.DB_HOST + ":" + config.DB_PORT,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
		Database: config.DB_NAME,
	})

	// untuk cek database bisa konek atau tidak
	_, err := db.ExecContext(context.Background(), "SELECT 1")
	if err != nil {
		log.Println(`Cannot Connect to Database ->`, err.Error())
	} else {
		log.Println(`Successful connect to Database`)
	}

	return db
}
