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
	teams := core.InitNewTeams(game)
	service.CreateTeams(game)

	res, err := service.FindTeams(game)
	if err != nil {
		t.Fatalf("teams for quiz %s not found", game.Id)
	}
	if len(res) != len(teams) {
		t.Fatalf("teams count mismatch")
	}
}
