package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestPlayerService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.PlayerService{db.PlayerCrud{Db: db.NewMockDb(true)}}

	t.Run("create and find player", func(t *testing.T) {
		player, err := service.FindOrCreatePlayer(tests.TestPlayer())
		if err != nil {
			t.Fatalf("player %s not found", player.Id)
		}
	})

	t.Run("find player by email", func(t *testing.T) {
		_, err := service.FindPlayerByEmail("random-bullshit@asdd.com")
		if err == nil {
			t.Fatalf("test failed")
		}
	})

	t.Run("find player", func(t *testing.T) {
		player := tests.TestPlayer()
		player, _ = service.FindOrCreatePlayer(player)
		_, err := service.FindPlayerByEmail(player.Email)
		if err != nil {
			t.Fatalf("test failed")
		}
	})
}
