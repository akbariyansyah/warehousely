package users

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-pg/pg"
)

var (
	DATABASE_URL = "postgres://root:admin@localhost:5432/default?sslmode=disable"
)

var db *pg.DB

func TestNewUserRepository(t *testing.T) {
	if db == nil {
		opt, err := pg.ParseURL(DATABASE_URL)
		if err != nil {
			t.Errorf("Cannot connect to database")
		}
		db = pg.Connect(opt)
	}

	t.Run("Connecting To Database", func(t *testing.T) {
		got := NewUserRepository(db)
		assert.NotNil(t, got)
	})
}
