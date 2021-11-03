package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
)

func TestCreateMatch(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.MatchService{}

	quizmaster := testPlayer()
	specs := core.Specs{
		Players: 1,
	}
	game, err := service.CreateMatch(quizmaster, specs)
	if err != nil {
		t.Fatalf("game not created")
	}
	if game.QuizMaster.Id != quizmaster.Id {
		t.Fatalf("quizmaster not set")
	}
}

// func TestFindActiveMatchesForPlayer(t *testing.T) {
// 	teardown := Setup(t)
// 	defer teardown(t)

// 	playerService := core.PlayerService{}
// 	quizmaster, _ := playerService.FindOrCreatePlayer(testPlayer())

// 	matchService := core.MatchService{}
// 	specs := testSpecs()
// 	game, err := matchService.CreateMatch(quizmaster, specs)
// 	if err != nil {
// 		t.Fatalf("game not created")
// 	}
// 	games, err := matchService.FindActiveMatchesForPlayer(quizmaster.Id)
// 	if err != nil {
// 		t.Fatalf("games not found")
// 	}
// 	for _, g := range games {
// 		if g.Id == game.Id {
// 			return
// 		}
// 	}
// 	t.Fatalf("game not found in active games")
// }
