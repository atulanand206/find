package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.Init(db.DB{})

	t.Run("create quiz", func(t *testing.T) {
		quizmaster := tests.TestPlayer()
		specs := tests.TestSpecs()
		response, err := service.GenerateCreateGameResponse(quizmaster, specs)
		if err != nil {
			t.Fatalf("game not created")
		}
		game := response.Quiz
		if game.QuizMaster.Id != quizmaster.Id {
			t.Fatalf("quizmaster not set")
		}
	})

}
