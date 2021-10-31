package core

import "errors"

type TeamService struct {
	crud TeamCrud

	subscriberService SubscriberService
}

func (service TeamService) CreateTeams(quiz Game) (teams []Team, err error) {
	teams = InitNewTeams(quiz)
	if err = service.crud.CreateTeams(teams); err != nil {
		err = errors.New(Err_TeamNotCreated)
	}
	return
}

func (service TeamService) FindAndFillTeamVacancy(match Game, teams []Team, player Player) (teamId string, err error) {
	teamIds := []string{}
	for _, team := range teams {
		teamIds = append(teamIds, team.Id)
	}
	teamPlayers, err := service.subscriberService.FindSubscribersForTag(teamIds)
	if err != nil {
		return
	}
	if len(teamPlayers) >= match.Specs.Players*match.Specs.Teams {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}
	teamId = service.FindVacantTeamId(teams, teamPlayers, match.Specs.Players)
	_, err = service.subscriberService.FindOrCreateSubscriber(teamId, player, TEAM)
	if err != nil {
		return
	}
	return
}
