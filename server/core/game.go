package core

import "errors"

func GenerateBeginGameResponse(player Player) (res Player, err error) {
	player, err = FindOrCreatePlayer(player.Email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	return
}

func GenerateCreateGameResponse(request CreateGameRequest) (quiz Game, err error) {
	player := request.Quizmaster
	player, err = FindOrCreatePlayer(player.Email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	quiz = InitNewMatch(player, request.Specs)
	if err = CreateMatch(quiz); err != nil {
		err = errors.New(Err_MatchNotCreated)
	}
	return
}

func GenerateEnterGameResponse(enterGameRequest EnterGameRequest) (match Game, err error) {
	match, err = FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	team, err := FindTeamInMatch(match, enterGameRequest.TeamId)
	if err != nil {
		err = errors.New(Err_TeamsNotPresentInMatch)
	}

	player := enterGameRequest.Person
	if IsQuizMasterInMatch(match, player) {
		err = errors.New(Err_QuizmasterCantPlay)
		return
	}

	if IsPlayerInMatch(match, player) {
		err = errors.New(Err_PlayerAlreadyInGame)
		return
	}

	if !PlayerCanBeAdded(team, match.Specs.Players) {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}

	player, err = FindPlayer(player.Email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	if err = UpdateMatchPlayer(match, player, team.Id); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

	match, err = FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}
	return
}

func FindTeamInMatch(match Game, teamId string) (team Team, err error) {
	for _, tx := range match.Teams {
		if tx.Id == teamId {
			team = tx
			return
		}
	}
	err = errors.New(Err_TeamsNotPresentInMatch)
	return
}

func FindOrCreatePlayer(email string) (player Player, err error) {
	player, err = FindPlayer(email)
	if err != nil {
		player = InitNewPlayer(player)
		if err = CreatePlayer(player); err != nil {
			err = errors.New(Err_PlayerNotCreated)
			return
		}
	}
	return
}

func GenerateWatchGameResponse(enterGameRequest EnterGameRequest) (match Game, err error) {
	match, err = FindMatch(enterGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}
	return
}

func GenerateStartGameResponse(startGameRequest StartGameRequest) (response StartGameResponse, err error) {
	match, err := FindMatch(startGameRequest.QuizId)
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		return
	}

	if !CanStart(match) {
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

	response = InitStartGameResponse(match, questions)
	return
}

func GenerateNextQuestionResponse(nextQuestionRequest NextQuestionRequest) (question Question, err error) {
	match, err := FindMatch(nextQuestionRequest.QuizId)
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
	question = questions[0]

	if err = UpdateMatchQuestions(match, questions); err != nil {
		err = errors.New(Err_MatchNotUpdated)
		return
	}

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
