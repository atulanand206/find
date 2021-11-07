package services

import "github.com/atulanand206/find/server/core/models"

type TargetService struct{}

func (target TargetService) TargetSelf(playerId string) (targets map[string]bool) {
	targets = make(map[string]bool)
	targets[playerId] = true
	return
}

func (target TargetService) TargetQuiz(quizId string, subscribers []models.Subscriber) (targets map[string]bool) {
	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}
	return
}
