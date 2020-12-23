package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserUsecase(t *testing.T) {
	userUsecase := NewUserUsecase(&mockUserRepository{})
	assert.NotNil(t, userUsecase)
}

func TestUserUsecase_HandleUserRegister(t *testing.T) {
	u := NewUserUsecase(&mockUserRepository{})

	for _, tt := range []struct {
		name    string
		value   *User
		wantErr bool
	}{
		{
			name:    "Test Inserting Mock User",
			value:   &User{Username: mockUser.Username, Password: mockUser.Password, Email: mockUser.Email},
			wantErr: false,
		},
		{
			name:    "Test Inserting Empty Username ",
			value:   &User{Username: "", Password: "admin", Email: "admin@yahoo.com"},
			wantErr: true,
		},
		{
			name:    "Test Inserting Empty Password ",
			value:   &User{Username: "user3", Password: "", Email: "admin@yahoo.com"},
			wantErr: true,
		},
		{
			name:    "Test Inserting Empty Email ",
			value:   &User{Username: "noemail", Password: "admin", Email: ""},
			wantErr: true,
		},
	} {
		result, err := u.HandleUserRegister(tt.value)
		if !tt.wantErr {
			assert.Equal(t, result.Username, tt.value.Username)
			assert.NoError(t, err, tt.name)
			assert.NotEmpty(t, result.Password)

			if result.Username == mockUser.Username {
				mockUser.Password = result.Password
			}

		} else {
			assert.Error(t, err, tt.name)
		}
	}
}

func TestUserUsecase_HandleUserLogin(t *testing.T) {
	u := NewUserUsecase(&mockUserRepository{})

	for _, tt := range []struct {
		name    string
		value   *User
		wantErr bool
	}{
		{
			name:    "Login with Username & Password",
			value:   &User{Username: mockUser.Username, Password: "admin"},
			wantErr: false,
		},
		{
			name:    "Login Username yang tidak ada",
			value:   &User{Username: "kucing", Password: "admin"},
			wantErr: true,
		},
		{
			name:    "Login No Username",
			value:   &User{Username: "", Password: "admin"},
			wantErr: true,
		},
		{
			name:    "Login No Password",
			value:   &User{Username: "admin", Password: ""},
			wantErr: true,
		},
	} {
		result, err := u.HandleUserLogin(tt.value)
		if !tt.wantErr {
			assert.Equal(t, result.Username, tt.value.Username)
			assert.NoError(t, err, tt.name)
		} else {
			assert.Error(t, err, tt.name)
		}
	}
}

func TestUserUsecase_HandleDeleteUser(t *testing.T) {
	u := NewUserUsecase(&mockUserRepository{})

	err := u.HandleDeleteUser("1000")
	assert.NoError(t, err)
}

type mockUserRepository struct {
}

func (r *mockUserRepository) HandleUserLogin(username string, status bool) (*User, error) {
	if username == "gerin" {
		return mockUser, nil
	} else {
		return nil, errors.New("Username Not Found")
	}
}

func (r *mockUserRepository) HandleUserRegister(user *User) (*User, error) {
	return user, nil
}

func (r *mockUserRepository) HandleDeleteUser(id string) error {
	return nil
}

var mockUser = &User{
	ID:       "1000",
	Username: "gerin",
	Password: "admin",
	Email:    "gerin@admin.com",
	Status:   true,
}
