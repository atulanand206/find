package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func testPlayer() core.Player {
	playerId, _ := gonanoid.New(8)
	playerName, _ := gonanoid.New(8)
	email, _ := gonanoid.New(8)
	player := core.Player{
		Id:    playerId,
		Name:  playerName,
		Email: email,
	}
	return player
}

func TestFindOrCreatePlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerCrud{}

	player := testPlayer()
	res, err := crud.FindOrCreatePlayer(player)
	if err != nil {
		t.Fatalf("player %s not found", player.Id)
	}
	if res.Email != player.Email {
		t.Fatalf("player email mismatch")
	}
}

func TestFindPlayerFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerCrud{}

	email, _ := gonanoid.New(8)
	_, err := crud.FindPlayer(email)
	if err == nil {
		t.Fatalf("test failed")
	}
}

func TestUpdatePlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerCrud{}

	player := testPlayer()
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
}
