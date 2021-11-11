package services

import (
	"errors"
	"fmt"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
)

type SubscriberService struct {
	Crud db.SubscriberCrud

	TargetService TargetService

	Creators Creators
}

func (service SubscriberService) SelfResponse(playerId string, action actions.Action, response interface{}) (res models.WebsocketMessage, targets map[string]bool) {
	res = service.Creators.MessageCreator.WebSocketsResponse(action, response)
	targets = service.TargetService.TargetSelf(playerId)
	return
}

func (service SubscriberService) QuizResponse(action string, quizId string, response models.Snapshot) (res models.WebsocketMessage, targets map[string]bool) {
	subscribers, er := service.FindSubscribersForTags([]string{quizId})
	if er != nil {
		res = service.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = service.TargetService.TargetQuiz(quizId, subscribers)
	res = service.Creators.MessageCreator.WebSocketsResponse(actions.S_GAME, &models.SnapshotResponse{Action: action, Snapshot: response})
	return
}

func (service SubscriberService) FindOrCreateSubscriber(tag string, audience models.Player, role actions.Role) (subscriber models.Subscriber, err error) {
	subscriber, err = service.Crud.FindSubscriber(tag, audience.Id)
	if err != nil {
		subscriber = service.Creators.InstanceCreator.InitSubscriber(tag, audience, role.String())
		err = service.Crud.CreateSubscriber(subscriber)
		if err != nil {
			err = errors.New(fmt.Sprint(service.Creators.ErrorCreator.SubscriberNotCreated(subscriber)))
		}
	}
	return
}

func (service SubscriberService) SubscribeAndRespond(match models.Game, player models.Player, snapshot models.Snapshot, role actions.Role) (response models.GameResponse, err error) {
	_, err = service.FindOrCreateSubscriber(match.Id, player, role)
	if err != nil {
		return
	}

	response = models.GameResponse{Quiz: match, Snapshot: snapshot, Role: role.String()}
	return
}

func (service SubscriberService) FindSubscribersForTags(tags []string) (subscribers []models.Subscriber, err error) {
	return service.Crud.FindSubscribers(bson.M{"tag": bson.M{"$in": tags}})
}

func (service SubscriberService) FindSubscriptionsForPlayerId(playerId string) (subscribers []models.Subscriber, err error) {
	return service.Crud.FindSubscribers(bson.M{"player_id": playerId})
}

func (service SubscriberService) FindSubscribers(tag string, role actions.Role) (subscribers []models.Subscriber, err error) {
	return service.Crud.FindSubscribers(bson.M{"tag": tag, "role": role.String()})
}

func (service SubscriberService) FindPlayerIdsFromSubscribers(subscribers []models.Subscriber) (playerIds []string) {
	playerIds = make([]string, 0)
	for _, subscriber := range subscribers {
		playerIds = append(playerIds, subscriber.PlayerId)
	}
	return playerIds
}
