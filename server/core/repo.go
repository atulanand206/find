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

	subscriberService SubscriberService
}

func (service SubscriberService) subscribeAndRespond(match Game, teams []Team, teamPlayers []Subscriber,
	players []Player, playerTeamId string, player Player, snapshot Snapshot, role Role) (response EnterGameResponse, err error) {
	_, err = service.FindOrCreateSubscriber(match.Id, player, role)
	if err != nil {
		return
	}

	response = InstanceCreator.InitEnterGameResponse(match, teams, teamPlayers, players, "", snapshot)
	return
}

func (service SubscriberService) FindOrCreateSubscriber(tag string, audience Player, role Role) (subscriber Subscriber, err error) {
	subscriber, err = service.db.FindSubscriberForTagAndPlayerId(tag, audience.Id)
	if err != nil {
		subscriber = InstanceCreator.InitSubscriber(tag, audience, role.String())
		err = service.db.CreateSubscriber(subscriber)
		if err != nil {
			err = errors.New(fmt.Sprint(ErrorCreator.SubscriberNotCreated(subscriber)))
		}
	}
	return
}

func (service SubscriberService) FindSubscribersForTag(tag string, role Role) (subscribers []Subscriber, err error) {
	return service.db.FindSubscribers(tag, role)
}

func (service MatchService) FindMatchFull(matchId string) (match Game, teams []Team, teamPlayers []Subscriber, players []Player, snapshot Snapshot, err error) {
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

func (service PlayerService) DeletePlayerLiveSession(playerId string) (res WebsocketMessage, targets map[string]bool, err error) {
	subscribers, err := service.db.FindSubscriptionsForPlayerId(playerId)
	tags := make([]string, 0)
	for _, subscriber := range subscribers {
		tags = append(tags, subscriber.Tag)
	}

	subscribers, err = service.db.FindSubscribersForTag(tags)
	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		if playerId != subscriber.PlayerId {
			targets[subscriber.PlayerId] = true
		}
	}

	res = MessageCreator.InitWebSocketMessage(S_REFRESH, "Player dropped. Please refresh.")
	err = service.db.DeleteSubscriber(playerId)
	return
}

func (service TeamService) CreateTeams(quiz Game) (teams []Team, err error) {
	teams = InitNewTeams(quiz)
	if err = service.db.CreateTeams(teams); err != nil {
		err = errors.New(Err_TeamNotCreated)
	}
	return
}

func (service TeamService) TeamIdForPlayer(teamPlayers []Subscriber, player Player) (teamId string) {
	for _, v := range teamPlayers {
		if v.PlayerId == player.Id {
			teamId = v.Tag
			return
		}
	}
	return
}

func (service TeamService) FindAndFillTeamVacancy(match Game, teams []Team, player Player) (teamId string, err error) {
	teamPlayers, err := service.subscriberService.FindSubscribersForTag(match.Id, TEAM)
	if err != nil {
		return
	}
	if len(teamPlayers) >= match.Specs.Players*match.Specs.Teams {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}
	teamId = FindVacantTeamId(teams, match.Specs.Players)
	_, err = service.subscriberService.FindOrCreateSubscriber(teamId, player, TEAM)
	if err != nil {
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
