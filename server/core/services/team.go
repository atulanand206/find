package services

import (
	e "errors"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
)

type TeamService struct {
	crud db.TeamCrud

	subscriberService SubscriberService
}

func (service TeamService) CreateTeams(quiz models.Game) (teams []models.Team, err error) {
	teams = models.InitNewTeams(quiz)
	if err = service.crud.CreateTeams(teams); err != nil {
		err = e.New(errors.Err_TeamNotCreated)
	}
	return
}

func (service TeamService) FindTeams(quiz models.Game) (teams []models.Team, err error) {
	return service.crud.FindTeams(quiz)
}

func (service TeamService) FindAndFillTeamVacancy(match models.Game, teams []models.Team, player models.Player) (teamId string, err error) {
	teamIds := []string{}
	for _, team := range teams {
		teamIds = append(teamIds, team.Id)
	}
	teamPlayers, err := service.subscriberService.FindSubscribersForTag(teamIds)
	if err != nil {
		return
	}
	if len(teamPlayers) >= match.Specs.Players*match.Specs.Teams {
		err = e.New(errors.Err_PlayersFullInTeam)
		return
	}
	teamId = service.FindVacantTeamId(teams, teamPlayers, match.Specs.Players)
	_, err = service.subscriberService.FindOrCreateSubscriber(teamId, player, actions.TEAM)
	if err != nil {
		return
	}
	return
}

func (service TeamService) FindVacantTeamId(teams []models.Team, teamPlayers []models.Subscriber, playersCount int) (teamId string) {
	mp := make(map[string]int)
	for _, v := range teams {
		mp[v.Id] = 0
	}
	for _, v := range teamPlayers {
		mp[v.Tag] = mp[v.Tag] + 1
	}
	var x = playersCount
	for _, v := range mp {
		if v < x {
			x = v
		}
	}
	for k, v := range mp {
		if v == x {
			teamId = k
			return
		}
	}
	return
}
