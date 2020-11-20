package users

import "github.com/go-pg/pg"

type UserUsecase struct {
	UserRepository UserRepositoryInterface
}

type UserUsecaseInterface interface {
}

func NewUserUsecase(db *pg.DB) UserUsecaseInterface {
	return &UserUsecase{NewUserRepository(db)}
}
