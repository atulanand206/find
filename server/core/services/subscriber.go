package services

import (
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
)

type SubscriberService struct {
	Crud db.SubscriberCrud

	targetService TargetService
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
