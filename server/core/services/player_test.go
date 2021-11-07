package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestCreateAndFindPlayer(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.PlayerService{}

	player, err := service.FindOrCreatePlayer(tests.TestPlayer())
	if err != nil {
		t.Fatalf("player %s not found", player.Id)
	}
}

func TestFindPlayerByEmailFail(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.PlayerService{}

	_, err := service.FindPlayerByEmail("random-bullshit@asdd.com")
	if err == nil {
		t.Fatalf("test failed")
	}
}

func TestFindPlayer(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.PlayerService{}

	player := tests.TestPlayer()
	player, _ = service.FindOrCreatePlayer(player)
	_, err := service.FindPlayerByEmail(player.Email)
	if err != nil {
		t.Fatalf("test failed")
	}
}
