package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		res, _ := mockDb.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
		var foundGame models.Game
		bson.Unmarshal(res, &foundGame)
		assert.Equal(t, game.Id, foundGame.Id)
	})

	t.Run("create and update match", func(t *testing.T) {
		mockDb.CreateCollection("matches")
		assert.Equal(t, 1, len(mockDb.Data))
		game := tests.TestGame()
		mockDb.Create(game, "matches")
		res, _ := mockDb.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
		var foundGame models.Game
		bson.Unmarshal(res, &foundGame)
		assert.Equal(t, game.Id, foundGame.Id)
		assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)

		newGame := tests.TestGame()
		newGame.Id = game.Id
		mockDb.Update("matches", bson.M{"_id": newGame.Id}, newGame)
		res, _ = mockDb.FindOne("matches", bson.M{"_id": newGame.Id}, &options.FindOneOptions{})
		bson.Unmarshal(res, &foundGame)
		assert.Equal(t, newGame.Id, foundGame.Id)
		assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)
	})

	t.Run("create and delete match", func(t *testing.T) {
		mockDb.CreateCollection("matches")
		assert.Equal(t, 1, len(mockDb.Data))
		game := tests.TestGame()
		mockDb.Create(game, "matches")
		res, _ := mockDb.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
		var foundGame models.Game
		bson.Unmarshal(res, &foundGame)
		assert.Equal(t, game.Id, foundGame.Id)
		assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)

		mockDb.Delete("matches", bson.M{"_id": game.Id})
		res, _ = mockDb.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
		assert.Nil(t, res)
	})
}
