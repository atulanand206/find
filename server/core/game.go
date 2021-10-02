package core

import (
	"errors"
	"fmt"
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

	team, err := FindTeamInMatch(teams, enterGameRequest.TeamId)
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

func FindTeamInMatch(teams []Team, teamId string) (team Team, err error) {
	for _, tx := range teams {
		if tx.Id == teamId {
			team = tx
			return
		}
	}
	err = errors.New(Err_TeamsNotPresentInMatch)
	return
}

func FindOrCreatePlayer(request Player) (player Player, err error) {
	player, err = FindPlayer(request.Email)
	if err != nil {
		player = InitNewPlayer(request)
		if err = CreatePlayer(player); err != nil {
			err = errors.New(Err_PlayerNotCreated)
			return
		}
	}
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

	questions, err := FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if err = UpdateMatchQuestions(match, questions); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	response = InitStartGameResponse(match.Id, teams, questions)
	return
}

func GenerateQuestionHintResponse(request GameSnapRequest) (response HintRevealResponse, err error) {
	_, err = FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	answer, err := FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	response = InitHintRevealResponse(request, answer)
	return
}

func GenerateQuestionAnswerResponse(request GameSnapRequest) (response AnswerRevealResponse, err error) {
	_, err = FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	answer, err := FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
	}

	response = InitAnswerRevealResponse(request, answer)
	return
}

func GenerateNextQuestionResponse(request NextQuestionRequest) (response GameNextResponse, err error) {
	match, err := FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	if !QuestionCanBeAdded(match) {
		err = errors.New(Err_QuestionsNotLeft)
		return
	}

	questions, err := FindQuestionForMatch(match)
	if err != nil {
		return
	}
	question := questions[0]

	if err = UpdateMatchQuestions(match, questions); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	teamsTurn := NextTeam(match.Teams, request.TeamSTurn)
	response = InitNextQuestionResponse(request, question, teamsTurn)
	return
}

func FindQuestionForMatch(match Game) (questions []Question, err error) {
	index, err := FindIndex()
	if err != nil {
		err = errors.New(Err_IndexNotPresent)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), 1)
	questions, err = FindQuestionsFromIndexes(indexes, int64(1))
	if len(questions) == 0 {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	if err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}
	return
}

func GenerateFindAnswerResponse(request FindAnswerRequest) (answer Answer, err error) {
	answer, err = FindAnswer(request.QuestionId)
	if err != nil {
		err = errors.New(Err_AnswerNotPresent)
		return
	}
	return
}

func GeneratePassQuestionResponse(request GameSnapRequest) (response GamePassResponse, err error) {
	match, err := FindMatch(request.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	teamsTurn := NextTeam(match.Teams, request.TeamSTurn)
	fmt.Println(teamsTurn)
	response = InitPassQuestionResponse(request, teamsTurn)
	return
}
