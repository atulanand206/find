package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
)

func TestCreateAndFindTeams(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	Db := db.DB{}
	subscriberService := services.SubscriberService{Crud: db.SubscriberCrud{Db: Db}, TargetService: services.TargetService{}, Creators: services.Creators{}}
	service := services.TeamService{Crud: db.TeamCrud{Db: Db}, SubscriberService: subscriberService}

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
