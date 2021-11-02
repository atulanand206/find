package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
)

func TestCreateAndFindTeams(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.TeamService{}

	game := testGame()
	service.CreateTeams(game)

	res, err := service.FindTeams(game)
	if err != nil {
		t.Fatalf("teams for quiz %s not found", game.Id)
	}
	if len(res) != game.Specs.Teams {
		t.Fatalf("teams count mismatch")
	}
}
