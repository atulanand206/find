package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestCreateAndFindTeams(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.TeamService{}

	game := tests.TestGame()
	service.CreateTeams(game)

	res, err := service.FindTeams(game)
	if err != nil {
		t.Fatalf("teams for quiz %s not found", game.Id)
	}
	if len(res) != game.Specs.Teams {
		t.Fatalf("teams count mismatch")
	}
}
