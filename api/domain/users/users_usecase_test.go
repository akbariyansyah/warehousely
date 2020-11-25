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
	u := &UserUsecase{UserRepository: NewUserRepository(db)}

	for _, tt := range []struct {
		name    string
		value   *User
		wantErr bool
	}{
		{
			name:    "Test Inserting Random String",
			value:   &User{Username: randomString(8), Password: randomString(15), Email: randomString(11)},
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
			assert.NoError(t, err, tt.name)
			assert.NotEmpty(t, result.Password)
		} else {
			assert.Error(t, err, tt.name)
		}
	}
}

func TestUserUsecase_HandleUserLogin(t *testing.T) {
	u := &UserUsecase{UserRepository: NewUserRepository(db)}

	for _, tt := range []struct {
		name    string
		value   *User
		wantErr bool
	}{
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
		_, err := u.HandleUserLogin(tt.value)
		if !tt.wantErr {
			assert.NoError(t, err, tt.name)
		} else {
			assert.Error(t, err, tt.name)
		}
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
