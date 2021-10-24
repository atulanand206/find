package core

import (
	"errors"
)

type MatchService struct {
	db DB
}

func (service MatchService) FindMatchFull(matchId string) (
	match Game, teams []Team,
	teamPlayers []Subscriber, players []Player,
	roster []TeamRoster,
	snapshot Snapshot, err error) {
	match, err = service.db.FindMatch(matchId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err = service.db.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	teamPlayers, err = service.db.FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	players, err = service.db.FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	roster = TableRoster(teams, teamPlayers, players)
	if match.Active {
		snapshot, err = service.db.FindLatestSnapshot(match.Id)
		if err != nil {
			err = errors.New(Err_SnapshotNotPresent)
			return
		}
	}
	return
}

func (service MatchService) CreateMatch(player Player, specs Specs) (quiz Game, err error) {
	quiz = InitNewMatch(player, specs)
	if err = service.db.CreateMatch(quiz); err != nil {
		err = errors.New(Err_MatchNotCreated)
	}
	return
}

func (service MatchService) FindActiveMatches() (matches []Game, err error) {
	matches, err = service.db.FindActiveMatches()
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
	}
	return
}
