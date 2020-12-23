package users

import (
	"testing"
)

func TestNewUserController(t *testing.T) {

}

func TestUserController_HandleUserLogin(t *testing.T) {

}

func TestUserController_HandleUserRegister(t *testing.T) {

}

func TestUserController_HandleDeleteUser(t *testing.T) {

}

type mockUserUsecase struct{}

func (u *mockUserUsecase) HandleUserLogin(user *User) (*User, error) {
	return user, nil
}

func (u *mockUserUsecase) HandleUserRegister(user *User) (*User, error) {
	return user, nil
}

func (u *mockUserUsecase) HandleDeleteUser(id string) error {
	return nil
}
