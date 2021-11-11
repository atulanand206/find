package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

func TestPlayerCrud(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.PlayerCrud{Db: db.NewDb()}

	t.Run("find or create player", func(t *testing.T) {
		player := tests.TestPlayer()
		res, err := crud.FindOrCreatePlayer(player)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, player.Email, res.Email, "player email must match")
	})

	t.Run("find player fail", func(t *testing.T) {
		email, _ := gonanoid.New(8)
		_, err := crud.FindPlayer(email)
		assert.NotNil(t, err, "error must not be nil")
	})

	t.Run("find players", func(t *testing.T) {
		player := tests.TestPlayer()
		res, err := crud.FindOrCreatePlayer(player)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, player, res, "player must be equal")

		players, err := crud.FindPlayers([]string{player.Id})
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 1, len(players), "players count must be equal to one")
		assert.Equal(t, player, players[0], "player must be equal")
	})

	t.Run("update player", func(t *testing.T) {
		player := tests.TestPlayer()
		crud.FindOrCreatePlayer(player)

		res, err := crud.FindPlayer(player.Email)
		assert.Nil(t, err, "error must be nil")

		name, _ := gonanoid.New(8)
		res.Name = name
		_, err = crud.UpdatePlayer(res)
		assert.Nil(t, err)

		res, err = crud.FindPlayer(player.Email)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, name, res.Name, "player must be equal")
	})
}
