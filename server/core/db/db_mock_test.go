package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
)

func TestMockDB(t *testing.T) {
	mockDb := &db.MockDB{}
	mockDb.Init()
	t.Run("create match collection", func(t *testing.T) {
		mockDb.CreateCollection("matches")
		assert.Equal(t, 1, len(mockDb.Data))
	})

	t.Run("create match", func(t *testing.T) {
		mockDb.CreateCollection("matches")
		assert.Equal(t, 1, len(mockDb.Data))
		game := tests.TestGame()
		mockDb.Create(game, "matches")
		assert.Equal(t, 1, len(mockDb.Data["matches"]))
	})
}
