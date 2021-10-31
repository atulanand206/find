package core

import (
	"errors"
)

type MatchService struct {
	crud MatchCrud
}

func (service Service) FindMatchFull(matchId string) (
	match Game, teams []Team,
	teamPlayers []Subscriber, players []Player,
	roster []TeamRoster,
	snapshot Snapshot, err error) {
	match, err = service.matchService.crud.FindMatch(matchId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err = service.teamService.crud.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	teamPlayers, err = service.subscriberService.crud.FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	players, err = service.playerService.crud.FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	roster = TableRoster(teams, teamPlayers, players)
	snapshot, err = service.snapshotService.crud.FindLatestSnapshot(match.Id)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
		return
	}
	return
}

func (service MatchService) CreateMatch(player Player, specs Specs) (quiz Game, err error) {
	quiz = InitNewMatch(player, specs)
	if err = service.crud.CreateMatch(quiz); err != nil {
		err = errors.New(Err_MatchNotCreated)
	}
	return
}

func (service MatchService) FindActiveMatches() (matches []Game, err error) {
	matches, err = service.crud.FindActiveMatches()
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
	}
	return
}
