package services

import (
	e "errors"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/utils"
)

type MatchService struct {
	Crud db.MatchCrud

	SubscriberService SubscriberService
}

func (service Service) FindMatchFull(matchId string) (
	match models.Game, teams []models.Team,
	subscribers []models.Subscriber, players []models.Player,
	roster []models.TeamRoster,
	snapshot models.Snapshot, err error) {
	match, err = service.MatchService.Crud.FindMatch(matchId)
	if err != nil {
		err = e.New(errors.Err_MatchNotPresent)
		return
	}

	teams, err = service.TeamService.FindTeams(match.Id)
	if err != nil {
		err = e.New(errors.Err_TeamsNotPresentInMatch)
		return
	}

	teamIds := service.TeamService.FindTeamIdsFromTeams(teams)
	subscribers, err = service.SubscriberService.FindSubscribersForTags(teamIds)
	if err != nil {
		err = e.New(errors.Err_SubscribersNotPresentInMatch)
		return
	}

	playerIds := service.SubscriberService.FindPlayerIdsFromSubscribers(subscribers)
	players, err = service.PlayerService.FindPlayers(playerIds)
	if err != nil {
		err = e.New(errors.Err_PlayerNotPresent)
		return
	}

	roster = utils.TableRoster(teams, subscribers, players)
	snapshot, err = service.SnapshotService.Crud.FindLatestSnapshot(match.Id)
	if err != nil {
		err = e.New(errors.Err_SnapshotNotPresent)
		return
	}
	return
}

func (service MatchService) CreateMatch(player models.Player, specs models.Specs) (quiz models.Game, err error) {
	quiz = models.InitNewMatch(player, specs)
	if err = service.Crud.CreateMatch(quiz); err != nil {
		err = e.New(errors.Err_MatchNotCreated)
	}
	return
}

func (service MatchService) FindActiveMatchesForPlayer(playerId string) (matches []models.Game, err error) {
	matches, err = service.Crud.FindActiveMatches()
	for ix, match := range matches {
		if match.QuizMaster.Id == playerId {
			match.CanJoin = true
			matches[ix] = match
		}
		subscribers, err := service.SubscriberService.FindSubscribers(match.Id, actions.PLAYER)
		if err == nil {
			match.PlayersJoined = len(subscribers)
			if !match.CanJoin {
				for _, subscriber := range subscribers {
					if subscriber.PlayerId == playerId {
						match.CanJoin = true
						matches[ix] = match
						break
					}
				}
				if !match.CanJoin {
					if match.Specs.Players*match.Specs.Teams > len(subscribers) {
						match.CanJoin = true
						matches[ix] = match
					}
				}
			}
		}
	}
	return
}

func (service MatchService) UpdateMatchTags(match models.Game, tag string) (bool, error) {
	match.Tags = append(match.Tags, tag)
	return service.Crud.UpdateMatch(match)
}
