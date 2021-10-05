package core

import (
	"errors"
)

func FindTeamVacancy(match Game, teams []Team, teamPlayers []TeamPlayer) (teamId string, err error) {
	if len(teamPlayers) >= match.Specs.Players*match.Specs.Teams {
		err = errors.New(Err_PlayersFullInTeam)
		return
	}
	if len(teamPlayers) < len(teams) {
		teamId = teams[len(teamPlayers)].Id
		return
	}
	mp := make(map[string]int)
	for _, v := range teamPlayers {
		if mp[v.TeamId] == 0 {
			mp[v.TeamId] = 1
		} else {
			mp[v.TeamId] = mp[v.TeamId] + 1
		}
	}
	for k, v := range mp {
		if v < match.Specs.Players {
			teamId = k
			return
		}
	}
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

func TeamIdForPlayer(teamPlayers []TeamPlayer, player Player) (teamId string) {
	for _, v := range teamPlayers {
		if v.PlayerId == player.Id {
			teamId = v.TeamId
			return
		}
	}
	return
}
