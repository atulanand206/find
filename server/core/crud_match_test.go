package core_test

import (
	"fmt"
	"testing"

	"github.com/atulanand206/find/server/core"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func testGame() core.Game {
	playerId, _ := gonanoid.New(8)
	quizmaster := core.Player{
		Id:    playerId,
		Name:  "quizmaster",
		Email: "master@quiz.com",
	}
	specs := core.Specs{
		Name:      "test",
		Teams:     2,
		Players:   2,
		Questions: 10,
		Rounds:    2,
		Points:    10,
	}
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
	if res.Id != game.Id || res.Specs.Name != "test" {
		t.Fatalf("match id mismatch")
	}
}

func TestFindMatchFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	_, err := crud.FindMatch("game.Id")
	if err == nil {
		t.Fatalf("match found")
	}
}

func TestFindActiveMatches(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	game := testGame()
	crud.CreateMatch(game)

	res, err := crud.FindActiveMatches()
	if err != nil {
		t.Fatalf("matches %s not found", game.Id)
	}
	for _, match := range res {
		if match.Id == game.Id {
			return
		}
	}
	t.Fatalf("match id not present")
}

func TestUpdateMatchFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	crud := core.MatchCrud{}

	game := testGame()
	count, _ := crud.UpdateMatch(game)
	fmt.Println(count)
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
	if res.Specs.Name != "test" {
		t.Fatalf("specs not correct")
	}

	game.Specs.Name = "updated"
	crud.UpdateMatch(game)

	res, err = crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Specs.Name != "updated" {
		t.Fatalf("specs not updated")
	}
}
