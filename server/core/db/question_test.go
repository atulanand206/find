package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
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

	t.Run("seed question and find question", func(t *testing.T) {

	})
}
