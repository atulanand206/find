package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
)

func TestAddQuestion(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.QuestionService{}

	// Create a new question
	question := core.NewQuestion{
		Statements: []string{"test", "question"},
		Answer:     []string{"Test Answer"},
	}

	err := service.AddQuestion("question.Tag", []core.NewQuestion{question})
	if err != nil {
		t.Fatalf("Error adding question: %s", err)
	}
}

func TestFindQuestionForMatch(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.QuestionService{}

	game := testGame()

	_, err := service.FindQuestionForMatch(game)
	if err != nil {
		t.Fatalf("Error finding question: %s", err)
	}

}
