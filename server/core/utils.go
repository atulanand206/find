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

func IsPlayerInTeam(team Team, person Player) (result bool) {
	for _, player := range team.Players {
		if player.Id == person.Id {
			result = true
			return
		}
	}
	return
}

func NewTeam(team Team) (result bool) {
	return team.Id == ""
}

func NextTeam(teams []TeamMini, teamsTurn string) (teamId string) {
	teamId = ""
	if len(teams) == 0 {
		return
	}
	for _, v := range teams {
		if teamId != "" {
			teamId = v.Id
			return
		}
		if v.Id == teamsTurn {
			teamId = v.Id
		}
	}
	teamId = teams[0].Id
	return
}

func PlayerCanBeAdded(match Game, team Team) (result bool) {
	if len(team.Players) < match.Specs.Players {
		result = true
		return
	}
	return
}

func MatchFull(match Game, teams []Team) (result bool) {
	for _, v := range teams {
		if len(v.Players) < match.Specs.Players {
			return false
		}
	}
	return true
}

func QuestionCanBeAdded(match Game) (result bool) {
	return len(match.Tags) < match.Specs.Questions
}
