package services

import (
	"errors"
	"fmt"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
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
	subscribers, er := service.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = service.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = service.TargetService.TargetQuiz(quizId, subscribers)
	res = service.Creators.MessageCreator.WebSocketsResponse(actions.S_GAME, &models.SnapshotResponse{Action: action, Snapshot: response})
	return
}

func (service SubscriberService) FindOrCreateSubscriber(tag string, audience models.Player, role actions.Role) (subscriber models.Subscriber, err error) {
	subscriber, err = service.Crud.FindSubscriberForTagAndPlayerId(tag, audience.Id)
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

func (service SubscriberService) FindSubscribersForTag(tags []string) (subscribers []models.Subscriber, err error) {
	return service.Crud.FindSubscribersForTag(tags)
}

func (service SubscriberService) FindTeamPlayers(teams []models.Team) (teamPlayers []models.Subscriber, err error) {
	return service.Crud.FindTeamPlayers(teams)
}

func (service SubscriberService) FindSubscriptionsForPlayerId(playerId string) (subscribers []models.Subscriber, err error) {
	return service.Crud.FindSubscriptionsForPlayerId(playerId)
}
