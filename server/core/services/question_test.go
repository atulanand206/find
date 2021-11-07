package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestAddQuestion(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.QuestionService{}

	// Create a new question
	question := models.NewQuestion{
		Statements: []string{"test", "question"},
		Answer:     []string{"Test Answer"},
	}

	err := service.AddQuestion("question.Tag", []models.NewQuestion{question})
	if err != nil {
		t.Fatalf("Error adding question: %s", err)
	}
}

func TestFindQuestionForMatch(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.QuestionService{}

	game := tests.TestGame()

	_, err := service.FindQuestionForMatch(game)
	if err != nil {
		t.Fatalf("Error finding question: %s", err)
	}

}
