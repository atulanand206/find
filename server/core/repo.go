package core

import "errors"

func FindTeamVacancy(match Game, teams []Team) (team Team, err error) {
	for _, tx := range teams {
		if len(tx.Players) < match.Specs.Players {
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

func FindQuestionForMatch(match Game) (question Question, err error) {
	index, err := FindIndex()
	if err != nil {
		err = errors.New(Err_IndexNotPresent)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), 1)
	questions, err := FindQuestionsFromIndexes(indexes, int64(1))
	if len(questions) != 1 || err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	question = questions[0]
	return
}
