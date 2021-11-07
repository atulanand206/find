package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestMatchService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	Db := db.DB{}
	subscriberService := services.SubscriberService{Crud: db.SubscriberCrud{Db: Db}, TargetService: services.TargetService{}, Creators: services.Creators{}}
	service := services.MatchService{Crud: db.MatchCrud{Db: Db}, SubscriberService: subscriberService}
	playerService := services.PlayerService{Crud: db.PlayerCrud{Db: Db}}

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
