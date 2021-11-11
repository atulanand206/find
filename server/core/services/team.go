package services

import (
	e "errors"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
)

type TeamService struct {
	Crud db.TeamCrud

	SubscriberService SubscriberService
}

func (service TeamService) CreateTeams(quiz models.Game) (teams []models.Team, err error) {
	teams = models.InitNewTeams(quiz)
	if err = service.Crud.CreateTeams(teams); err != nil {
		err = e.New(errors.Err_TeamNotCreated)
	}
	return
}

func (service TeamService) FindTeams(quiz models.Game) (teams []models.Team, err error) {
	return service.Crud.FindTeams(quiz.Id)
}

func (service TeamService) FindAndFillTeamVacancy(match models.Game, teams []models.Team, player models.Player) (teamId string, err error) {
	teamIds := service.Crud.FindTeamIdsFromTeams(teams)
	subscribers, err := service.SubscriberService.FindSubscribersForTags(teamIds)
	if err != nil {
		return
	}
	if len(subscribers) >= match.Specs.Players*match.Specs.Teams {
		err = e.New(errors.Err_PlayersFullInTeam)
		return
	}
	teamId = service.FindVacantTeamId(teams, subscribers, match.Specs.Players)
	_, err = service.SubscriberService.FindOrCreateSubscriber(teamId, player, actions.TEAM)
	if err != nil {
		return
	}
	return
}

func (service TeamService) FindVacantTeamId(teams []models.Team, subscribers []models.Subscriber, playersCount int) (teamId string) {
	mp := make(map[string]int)
	for _, v := range teams {
		mp[v.Id] = 0
	}
	for _, v := range subscribers {
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
