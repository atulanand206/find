package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

func TestMatchCrud(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)
	crud := db.MatchCrud{Db: db.NewDb()}
	crud.Db.DropCollections()

	t.Run("create and find match", func(t *testing.T) {
		game := tests.TestGame()
		err := crud.CreateMatch(game)
		assert.Nil(t, err)
		res, err := crud.FindMatch(game.Id)
		if err != nil {
			t.Fatalf("match %s not found", game.Id)
		}
		if res.Id != game.Id || res.Specs.Name != game.Specs.Name {
			t.Fatalf("match id mismatch")
		}
	})

	t.Run("find match fail", func(t *testing.T) {
		id, _ := gonanoid.New(8)
		_, err := crud.FindMatch(id)
		if err == nil {
			t.Fatalf("match found")
		}
	})

	t.Run("find active matches", func(t *testing.T) {
		game := tests.TestGame()
		crud.CreateMatch(game)

		res, err := crud.FindActiveMatches()
		assert.Nil(t, err, "error must be nil")
		for _, match := range res {
			if match.Id == game.Id {
				return
			}
		}
		t.Fatalf("match id not present")
	})

	t.Run("update match fail", func(t *testing.T) {
		game := tests.TestGame()
		updated, _ := crud.UpdateMatch(game)
		assert.False(t, updated, "match must not be updated")
	})

	t.Run("update match", func(t *testing.T) {
		game := tests.TestGame()
		crud.CreateMatch(game)

		res, err := crud.FindMatch(game.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, game.Specs.Name, res.Specs.Name, "quiz name must be equal")

		game.Specs.Name, _ = gonanoid.New(8)
		crud.UpdateMatch(game)

		res, err = crud.FindMatch(game.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, game.Specs.Name, res.Specs.Name, "quiz name must be equal")
	})
}
