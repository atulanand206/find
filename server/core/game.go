package core

import (
	"errors"
)

type Service struct {
	matchService      MatchService
	subscriberService SubscriberService
	playerService     PlayerService
	teamService       TeamService
	validator         Validator
}

func (service Service) GenerateActiveQuizResponse() (matches []Game, err error) {
	return service.matchService.FindActiveMatches()
}

func (service Service) GenerateBeginGameResponse(player Player) (res Player, err error) {
	return service.playerService.FindOrCreatePlayer(player)
}

func (service Service) GenerateCreateGameResponse(quizmaster Player, specs Specs) (response Snapshot, err error) {
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
	roster := TableRoster(teams, []Subscriber{}, []Player{})

	return service.subscriberService.subscribeAndRespond(quiz, roster, player, Snapshot{}, QUIZMASTER)
}

func (service Service) GenerateEnterGameResponse(request Request) (response Snapshot, err error) {
	player, err := service.playerService.FindPlayerByEmail(request.Person.Email)
	if err != nil {
		return
	}

	match, teams, teamPlayers, _, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if IsQuizMasterInMatch(match, player) {
		return service.subscriberService.subscribeAndRespond(match, roster, player, snapshot, QUIZMASTER)
	}

	if IsPlayerInTeams(teamPlayers, player) {
		return service.subscriberService.subscribeAndRespond(match, roster, player, snapshot, PLAYER)
	}

	_, err = service.teamService.FindAndFillTeamVacancy(match, teams, player)
	if err != nil {
		return
	}

	match, teams, teamPlayers, _, roster, snapshot, err = service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(match, roster, player, snapshot, PLAYER)
}

func (service Service) GenerateFullMatchResponse(quizId string) (response Snapshot, err error) {
	_, _, _, _, _, snapshot, err := service.matchService.FindMatchFull(quizId)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateWatchGameResponse(request Request) (response Snapshot, err error) {
	audience, err := service.playerService.FindPlayerByEmail(request.Person.Email)
	if err != nil {
		return
	}

	match, _, _, _, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	return service.subscriberService.subscribeAndRespond(match, roster, audience, snapshot, AUDIENCE)
}

func (service Service) GenerateStartGameResponse(request Request) (response Snapshot, err error) {
	match, teams, teamPlayers, _, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
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

	snapshot = InitialSnapshot(match.Id, question, roster, teams[0].Id)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateQuestionHintResponse(request Request) (response Snapshot, err error) {
	_, _, _, _, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	answer, err := Db.FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	snapshot = SnapshotWithHint(snapshot, answer.Hint, roster)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateQuestionAnswerResponse(request Request) (response Snapshot, err error) {
	match, teams, teamPlayers, players, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
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
	team.Score += ScoreAnswer(match.Specs.Points, snapshot.RoundNo)

	err = service.teamService.db.UpdateTeam(team)
	if err != nil {
		err = errors.New(Err_TeamNotUpdated)
	}

	teams, err = service.teamService.db.FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamNotPresent)
	}

	answer, err := Db.FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	roster = TableRoster(teams, teamPlayers, players)

	snapshot = SnapshotWithAnswer(snapshot, answer.Answer, match.Specs.Points, roster)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateNextQuestionResponse(request Request) (response Snapshot, err error) {

	match, teams, _, _, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
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

	teamsTurn := NextTeam(teams, request.TeamSTurn)
	snapshot = SnapshotWithNext(snapshot, roster, teamsTurn, question)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func (service Service) GeneratePassQuestionResponse(request Request) (response Snapshot, err error) {
	_, teams, _, _, roster, snapshot, err := service.matchService.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	teamsTurn := NextTeam(teams, request.TeamSTurn)
	snapshot = SnapshotWithPass(snapshot, roster, teamsTurn)
	if err = Db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateScoreResponse(request Request) (response ScoreResponse, err error) {
	snapshots, err := Db.FindSnapshots(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
		return
	}

	response = InitScoreResponse(request.QuizId, snapshots)
	return
}
