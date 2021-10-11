package core

import (
	"errors"
	"fmt"
)

type Repository struct {
	db DB
}

type MatchService struct {
	db DB
}

type SubscriberService struct {
	db DB
}

type PlayerService struct {
	db DB
}

type TeamService struct {
	db DB
}

func (service SubscriberService) subscribeAndRespond(match Game, teams []Team, teamPlayers []TeamPlayer,
	players []Player, playerTeamId string, player Player, snapshot Snapshot, role Role) (response EnterGameResponse, err error) {
	_, err = service.FindOrCreateSubscriber(match, player, role)
	if err != nil {
		return
	}

	response = InstanceCreator.InitEnterGameResponse(match, teams, teamPlayers, players, "", snapshot)
	return
}

func (service SubscriberService) FindOrCreateSubscriber(match Game, audience Player, role Role) (subscriber Subscriber, err error) {
	subscriber, err = service.db.FindSubscriberForTagAndPlayerId(match.Id, audience.Id)
	if err != nil {
		subscriber = InstanceCreator.InitSubscriber(match, audience, role.String())
		err = service.db.CreateSubscriber(subscriber)
		if err != nil {
			err = errors.New(fmt.Sprint(ErrorCreator.SubscriberNotCreated(subscriber)))
		}
	}
	return
}

func (service MatchService) FindMatchFull(matchId string) (match Game, teams []Team, teamPlayers []TeamPlayer, players []Player, snapshot Snapshot, err error) {
	match, err = service.db.FindMatch(matchId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err = service.db.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
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

func (service PlayerService) FindOrCreatePlayer(request Player) (player Player, err error) {
	player, err = service.db.FindPlayer(request.Email)
	if err != nil {
		player = InitNewPlayer(request)
		if err = service.db.CreatePlayer(player); err != nil {
			err = errors.New(Err_PlayerNotCreated)
			return
		}
	}
	return
}

func (service PlayerService) FindPlayerByEmail(email string) (player Player, err error) {
	player, err = service.db.FindPlayer(email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	return
}

func (service PlayerService) DeletePlayerLiveSession(playerId string) error {
	return service.db.DeleteSubscriber(playerId)
}

func (service TeamService) CreateTeams(quiz Game) (teams []Team, err error) {
	teams = InitNewTeams(quiz)
	if err = service.db.CreateTeams(teams); err != nil {
		err = errors.New(Err_TeamNotCreated)
	}
	return
}

func (service TeamService) TeamIdForPlayer(teamPlayers []TeamPlayer, player Player) (teamId string) {
	for _, v := range teamPlayers {
		if v.PlayerId == player.Id {
			teamId = v.TeamId
			return
		}
	}
	return
}

func (service TeamService) FindAndFillTeamVacancy(match Game, teams []Team, teamPlayers []TeamPlayer, player Player) (teamId string, err error) {
	if len(teamPlayers) >= match.Specs.Players*match.Specs.Teams {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}
	mp := make(map[string]int)
	for _, v := range teamPlayers {
		if mp[v.TeamId] == 0 {
			mp[v.TeamId] = 1
		} else {
			mp[v.TeamId] = mp[v.TeamId] + 1
		}
	}
	var x = match.Specs.Players
	for _, v := range mp {
		if v < x {
			x = v
			return
		}
	}
	for k, v := range mp {
		if v == x {
			teamId = k
			return
		}
	}

	if err = service.db.CreateTeamPlayer(InitTeamPlayer(teamId, player)); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}
	return
}

func (repo Repository) FindQuestionForMatch(match Game) (question Question, err error) {
	index, err := repo.db.FindIndex()
	if err != nil {
		err = errors.New(Err_IndexNotPresent)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), 1)
	questions, err := repo.db.FindQuestionsFromIndexes(indexes, int64(1))
	if len(questions) != 1 || err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	question = questions[0]
	return
}
