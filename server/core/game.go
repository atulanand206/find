package core

import (
	"errors"
	"math"
)

func GenerateActiveQuizResponse() (matches []Game, err error) {
	matches, err = FindActiveMatches()
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
	}
	return
}

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

	teams := InitNewTeams(quiz)
	if err = CreateTeams(teams); err != nil {
		err = errors.New(Err_TeamNotCreated)
	}

	response = InitEnterGameResponse(quiz, teams, []TeamPlayer{}, []Player{}, "", Snapshot{})
	return
}

func GenerateEnterGameResponse(enterGameRequest EnterGameRequest) (response EnterGameResponse, err error) {
	player, err := FindPlayer(enterGameRequest.Person.Email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	match, err := FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	var snapshot Snapshot
	if match.Active {
		snapshot, err = FindLatestSnapshot(match.Id)
		if err != nil {
			err = errors.New(Err_SnapshotNotPresent)
			return
		}
	}

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamPlayers, err := FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamPlayersNotPresent)
		return
	}

	players, err := FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	if IsQuizMasterInMatch(match, player) {
		response = InitEnterGameResponse(match, teams, teamPlayers, players, "", snapshot)
		return
	}

	if IsPlayerInTeams(teamPlayers, player) {
		response = InitEnterGameResponse(match, teams, teamPlayers, players, TeamIdForPlayer(teamPlayers, player), snapshot)
		return
	}

	teamId, err := FindTeamVacancy(match, teams, teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}

	if err = CreateTeamPlayer(InitTeamPlayer(teamId, player)); err != nil {
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

	teamPlayers, err = FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	players, err = FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	response = InitEnterGameResponse(match, teams, teamPlayers, players, TeamIdForPlayer(teamPlayers, player), snapshot)
	return
}

func GenerateWatchGameResponse(enterGameRequest EnterGameRequest) (response EnterGameResponse, err error) {
	_, err = FindPlayer(enterGameRequest.Person.Email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	match, err := FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamPlayers, err := FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	players, err := FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	var snapshot Snapshot
	if match.Active {
		snapshot, err = FindLatestSnapshot(match.Id)
		if err != nil {
			err = errors.New(Err_SnapshotNotPresent)
			return
		}
	}

	response = InitEnterGameResponse(match, teams, teamPlayers, players, "", snapshot)
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

	teamPlayers, err := FindTeamPlayers(teams)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
		return
	}

	if result := MatchFull(match, teamPlayers); !result {
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

	players, err := FindPlayers(teamPlayers)
	if err != nil {
		err = errors.New(Err_PlayerNotPresent)
		return
	}

	response = InitStartGameResponse(match.Id, teams, teamPlayers, players, question, snapshot)
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

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamsTurn := NextTeam(teams, request.TeamSTurn)
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

	teams, err := FindTeams(match)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	teamsTurn := NextTeam(teams, request.TeamSTurn)
	snapshot := snapshots[len(snapshots)-1]
	snapshot = InitSnapshotDtoF(match.Id, request.QuestionId, teamsTurn, PASS.String(), 0, snapshot.QuestionNo, snapshot.RoundNo, []string{})
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
