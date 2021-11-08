package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestQuestionService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.QuestionService{db.QuestionCrud{Db: db.NewMockDb()}}

	t.Run("add question", func(t *testing.T) {
		// Create a new question
		question := models.NewQuestion{
			Statements: []string{"test", "question"},
			Answer:     []string{"Test Answer"},
		}

		err := service.AddQuestion("question.Tag", []models.NewQuestion{question})
		if err != nil {
			t.Fatalf("Error adding question: %s", err)
		}
	})

	t.Run("find question for match", func(t *testing.T) {
		game := tests.TestGame()

		_, err := service.FindQuestionForMatch(game)
		if err != nil {
			t.Fatalf("Error finding question: %s", err)
		}
	})
}
