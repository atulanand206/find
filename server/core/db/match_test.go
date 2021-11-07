package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TestCreateAndFindMatch(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.MatchCrud{}

	game := tests.TestGame()
	crud.CreateMatch(game)

	res, err := crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Id != game.Id || res.Specs.Name != game.Specs.Name {
		t.Fatalf("match id mismatch")
	}
}

func TestFindMatchFail(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.MatchCrud{}

	id, _ := gonanoid.New(8)
	_, err := crud.FindMatch(id)
	if err == nil {
		t.Fatalf("match found")
	}
}

func TestFindActiveMatches(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.MatchCrud{}

	game := tests.TestGame()
	crud.CreateMatch(game)

	res, err := crud.FindActiveMatches()
	if err != nil {
		t.Fatalf("matches %s not found", game.Id)
	}
	for _, match := range res {
		if match.Id == game.Id {
			return
		}
	}
	t.Fatalf("match id not present")
}

func TestUpdateMatchFail(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.MatchCrud{}

	game := tests.TestGame()
	count, _ := crud.UpdateMatch(game)
	if count {
		t.Fatalf("test failed")
	}
}

func TestUpdateMatch(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.MatchCrud{}

	game := tests.TestGame()
	crud.CreateMatch(game)

	res, err := crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Specs.Name != game.Specs.Name {
		t.Fatalf("specs not correct")
	}

	game.Specs.Name, _ = gonanoid.New(8)
	crud.UpdateMatch(game)

	res, err = crud.FindMatch(game.Id)
	if err != nil {
		t.Fatalf("match %s not found", game.Id)
	}
	if res.Specs.Name != game.Specs.Name {
		t.Fatalf("specs not updated")
	}
}