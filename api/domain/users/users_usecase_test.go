package users

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-pg/pg"
)

func TestNewUserUsecase(t *testing.T) {
	if db == nil {
		opt, err := pg.ParseURL(DATABASE_URL)
		if err != nil {
			t.Errorf("Cannot connect to database")
		}
		db = pg.Connect(opt)
	}

	t.Run("Connecting To Database", func(t *testing.T) {
		result := NewUserUsecase(db)
		assert.NotNil(t, result)
	})
}

func TestUserUsecase_HandleUserRegister(t *testing.T) {
	type args struct {
		user *User
	}

	tests := []struct {
		name    string
		u       *UserUsecase
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:    "Test Inserting New User",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "root", Password: "admin", Email: "admin@yahoo.com"}},
			wantErr: false,
		},
		{
			name:    "Test Inserting Empty Username ",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "", Password: "admin", Email: "admin@yahoo.com"}},
			wantErr: true,
		},
		{
			name:    "Test Inserting Empty Password ",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "user3", Password: "", Email: "admin@yahoo.com"}},
			wantErr: true,
		},
		{
			name:    "Test Inserting Empty Email ",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "noemail", Password: "admin", Email: ""}},
			wantErr: true,
		},
		{
			name:    "Test Inserting Random String",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: randomString(7), Password: randomString(15), Email: randomString(10)}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.HandleUserRegister(tt.args.user)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, got.Username, tt.args.user.Username, "Username should be equal")
			}
		})
	}
}

func TestUserUsecase_HandleUserLogin(t *testing.T) {
	type args struct {
		user *User
	}

	tests := []struct {
		name    string
		u       *UserUsecase
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:    "Login User with right Password",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "root", Password: "admin"}},
			wantErr: false,
		},
		{
			name:    "Login User Wrong Password",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "root", Password: "admin123"}},
			wantErr: true,
		},
		{
			name:    "Login No Username",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "", Password: "admin"}},
			wantErr: true,
		},
		{
			name:    "Login No Password",
			u:       &UserUsecase{UserRepository: NewUserRepository(db)},
			args:    args{&User{Username: "admin", Password: ""}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.u.HandleUserLogin(tt.args.user)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				// assert.NotNil(t, res.Username)
			}
		})
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
