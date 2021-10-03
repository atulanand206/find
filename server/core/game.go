package core

import (
	"errors"
	"fmt"
	"math"
)

func GenerateBeginGameResponse(player Player) (res Player, err error) {
	res, err = FindOrCreatePlayer(player)
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	return
}

func GenerateCreateGameResponse(request CreateGameRequest) (response EnterGameResponse, err error) {
	player := request.Quizmaster
	player, err = FindOrCreatePlayer(player)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	quiz := InitNewMatch(player, request.Specs)
	if err = CreateMatch(quiz); err != nil {
		err = errors.New(Err_MatchNotCreated)
	}

	if err = CreateTeams(quiz); err != nil {
		err = errors.New(Err_TeamNotCreated)
	}

	teams, err := FindTeams(quiz)
	if err != nil {
		err = errors.New(Err_TeamNotPresent)
	}

	response = InitEnterGameResponse(quiz, teams)
	fmt.Println(response)
	return
}

func GenerateEnterGameResponse(enterGameRequest EnterGameRequest) (response EnterGameResponse, err error) {
	match, err := FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	team, err := FindTeamVacancy(match, teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	player := enterGameRequest.Person
	if IsQuizMasterInMatch(match, player) {
		err = errors.New(Err_QuizmasterCantPlay)
		return
	}

	if IsPlayerInTeam(team, player) {
		err = errors.New(Err_PlayerAlreadyInGame)
		return
	}

	fmt.Println(match)
	if !PlayerCanBeAdded(match, team) {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}

	player, err = FindPlayer(player.Email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	if err = UpdatePlayerInTeam(team, player); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	match, err = FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err = FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	response = InitEnterGameResponse(match, teams)
	return
}

func GenerateWatchGameResponse(enterGameRequest EnterGameRequest) (response EnterGameResponse, err error) {
	match, err := FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	response = InitEnterGameResponse(match, teams)
	return
}

func GenerateStartGameResponse(startGameRequest StartGameRequest) (response StartGameResponse, err error) {
	match, err := FindMatch(startGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	if result := MatchFull(match, teams); !result {
		err = errors.New(Err_WaitingForPlayers)
		return
	}

	question, err := FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if err = UpdateMatchQuestions(match, question); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	snapshot := InitSnapshotDtoF(match.Id, question.Id, teams[0].Id,
		START.String(), 0, 1, 1, question.Statements)
	if err = CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = InitStartGameResponse(match.Id, teams, question, snapshot)
	return
}

func GenerateQuestionHintResponse(request GameSnapRequest) (response Snapshot, err error) {
	_, err = FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	answer, err := FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	snapshot, err := FindLatestSnapshot(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	snapshot = InitSnapshotDtoF(request.QuizId, request.QuestionId, snapshot.TeamSTurn,
		HINT.String(), 0,
		snapshot.QuestionNo, snapshot.RoundNo, answer.Answer)

	if err = CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GenerateQuestionAnswerResponse(request GameSnapRequest) (response Snapshot, err error) {
	match, err := FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	answer, err := FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	snapshot, err := FindLatestSnapshot(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	snapshot = InitSnapshotDtoF(request.QuizId, request.QuestionId, request.TeamSTurn,
		RIGHT.String(), match.Specs.Points/int(math.Pow(2, float64(snapshot.RoundNo))),
		snapshot.QuestionNo, snapshot.RoundNo, answer.Answer)
	if err = CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GenerateNextQuestionResponse(request GameSnapRequest) (response Snapshot, err error) {
	match, err := FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	if !QuestionCanBeAdded(match) {
		err = errors.New(Err_QuestionsNotLeft)
		return
	}

	question, err := FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if err = UpdateMatchQuestions(match, question); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	snapshot, err := FindLatestSnapshot(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	teamsTurn := NextTeam(match.Teams, request.TeamSTurn)
	snapshot = InitSnapshotDtoF(request.QuestionId, question.Id, teamsTurn, NEXT.String(), 0, snapshot.QuestionNo+1, 1, question.Statements)
	if err = CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GeneratePassQuestionResponse(request GameSnapRequest) (response Snapshot, err error) {
	match, err := FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	snapshots, err := FindQuestionSnapshots(request.QuizId, request.QuestionId)
	if len(snapshots) == 0 || err != nil {
		err = errors.New(Err_SnapshotNotPresent)
	}

	teamsTurn := NextTeam(match.Teams, request.TeamSTurn)
	snapshot := snapshots[len(snapshots)-1]
	snapshot = InitSnapshotDtoF(request.QuizId, request.QuestionId, teamsTurn, PASS.String(), 0, snapshot.QuestionNo, snapshot.RoundNo, []string{})
	if err = CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}

	response = snapshot
	return
}

func GenerateScoreResponse(request ScoreRequest) (response ScoreResponse, err error) {
	snapshots, err := FindSnapshots(request.QuizId)
	if err != nil {
		err = errors.New(Err_SnapshotNotPresent)
		return
	}

	response = InitScoreResponse(request, snapshots)
	return
}
