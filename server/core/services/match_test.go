package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestCreateMatch(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.MatchService{}

	quizmaster := tests.TestPlayer()
	specs := models.Specs{
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

func TestFindActiveMatchesForPlayer(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	playerService := services.PlayerService{}
	quizmaster, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())

	matchService := services.MatchService{}
	specs := tests.TestSpecs()
	game, err := matchService.CreateMatch(quizmaster, specs)
	if err != nil {
		t.Fatalf("game not created")
	}
	games, err := matchService.FindActiveMatchesForPlayer(quizmaster.Id)
	if err != nil {
		t.Fatalf("games not found")
	}
	for _, g := range games {
		if g.Id == game.Id && g.CanJoin {
			return
		}
	}
	t.Fatalf("game not found in active games")
}
