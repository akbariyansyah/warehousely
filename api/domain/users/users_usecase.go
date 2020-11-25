package users

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-pg/pg"
)

type UserUsecase struct {
	UserRepository UserRepositoryInterface
}

type UserUsecaseInterface interface {
	HandleUserLogin(user *User) (*User, error)
	HandleUserRegister(user *User) (*User, error)
	HandleDeleteUser(id string) error
}

func NewUserUsecase(db *pg.DB) UserUsecaseInterface {
	return &UserUsecase{NewUserRepository(db)}
}

func (u *UserUsecase) HandleUserLogin(user *User) (*User, error) {
	if user.Username == "" || user.Password == "" {
		return nil, errors.New(`Username or Password cannot empty`)
	}

	result, err := u.UserRepository.HandleUserLogin(user.Username, true)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// membandingkan user input (user.password) dengan password yang ada di db (result.password)
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Username atau Password salah")
	}

	log.Println(`User Login ->`, result.Username)

	result.Password = ""
	return result, nil
}

func (u *UserUsecase) HandleUserRegister(user *User) (*User, error) {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return nil, errors.New(`Username or Password cannot empty`)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user.Password = string(hash)
	result, err := u.UserRepository.HandleUserRegister(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(`User Register ->`, result.Username)

	return result, nil
}

func (u *UserUsecase) HandleDeleteUser(id string) error {
	err := u.UserRepository.HandleDeleteUser(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println(`Delete User id ->`, id)

	return nil
}
