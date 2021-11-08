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

func TestDB(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)
	dbConns := []db.DBConn{db.NewMockDb(false), db.NewDb()}

	for _, dbConn := range dbConns {
		t.Run("create match", func(t *testing.T) {
			dbConn.CreateCollection("matches")
			game := tests.TestGame()
			err := dbConn.Create(game, "matches")
			assert.Nil(t, err)
			res, err := dbConn.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
			assert.Nil(t, err)
			var foundGame models.Game
			bson.Unmarshal(res, &foundGame)
			assert.Equal(t, game.Id, foundGame.Id)
		})

		t.Run("create and update match", func(t *testing.T) {
			dbConn.CreateCollection("matches")
			game := tests.TestGame()
			err := dbConn.Create(game, "matches")
			assert.Nil(t, err)
			res, err := dbConn.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
			assert.Nil(t, err)

			var foundGame models.Game
			bson.Unmarshal(res, &foundGame)
			assert.Equal(t, game.Id, foundGame.Id)
			assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)

			newGame := tests.TestGame()
			newGame.Id = game.Id
			dbConn.Update("matches", bson.M{"_id": newGame.Id}, newGame)
			res, err = dbConn.FindOne("matches", bson.M{"_id": newGame.Id}, &options.FindOneOptions{})
			assert.Nil(t, err)

			bson.Unmarshal(res, &foundGame)
			assert.Equal(t, newGame.Id, foundGame.Id)
			assert.Equal(t, newGame.Specs.Name, foundGame.Specs.Name)
		})

		t.Run("create and delete match", func(t *testing.T) {
			dbConn.CreateCollection("matches")
			game := tests.TestGame()
			err := dbConn.Create(game, "matches")
			assert.Nil(t, err)
			res, err := dbConn.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
			assert.Nil(t, err)

			var foundGame models.Game
			bson.Unmarshal(res, &foundGame)
			assert.Equal(t, game.Id, foundGame.Id)
			assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)

			dbConn.Delete("matches", bson.M{"_id": game.Id})
			res, _ = dbConn.FindOne("matches", bson.M{"_id": game.Id}, &options.FindOneOptions{})
			assert.Nil(t, res)
		})

		t.Run("create and find match", func(t *testing.T) {
			dbConn.CreateCollection("matches")
			game := tests.TestGame()
			err := dbConn.Create(game, "matches")
			assert.Nil(t, err)
			res, err := dbConn.Find("matches", bson.M{"_id": game.Id}, &options.FindOptions{})
			assert.Nil(t, err)

			assert.Equal(t, 1, len(res))
			var foundGame models.Game
			bson.Unmarshal(res[0], &foundGame)
			assert.Equal(t, game.Id, foundGame.Id)
			assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)
		})

		t.Run("create and find match using in filter", func(t *testing.T) {
			dbConn.CreateCollection("matches")
			game := tests.TestGame()
			err := dbConn.Create(game, "matches")
			assert.Nil(t, err)
			res, err := dbConn.Find("matches", bson.M{"_id": bson.M{"$in": []string{game.Id}}}, &options.FindOptions{})
			assert.Nil(t, err)

			assert.Equal(t, 1, len(res))
			var foundGame models.Game
			bson.Unmarshal(res[0], &foundGame)
			assert.Equal(t, game.Id, foundGame.Id)
			assert.Equal(t, game.Specs.Name, foundGame.Specs.Name)
		})
	}
}
