package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
)

func TestCreateAndFindPlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.PlayerService{}

	player, err := service.FindOrCreatePlayer(testPlayerWithoutId())
	if err != nil {
		t.Fatalf("player %s not found", player.Id)
	}
}

func TestFindPlayerByEmailFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerService{}

	_, err := crud.FindPlayerByEmail("random-bullshit@asdd.com")
	if err == nil {
		t.Fatalf("test failed")
	}
}

func TestFindPlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerService{}

	player := testPlayerWithoutId()
	player, _ = crud.FindOrCreatePlayer(player)
	_, err := crud.FindPlayerByEmail(player.Email)
	if err != nil {
		t.Fatalf("test failed")
	}
}
