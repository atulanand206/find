package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
)

func TestTableRoster(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	playerService := core.PlayerService{}
	matchService := core.MatchService{}
	teamService := core.TeamService{}
	subscriberService := core.SubscriberService{}

	quizmaster, _ := playerService.FindOrCreatePlayer(testPlayer())
	specs := testSpecs()
	game, _ := matchService.CreateMatch(quizmaster, specs)
	teams, _ := teamService.CreateTeams(game)
	for _, team := range teams {
		player, _ := playerService.FindOrCreatePlayer(testPlayer())
		subscriberService.FindOrCreateSubscriber(team.Id, player, core.TEAM)
	}

	subscribers, _ := subscriberService.FindTeamPlayers(teams)
	players, _ := playerService.FindPlayers(subscribers)

	roster := core.TableRoster(teams, subscribers, players)

	if len(roster) != len(teams) {
		t.Errorf("Expected %d, got %d", len(teams), len(roster))
	}

	for ix, team := range teams {
		if roster[ix].Score != team.Score {
			t.Errorf("Expected %d, got %d", team.Score, roster[ix].Score)
		}
	}

	for _, team := range roster {
		if len(team.Players) != 1 {
			t.Errorf("Expected 1, got %d", len(team.Players))
		}
	}
}
