package core

import (
	"errors"
	"fmt"
)

type SubscriberService struct {
	crud SubscriberCrud

	target Target
}

func (service SubscriberService) selfResponse(playerId string, action Action, response interface{}) (res WebsocketMessage, targets map[string]bool) {
	res = MessageCreator.WebSocketsResponse(action, response)
	targets = service.target.TargetSelf(playerId)
	return
}

func (service SubscriberService) quizResponse(action string, quizId string, response Snapshot) (res WebsocketMessage, targets map[string]bool) {
	subscribers, er := service.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = service.target.TargetQuiz(quizId, subscribers)
	res = MessageCreator.WebSocketsResponse(S_GAME, &SnapshotResponse{Action: action, Snapshot: response})
	return
}

func (service SubscriberService) subscribeAndRespond(match Game, player Player, snapshot Snapshot, role Role) (response GameResponse, err error) {
	_, err = service.FindOrCreateSubscriber(match.Id, player, role)
	if err != nil {
		return
	}

	response = GameResponse{Quiz: match, Snapshot: snapshot, Role: role.String()}
	return
}

func (service SubscriberService) FindOrCreateSubscriber(tag string, audience Player, role Role) (subscriber Subscriber, err error) {
	subscriber, err = service.crud.FindSubscriberForTagAndPlayerId(tag, audience.Id)
	if err != nil {
		subscriber = InstanceCreator.InitSubscriber(tag, audience, role.String())
		err = service.crud.CreateSubscriber(subscriber)
		if err != nil {
			err = errors.New(fmt.Sprint(ErrorCreator.SubscriberNotCreated(subscriber)))
		}
	}
	return
}

func (service SubscriberService) FindSubscribersForTag(tags []string) (subscribers []Subscriber, err error) {
	return service.crud.FindSubscribersForTag(tags)
}

func (service SubscriberService) FindTeamPlayers(teams []Team) (teamPlayers []Subscriber, err error) {
	return service.crud.FindTeamPlayers(teams)
}

func (service SubscriberService) FindSubscriptionsForPlayerId(playerId string) (subscribers []Subscriber, err error) {
	return service.crud.FindSubscriptionsForPlayerId(playerId)
}
