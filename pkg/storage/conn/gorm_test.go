package conn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckConnect(t *testing.T) {
	t.Run("when connect to mysql will return DB and nil", func(t *testing.T) {
		db := CheckConnect()
		assert.NotEmpty(t, db, nil)
	})
}
