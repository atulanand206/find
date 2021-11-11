package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
)

func TestQuestionCrud(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.QuestionCrud{Db: db.NewDb()}

	t.Run("seed index and find index by tag", func(t *testing.T) {
		indexes := tests.TestIndexes(4)
		err := crud.SeedIndexes(indexes)
		assert.Nil(t, err, "error must be nil")

		index, err := crud.FindIndexByTag(indexes[0].Tag)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, indexes[0].Id, index.Id, "index id must be equal")

		index, err = crud.FindIndexByTag(indexes[1].Tag)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, indexes[1].Id, index.Id, "index id must be equal")
	})

	t.Run("fail find index by tag", func(t *testing.T) {
		_, err := crud.FindIndexByTag(tests.TestRandomString())
		assert.NotNil(t, err, "error must not be nil")
	})

	t.Run("find indexes", func(t *testing.T) {
		crud.Db.DropCollections()
		indexes := tests.TestIndexes(4)
		err := crud.SeedIndexes(indexes)
		assert.Nil(t, err, "error must be nil")

		foundIndexes, err := crud.FindIndexes()
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, len(indexes), len(foundIndexes), "indexes must be equal")
		for _, index := range indexes {
			assert.Contains(t, foundIndexes, index, "index must be found")
		}
	})

	t.Run("fail find indexes", func(t *testing.T) {
		crud.Db.DropCollections()
		indexes, err := crud.FindIndexes()
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 0, len(indexes), "indexes must be empty")
	})

	t.Run("seed question and find question", func(t *testing.T) {
		index := tests.TestIndex()
		err := crud.SeedIndexes([]models.Index{index})
		assert.Nil(t, err, "error must be nil")

		newQuestion := tests.TestNewQuestion()
		question := tests.TestQuestion(index.Tag, newQuestion)
		err = crud.SeedQuestions([]models.Question{question})
		assert.Nil(t, err, "error must be nil")

		foundQuestion, err := crud.FindQuestion(question.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, question.Id, foundQuestion.Id, "question id must be equal")
		assert.Equal(t, question.Tag, foundQuestion.Tag, "question index tag must be equal")
		assert.Equal(t, question.Statements, foundQuestion.Statements, "question statements must be equal")
	})

	t.Run("fail find question", func(t *testing.T) {
		_, err := crud.FindQuestion(tests.TestRandomString())
		assert.NotNil(t, err, "error must not be nil")
	})

	t.Run("seed answer and find answer", func(t *testing.T) {
		index := tests.TestIndex()
		err := crud.SeedIndexes([]models.Index{index})
		assert.Nil(t, err, "error must be nil")

		newQuestion := tests.TestNewQuestion()
		question := tests.TestQuestion(index.Tag, newQuestion)
		err = crud.SeedQuestions([]models.Question{question})
		assert.Nil(t, err, "error must be nil")

		answer := tests.TestAnswer(question.Id, newQuestion.Answer)
		err = crud.SeedAnswers([]models.Answer{answer})
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, newQuestion.Answer, answer.Answer, "answer must be equal")

		foundQuestion, err := crud.FindQuestion(question.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, question.Id, foundQuestion.Id, "question id must be equal")
		assert.Equal(t, question.Tag, foundQuestion.Tag, "question index tag must be equal")
		assert.Equal(t, question.Statements, foundQuestion.Statements, "question statements must be equal")

		foundAnswer, err := crud.FindAnswer(question.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, answer.Id, foundAnswer.Id, "answer id must be equal")
		assert.Equal(t, answer.Answer, foundAnswer.Answer, "answer must be equal")
		assert.NotEmpty(t, foundAnswer.Hint, "answer hint must be not empty")
	})

	t.Run("fail find answer", func(t *testing.T) {
		_, err := crud.FindAnswer(tests.TestRandomString())
		assert.NotNil(t, err, "error must not be nil")
	})

	t.Run("find questions from index", func(t *testing.T) {
		index := tests.TestIndex()
		err := crud.SeedIndexes([]models.Index{index})
		assert.Nil(t, err, "error must be nil")

		newQuestion := tests.TestNewQuestion()
		question := tests.TestQuestion(index.Tag, newQuestion)
		err = crud.SeedQuestions([]models.Question{question})
		assert.Nil(t, err, "error must be nil")

		answer := tests.TestAnswer(question.Id, question.Statements)
		err = crud.SeedAnswers([]models.Answer{answer})
		assert.Nil(t, err, "error must be nil")

		foundQuestions, err := crud.FindQuestionsFromIndex(index, 1)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 1, len(foundQuestions), "number of questions must be one")
	})

	t.Run("fail find questions from index", func(t *testing.T) {
		index := tests.TestIndex()
		foundQuestions, err := crud.FindQuestionsFromIndex(index, 1)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 0, len(foundQuestions), "number of questions must be 0")
	})

	t.Run("find questions from indexes", func(t *testing.T) {
		index := tests.TestIndex()
		err := crud.SeedIndexes([]models.Index{index})
		assert.Nil(t, err, "error must be nil")

		newQuestion := tests.TestNewQuestion()
		question := tests.TestQuestion(index.Tag, newQuestion)
		err = crud.SeedQuestions([]models.Question{question})
		assert.Nil(t, err, "error must be nil")

		answer := tests.TestAnswer(question.Id, question.Statements)
		err = crud.SeedAnswers([]models.Answer{answer})
		assert.Nil(t, err, "error must be nil")

		foundQuestions, err := crud.FindQuestionsFromIndexes([]models.Index{index}, 1)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 1, len(foundQuestions), "number of questions must be one")
	})

	t.Run("fail find questions from indexes", func(t *testing.T) {
		index := tests.TestIndex()
		foundQuestions, err := crud.FindQuestionsFromIndexes([]models.Index{index}, 1)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 0, len(foundQuestions), "number of questions must be 0")
	})
}
