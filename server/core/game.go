package core

import (
	"errors"
	"math"
)

type Service struct {
	matchService      MatchService
	subscriberService SubscriberService
	playerService     PlayerService
	teamService       TeamService
}

func (service Service) GenerateActiveQuizResponse() (matches []Game, err error) {
	return service.matchService.FindActiveMatches()
}

func (service Service) GenerateBeginGameResponse(player Player) (res Player, err error) {
	return service.playerService.FindOrCreatePlayer(player)
}

func (service Service) GenerateCreateGameResponse(request CreateGameRequest) (response EnterGameResponse, err error) {
	player := request.Quizmaster
	player, err = service.playerService.FindOrCreatePlayer(player)
	if err != nil {
		return
	}

	quiz, err := service.matchService.CreateMatch(player, request.Specs)
	if err != nil {
		return
	}

	teams, err := service.teamService.CreateTeams(quiz)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(quiz, teams, []Subscriber{}, []Player{}, "", player, Snapshot{}, QUIZMASTER)
}

func (service Service) GenerateEnterGameResponse(enterGameRequest EnterGameRequest) (response EnterGameResponse, err error) {
	player, err := service.playerService.FindPlayerByEmail(enterGameRequest.Person.Email)
	if err != nil {
		return
	}

	match, teams, teamPlayers, players, snapshot, err := service.matchService.FindMatchFull(enterGameRequest.QuizId)
	if err != nil {
		return
	}

	if IsQuizMasterInMatch(match, player) {
		return service.subscriberService.subscribeAndRespond(match, teams, teamPlayers, players, "", player, snapshot, QUIZMASTER)
	}

	if IsPlayerInTeams(teamPlayers, player) {
		return service.subscriberService.subscribeAndRespond(match, teams, teamPlayers, players, service.teamService.TeamIdForPlayer(teamPlayers, player), player, snapshot, PLAYER)
	}

	_, err = service.teamService.FindAndFillTeamVacancy(match, teams, player)
	if err != nil {
		return
	}

	match, teams, teamPlayers, players, snapshot, err = service.matchService.FindMatchFull(enterGameRequest.QuizId)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(match, teams, teamPlayers, players, service.teamService.TeamIdForPlayer(teamPlayers, player), player, snapshot, PLAYER)
}

func (service Service) GenerateWatchGameResponse(enterGameRequest EnterGameRequest) (response EnterGameResponse, err error) {
	audience, err := service.playerService.FindPlayerByEmail(enterGameRequest.Person.Email)
	if err != nil {
		return
	}

	match, teams, teamPlayers, players, snapshot, err := service.matchService.FindMatchFull(enterGameRequest.QuizId)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(match, teams, teamPlayers, players, "", audience, snapshot, AUDIENCE)
}

func GenerateStartGameResponse(startGameRequest StartGameRequest) (response StartGameResponse, err error) {
	match, err := Db.FindMatch(startGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err := Db.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamPlayers, err := Db.FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	if result := MatchFull(match, teamPlayers); !result {
		err = errors.New(Err_WaitingForPlayers)
		return
	}

	question, err := Repo.FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if err = Db.UpdateMatchQuestions(match, question); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	snapshot := InitSnapshotDtoF(match.Id, question.Id, teams[0].Id,
		START.String(), 0, 1, 1, question.Statements)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	players, err := Db.FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	response = InitStartGameResponse(match.Id, teams, teamPlayers, players, question, snapshot)
	return
}

func GenerateQuestionHintResponse(request GameSnapRequest) (response Snapshot, err error) {
	_, err = Db.FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	answer, err := Db.FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	snapshot, err := Db.FindLatestSnapshot(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	snapshot = InitSnapshotDtoF(request.QuizId, request.QuestionId, snapshot.TeamSTurn,
		HINT.String(), 0,
		snapshot.QuestionNo, snapshot.RoundNo, answer.Answer)

	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GenerateQuestionAnswerResponse(request GameSnapRequest) (response Snapshot, err error) {
	match, err := Db.FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	answer, err := Db.FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	snapshot, err := Db.FindLatestSnapshot(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	snapshot = InitSnapshotDtoF(request.QuizId, request.QuestionId, request.TeamSTurn,
		RIGHT.String(), match.Specs.Points/int(math.Pow(2, float64(snapshot.RoundNo))),
		snapshot.QuestionNo, snapshot.RoundNo, answer.Answer)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GenerateNextQuestionResponse(request GameSnapRequest) (response Snapshot, err error) {
	match, err := Db.FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	if !QuestionCanBeAdded(match) {
		err = errors.New(Err_QuestionsNotLeft)
		return
	}

	question, err := Repo.FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if err = Db.UpdateMatchQuestions(match, question); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	snapshot, err := Db.FindLatestSnapshot(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	teams, err := Db.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamsTurn := NextTeam(teams, request.TeamSTurn)
	snapshot = InitSnapshotDtoF(request.QuestionId, question.Id, teamsTurn, NEXT.String(), 0, snapshot.QuestionNo+1, 1, question.Statements)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GeneratePassQuestionResponse(request GameSnapRequest) (response Snapshot, err error) {
	match, err := Db.FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	snapshots, err := Db.FindQuestionSnapshots(request.QuizId, request.QuestionId)
	if len(snapshots) == 0 || err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	teams, err := Db.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamsTurn := NextTeam(teams, request.TeamSTurn)
	snapshot := snapshots[len(snapshots)-1]
	snapshot = InitSnapshotDtoF(match.Id, request.QuestionId, teamsTurn, PASS.String(), 0, snapshot.QuestionNo, snapshot.RoundNo, []string{})
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GenerateScoreResponse(request ScoreRequest) (response ScoreResponse, err error) {
	snapshots, err := Db.FindSnapshots(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
		return
	}

	response = InitScoreResponse(request, snapshots)
	return
}
