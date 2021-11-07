package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestMatchService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.MatchService{}

	t.Run("create match", func(t *testing.T) {
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
	})

	t.Run("find active matches", func(t *testing.T) {
		specs := tests.TestSpecs()

		playerService := services.PlayerService{}
		quizmaster, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())

		game, err := service.CreateMatch(quizmaster, specs)
		if err != nil {
			t.Fatalf("game not created")
		}
		games, err := service.FindActiveMatchesForPlayer(quizmaster.Id)
		if err != nil {
			t.Fatalf("games not found")
		}
		for _, g := range games {
			if g.Id == game.Id && g.CanJoin {
				return
			}
		}
		t.Fatalf("game not found in active games")
	})
}
