package core

import (
	"errors"
	"fmt"
)

type Service struct {
	authService       AuthService
	permissionService PermissionService
	matchService      MatchService
	subscriberService SubscriberService
	playerService     PlayerService
	teamService       TeamService
	questionService   QuestionService
	snapshotService   SnapshotService
	validator         Validator
}

func (service Service) GenerateGameResponse(request Request) (response Snapshot, err error) {

	return
}

func (service Service) GenerateBeginGameResponse(player Player) (res LoginResponse, err error) {
	response, err := service.playerService.FindOrCreatePlayer(player)
	if err != nil {
		return
	}

	tokens, err := service.authService.GenerateTokens(response)
	if err != nil {
		return
	}

	quizmaster := service.permissionService.HasPermissions(player.Id)
	res = LoginResponse{response, tokens, quizmaster}
	return
}

func (service Service) GenerateCreateGameResponse(quizmaster Player, specs Specs) (response GameResponse, err error) {
	player := quizmaster
	player, err = service.playerService.FindOrCreatePlayer(player)
	if err != nil {
		return
	}

	quiz, err := service.matchService.CreateMatch(player, specs)
	if err != nil {
		return
	}

	teams, err := service.teamService.CreateTeams(quiz)
	if err != nil {
		return
	}

	snapshot, err := service.snapshotService.InitialSnapshot(quiz.Id, teams)
	if err != nil {
		return
	}
	fmt.Println(snapshot)

	return service.subscriberService.subscribeAndRespond(quiz, player, snapshot, QUIZMASTER)
}

func (service Service) GenerateEnterGameResponse(request Request) (response GameResponse, err error) {
	player, err := service.playerService.FindPlayerByEmail(request.Person.Email)
	if err != nil {
		return
	}

	match, teams, teamPlayers, _, _, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if IsQuizMasterInMatch(match, player) {
		match, _, _, _, roster, snapshot, er := service.FindMatchFull(request.QuizId)
		if er != nil {
			err = er
			return
		}

		snapshot, er = service.snapshotService.SnapshotJoin(snapshot, roster)
		if er != nil {
			err = er
			return
		}

		return service.subscriberService.subscribeAndRespond(match, player, snapshot, QUIZMASTER)
	}

	if IsPlayerInTeams(teamPlayers, player) {
		match, _, _, _, roster, snapshot, er := service.FindMatchFull(request.QuizId)
		if er != nil {
			err = er
			return
		}

		snapshot, er = service.snapshotService.SnapshotJoin(snapshot, roster)
		if er != nil {
			err = er
			return
		}
		return service.subscriberService.subscribeAndRespond(match, player, snapshot, PLAYER)
	}

	_, err = service.teamService.FindAndFillTeamVacancy(match, teams, player)
	if err != nil {
		return
	}

	match, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	snapshot, err = service.snapshotService.SnapshotJoin(snapshot, roster)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(match, player, snapshot, PLAYER)
}

func (service Service) GenerateFullMatchResponse(quizId string) (response Snapshot, err error) {
	_, _, _, _, _, snapshot, err := service.FindMatchFull(quizId)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateWatchGameResponse(request Request) (response GameResponse, err error) {
	audience, err := service.playerService.FindPlayerByEmail(request.Person.Email)
	if err != nil {
		return
	}

	match, _, _, _, _, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(match, audience, snapshot, AUDIENCE)
}

func (service Service) GenerateStartGameResponse(request Request) (response Snapshot, err error) {
	match, teams, teamPlayers, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if result := MatchFull(match, teamPlayers); !result {
		err = errors.New(Err_WaitingForPlayers)
		return
	}

	question, err := service.questionService.FindQuestionForMatch(match)
	if err != nil {
		return
	}

	match.Started = true
	if err = service.matchService.crud.UpdateMatchQuestions(match, question); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	snapshot, err = service.snapshotService.SnapshotStart(snapshot, roster, question, teams[0].Id)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateQuestionHintResponse(request Request) (response Snapshot, err error) {
	_, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	answer, err := service.questionService.crud.FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	snapshot, err = service.snapshotService.SnapshotHint(snapshot, roster, answer.Hint)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateQuestionAnswerResponse(request Request) (response Snapshot, err error) {
	match, teams, _, _, _, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	teamId := snapshot.TeamSTurn
	var team Team
	for _, tm := range teams {
		if tm.Id == teamId {
			team = tm
		}
	}
	points := match.Specs.Points
	team.Score += ScoreAnswer(points, snapshot.RoundNo)

	err = service.teamService.crud.UpdateTeam(team)
	if err != nil {
		err = errors.New(Err_TeamNotUpdated)
		return
	}

	answer, err := service.questionService.crud.FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	_, _, _, _, roster, _, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	snapshot, err = service.snapshotService.SnapshotAnswer(snapshot, roster, answer, points)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateNextQuestionResponse(request Request) (response Snapshot, err error) {

	match, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if !QuestionCanBeAdded(match) {
		err = errors.New(Err_QuestionsNotLeft)
		return
	}

	question, err := service.questionService.FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if err = service.matchService.crud.UpdateMatchQuestions(match, question); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	snapshot, err = service.snapshotService.SnapshotNext(snapshot, roster, question, snapshot.TeamSTurn)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GeneratePassQuestionResponse(request Request) (response Snapshot, err error) {
	match, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	snapshot, err = service.snapshotService.SnapshotPass(snapshot, roster, snapshot.TeamSTurn, snapshot.RoundNo, match.Specs)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateScoreResponse(request Request) (response ScoreResponse, err error) {
	snapshots, err := service.snapshotService.crud.FindSnapshotsForMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
		return
	}

	response = InitScoreResponse(request.QuizId, snapshots)
	return
}

func (service Service) DeletePlayerLiveSession(playerId string) (res WebsocketMessage, targets map[string]bool, err error) {
	subscribers, err := service.subscriberService.crud.FindSubscriptionsForPlayerId(playerId)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	tags := make([]string, 0)
	for _, subscriber := range subscribers {
		tags = append(tags, subscriber.Tag)
	}

	subscribers, err = service.subscriberService.crud.FindSubscribersForTag(tags)
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
	err = service.subscriberService.crud.DeleteSubscriber(playerId)
	return
}
