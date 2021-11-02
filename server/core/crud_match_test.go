package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TestCreateMatch(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	playerId, _ := gonanoid.New(8)
	quizmaster := core.InitNewPlayer(core.Player{
		Id:    playerId,
		Name:  "quizmaster",
		Email: "master@quiz.com",
	})
	specs := core.Specs{
		Name:      "test",
		Teams:     2,
		Players:   2,
		Questions: 10,
		Rounds:    2,
		Points:    10,
	}
	game := core.InitNewMatch(quizmaster, specs)
	crud := core.MatchCrud{}
	crud.CreateMatch(game)

	res, err := crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Id != game.Id {
		t.Fatalf("match id mismatch")
	}
}
