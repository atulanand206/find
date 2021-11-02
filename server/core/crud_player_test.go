package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func testPlayerWithoutId() core.Player {
	player := core.Player{
		Name:  "quizmaster",
		Email: "master@quiz.com",
	}
	return player
}

func testPlayer(email string) core.Player {
	playerId, _ := gonanoid.New(8)
	player := core.Player{
		Id:    playerId,
		Name:  "quizmaster",
		Email: email,
	}
	return player
}

func TestFindOrCreatePlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerCrud{}

	player := testPlayerWithoutId()
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

	email := "master@quiz.co.m"
	_, err := crud.FindPlayer(email)
	if err == nil {
		t.Fatalf("test failed")
	}
}

func TestUpdatePlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.PlayerCrud{}

	player := testPlayer("magic@dark.gov")
	crud.FindOrCreatePlayer(player)

	res, err := crud.FindPlayer(player.Email)
	if err != nil {
		t.Fatalf("player %s not found", player.Email)
	}

	res.Name = "updated"
	crud.UpdatePlayer(res)

	res, err = crud.FindPlayer(player.Email)
	if err != nil {
		t.Fatalf("player %s not found", player.Email)
	}
	if res.Name != "updated" {
		t.Fatalf("specs not updated")
	}
}
