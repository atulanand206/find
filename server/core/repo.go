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

	target Target
}

type PlayerService struct {
	db DB
}

type TeamService struct {
	db DB

	subscriberService SubscriberService
}

func (service SubscriberService) selfResponse(quizId string, action Action, response interface{}) (res WebsocketMessage, targets map[string]bool) {
	res = MessageCreator.WebSocketsResponse(action, response)
	targets = Controller.subscriberService.target.TargetSelf(quizId)
	return
}

func (service SubscriberService) joinResponse(quizId string, response GameResponse) (res WebsocketMessage, targets map[string]bool) {
	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = Controller.subscriberService.target.TargetQuiz(quizId, subscribers)
	res = MessageCreator.WebSocketsResponse(S_GAME, response)
	return
}

func (service SubscriberService) quizResponse(quizId string, response Snapshot) (res WebsocketMessage, targets map[string]bool) {
	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = Controller.subscriberService.target.TargetQuiz(quizId, subscribers)
	res = MessageCreator.WebSocketsResponse(S_GAME, response)
	return
}

func (service SubscriberService) subscribeAndRespond(match Game, roster []TeamRoster, player Player, snapshot Snapshot, role Role) (response GameResponse, err error) {
	_, err = service.FindOrCreateSubscriber(match.Id, player, role)
	if err != nil {
		return
	}

	response = GameResponse{Quiz: match, Snapshot: snapshot}
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

func (service SubscriberService) FindSubscribersForTag(tags []string) (subscribers []Subscriber, err error) {
	return service.db.FindSubscribersForTag(tags)
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
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	tags := make([]string, 0)
	for _, subscriber := range subscribers {
		tags = append(tags, subscriber.Tag)
	}

	subscribers, err = service.db.FindSubscribersForTag(tags)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

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
