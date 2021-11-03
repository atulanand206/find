package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func testSpecs() core.Specs {
	gameName, _ := gonanoid.New(8)
	return core.Specs{
		Name:      gameName,
		Teams:     2,
		Players:   2,
		Questions: 10,
		Rounds:    2,
		Points:    10,
	}
}

func testGame() core.Game {
	playerId, _ := gonanoid.New(8)
	playerName, _ := gonanoid.New(8)
	email, _ := gonanoid.New(8)
	quizmaster := core.Player{
		Id:    playerId,
		Name:  playerName,
		Email: email,
	}
	specs := testSpecs()
	game := core.InitNewMatch(quizmaster, specs)
	return game
}

func TestCreateAndFindMatch(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	game := testGame()
	crud.CreateMatch(game)

	res, err := crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Id != game.Id || res.Specs.Name != game.Specs.Name {
		t.Fatalf("match id mismatch")
	}
}

func TestFindMatchFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	id, _ := gonanoid.New(8)
	_, err := crud.FindMatch(id)
	if err == nil {
		t.Fatalf("match found")
	}
}

func TestUpdateMatchFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	game := testGame()
	count, _ := crud.UpdateMatch(game)
	if count {
		t.Fatalf("test failed")
	}
}

func TestUpdateMatch(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	game := testGame()
	crud.CreateMatch(game)

	res, err := crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Specs.Name != game.Specs.Name {
		t.Fatalf("specs not correct")
	}

	game.Specs.Name, _ = gonanoid.New(8)
	crud.UpdateMatch(game)

	res, err = crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Specs.Name != game.Specs.Name {
		t.Fatalf("specs not updated")
	}
}
