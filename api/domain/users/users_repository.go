package users

import (
	"github.com/go-pg/pg"
)

type UserRepository struct {
	db *pg.DB
}

type UserRepositoryInterface interface {
	HandleUserLogin(username string, status bool) (*User, error)
	HandleUserRegister(user *User) (*User, error)
}

func NewUserRepository(db *pg.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (u *UserRepository) HandleUserRegister(user *User) (*User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	statement, err := tx.Prepare(`INSERT INTO m_user(username, password, email) VALUES($1,$2,$3) RETURNING id`)
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(&user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return user, nil
}

func (u *UserRepository) HandleUserLogin(username string, status bool) (*User, error) {
	user := new(User)

	// ngambil satu data berdasarkan username dan status yang aktif
	_, err := u.db.QueryOne(user, "SELECT * FROM m_user WHERE username=? AND status=?", username, status)
	if err != nil {
		return nil, err
	}

	return user, nil
}
