package core

type Target struct{}

func (target Target) TargetSelf(playerId string) (targets map[string]bool) {
	targets = make(map[string]bool)
	targets[playerId] = true
	return
}

func (target Target) TargetQuiz(quizId string, subscribers []Subscriber) (targets map[string]bool) {
	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}
	return
}
