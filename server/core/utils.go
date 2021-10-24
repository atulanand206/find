package core

import (
	"math/rand"
	"time"
)

func CreateDeckDtos() (indexs []Index, questions []Question, answers []Answer, err error) {
	indexsWrapper, err := ReadIndexes()
	if err != nil {
		return
	}
	indexs, questions, answers = CreateDeck(indexsWrapper)
	return
}

func CreateDeck(indexsWrapper IndexStore) (indexs []Index, questions []Question, answers []Answer) {
	indexs = make([]Index, 0)
	questions = make([]Question, 0)
	answers = make([]Answer, 0)

	for _, index := range indexsWrapper.Indexes {
		idx := InitNewIndex(index)
		indexs = append(indexs, idx)
		newQuestions := ReadQuestionsForIndex(idx)

		for _, newQuestion := range newQuestions {
			question := InitNewQuestion(idx, newQuestion)
			questions = append(questions, question)

			answers = append(answers, InitNewAnswer(question, newQuestion))
		}
	}
	return
}

func MapSansTags(tags []string) (tagsMap map[string]bool) {
	tagsMap = make(map[string]bool)
	for _, tag := range tags {
		tagsMap[tag] = true
	}
	return
}

func FilterIndex(input []Index, sans map[string]bool, limit int) (indexes []Index) {
	return randomFilteredIndicesList(randomFilteredIndicesMap(filteredIndices(input, sans), limit))
}

func filteredIndices(input []Index, sans map[string]bool) (indxs []Index) {
	indxs = make([]Index, 0)
	for _, indx := range input {
		if !sans[indx.Id] {
			indxs = append(indxs, indx)
		}
	}
	return
}

func randomFilteredIndicesMap(indxs []Index, limit int) (mp map[Index]bool) {
	mp = make(map[Index]bool)
	rand.Seed(time.Now().UnixNano())
	indexCount := len(indxs)
	for {
		if len(mp) < limit {
			tag := indxs[rand.Intn(indexCount)]
			mp[tag] = true
		} else {
			break
		}
	}
	return
}

func randomFilteredIndicesList(mp map[Index]bool) (indexes []Index) {
	for key := range mp {
		indexes = append(indexes, key)
	}
	return
}

func IsQuizMasterInMatch(match Game, person Player) (result bool) {
	return match.QuizMaster.Id == person.Id
}

func IsPlayerInTeams(teamPlayers []Subscriber, person Player) (result bool) {
	for _, player := range teamPlayers {
		if player.PlayerId == person.Id {
			result = true
			return
		}
	}
	return
}

func NewTeam(team Team) (result bool) {
	return team.Id == ""
}

func NextTeam(teams []TeamRoster, teamsTurn string) (teamId string) {
	teamId = ""
	if len(teams) == 0 {
		return
	}
	idx := 0
	for i, v := range teams {
		if v.Id == teamsTurn {
			idx = i
		}
	}
	idx = (idx + 1) % len(teams)
	teamId = teams[idx].Id
	return
}

func MatchFull(match Game, teamPlayers []Subscriber) (result bool) {
	return len(teamPlayers) == match.Specs.Players*match.Specs.Teams
}

func QuestionCanBeAdded(match Game) (result bool) {
	return len(match.Tags) < match.Specs.Questions
}

func TableRoster(teams []Team, teamPlayers []Subscriber, players []Player) (roster []TeamRoster) {
	roster = make([]TeamRoster, 0)
	for _, v := range teams {
		var entry TeamRoster
		entry.Id = v.Id
		entry.Name = v.Name
		entry.Score = v.Score
		playrs := make([]Player, 0)
		for _, w := range teamPlayers {
			if w.Tag == v.Id {
				for _, x := range players {
					if x.Id == w.PlayerId {
						playrs = append(playrs, x)
					}
				}
			}
		}
		entry.Players = playrs
		roster = append(roster, entry)
	}
	return
}

func (service TeamService) FindVacantTeamId(teams []Team, teamPlayers []Subscriber, playersCount int) (teamId string) {
	mp := make(map[string]int)
	for _, v := range teams {
		mp[v.Id] = 0
	}
	for _, v := range teamPlayers {
		mp[v.Tag] = mp[v.Tag] + 1
	}
	var x = playersCount
	for _, v := range mp {
		if v < x {
			x = v
		}
	}
	for k, v := range mp {
		if v == x {
			teamId = k
			return
		}
	}
	return
}
