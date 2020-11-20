package users

import "github.com/go-pg/pg"

type UserRepository struct {
	db *pg.DB
}

type UserRepositoryInterface interface {
}

func NewUserRepository(db *pg.DB) UserRepositoryInterface {
	return &UserRepository{db}
}
