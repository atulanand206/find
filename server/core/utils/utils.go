package utils

import (
	"math/rand"
	"time"

	"github.com/atulanand206/find/core/models"
)

func CreateDeckDtos() (indexs []models.Index, questions []models.Question, answers []models.Answer, err error) {
	indexsWrapper, err := ReadIndexes()
	if err != nil {
		return
	}
	indexs, questions, answers = CreateDeck(indexsWrapper)
	return
}

func CreateDeck(indexsWrapper models.IndexStore) (indexs []models.Index, questions []models.Question, answers []models.Answer) {
	indexs = make([]models.Index, 0)
	questions = make([]models.Question, 0)
	answers = make([]models.Answer, 0)

	for _, index := range indexsWrapper.Indexes {
		idx := models.InitNewIndex(index)
		indexs = append(indexs, idx)
		newQuestions := ReadQuestionsForIndex(idx)

		for _, newQuestion := range newQuestions {
			question := models.InitNewQuestion(idx, newQuestion)
			questions = append(questions, question)

			answers = append(answers, models.InitNewAnswer(question, newQuestion))
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

func FilterIndex(input []models.Index, sans map[string]bool, limit int) (indexes []models.Index) {
	return randomFilteredIndicesList(randomFilteredIndicesMap(filteredIndices(input, sans), limit))
}

func filteredIndices(input []models.Index, sans map[string]bool) (indxs []models.Index) {
	indxs = make([]models.Index, 0)
	for _, indx := range input {
		if !sans[indx.Id] {
			indxs = append(indxs, indx)
		}
	}
	return
}

func randomFilteredIndicesMap(indxs []models.Index, limit int) (mp map[models.Index]bool) {
	mp = make(map[models.Index]bool)
	rand.Seed(time.Now().UnixNano())
	indexCount := len(indxs)
	for {
		if len(mp) < limit && indexCount > 0 {
			tag := indxs[rand.Intn(indexCount)]
			mp[tag] = true
		} else {
			break
		}
	}
	return
}

func randomFilteredIndicesList(mp map[models.Index]bool) (indexes []models.Index) {
	for key := range mp {
		indexes = append(indexes, key)
	}
	return
}

func IsQuizMasterInMatch(match models.Game, person models.Player) (result bool) {
	return match.QuizMaster.Id == person.Id
}

func IsPlayerInTeams(teamPlayers []models.Subscriber, person models.Player) (result bool) {
	for _, player := range teamPlayers {
		if player.PlayerId == person.Id {
			result = true
			return
		}
	}
	return
}

func NewTeam(team models.Team) (result bool) {
	return team.Id == ""
}

func NextTeam(teams []models.TeamRoster, teamsTurn string) (teamId string) {
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

func MatchFull(match models.Game, teamPlayers []models.Subscriber) (result bool) {
	return len(teamPlayers) == match.Specs.Players*match.Specs.Teams
}

func QuestionCanBeAdded(match models.Game) (result bool) {
	return len(match.Tags) < match.Specs.Questions
}

func TableRoster(teams []models.Team, teamPlayers []models.Subscriber, players []models.Player) (roster []models.TeamRoster) {
	roster = make([]models.TeamRoster, 0)
	for _, v := range teams {
		var entry models.TeamRoster
		entry.Id = v.Id
		entry.Name = v.Name
		entry.Score = v.Score
		playrs := make([]models.Player, 0)
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
