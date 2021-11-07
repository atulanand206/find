package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TestPlayerCrud(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.PlayerCrud{Db: db.DB{}}

	t.Run("find or create player", func(t *testing.T) {
		player := tests.TestPlayer()
		res, err := crud.FindOrCreatePlayer(player)
		if err != nil {
			t.Fatalf("player %s not found", player.Id)
		}
		if res.Email != player.Email {
			t.Fatalf("player email mismatch")
		}
	})

	t.Run("find player fail", func(t *testing.T) {
		email, _ := gonanoid.New(8)
		_, err := crud.FindPlayer(email)
		if err == nil {
			t.Fatalf("test failed")
		}
	})

	t.Run("update player", func(t *testing.T) {
		player := tests.TestPlayer()
		crud.FindOrCreatePlayer(player)

		res, err := crud.FindPlayer(player.Email)
		if err != nil {
			t.Fatalf("player %s not found", player.Email)
		}

		name, _ := gonanoid.New(8)
		res.Name = name
		crud.UpdatePlayer(res)

		res, err = crud.FindPlayer(player.Email)
		if err != nil {
			t.Fatalf("player %s not found", player.Email)
		}
		if res.Name != name {
			t.Fatalf("specs not updated")
		}
	})
}
