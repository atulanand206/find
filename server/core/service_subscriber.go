package core

import (
	"errors"
	"fmt"
)

type SubscriberService struct {
	db DB

	target Target
}

func (service SubscriberService) selfResponse(quizId string, action Action, response interface{}) (res WebsocketMessage, targets map[string]bool) {
	res = MessageCreator.WebSocketsResponse(action, response)
	targets = Controller.subscriberService.target.TargetSelf(quizId)
	return
}

func (service SubscriberService) joinResponse(quizId string, response GameResponse) (res WebsocketMessage, targets map[string]bool) {
	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = Controller.subscriberService.target.TargetQuiz(quizId, subscribers)
	res = MessageCreator.WebSocketsResponse(S_GAME, response)
	return
}

func (service SubscriberService) quizResponse(quizId string, response Snapshot) (res WebsocketMessage, targets map[string]bool) {
	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = Controller.subscriberService.target.TargetQuiz(quizId, subscribers)
	res = MessageCreator.WebSocketsResponse(S_GAME, response)
	return
}

func (service SubscriberService) subscribeAndRespond(match Game, player Player, snapshot Snapshot, role Role) (response GameResponse, err error) {
	_, err = service.FindOrCreateSubscriber(match.Id, player, role)
	if err != nil {
		return
	}

	response = GameResponse{Quiz: match, Snapshot: snapshot}
	return
}

func (service SubscriberService) FindOrCreateSubscriber(tag string, audience Player, role Role) (subscriber Subscriber, err error) {
	subscriber, err = service.db.FindSubscriberForTagAndPlayerId(tag, audience.Id)
	if err != nil {
		subscriber = InstanceCreator.InitSubscriber(tag, audience, role.String())
		err = service.db.CreateSubscriber(subscriber)
		if err != nil {
			err = errors.New(fmt.Sprint(ErrorCreator.SubscriberNotCreated(subscriber)))
		}
	}
	return
}

func (service SubscriberService) FindSubscribersForTag(tags []string) (subscribers []Subscriber, err error) {
	return service.db.FindSubscribersForTag(tags)
}
