package core

import "errors"

type PlayerService struct {
	db DB
}

func (service PlayerService) FindOrCreatePlayer(request Player) (player Player, err error) {
	player, err = service.db.FindPlayer(request.Email)
	if err != nil {
		player = InitNewPlayer(request)
		if err = service.db.CreatePlayer(player); err != nil {
			err = errors.New(Err_PlayerNotCreated)
			return
		}
	}
	return
}

func (service PlayerService) FindPlayerByEmail(email string) (player Player, err error) {
	player, err = service.db.FindPlayer(email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	return
}

func (service PlayerService) DeletePlayerLiveSession(playerId string) (res WebsocketMessage, targets map[string]bool, err error) {
	subscribers, err := service.db.FindSubscriptionsForPlayerId(playerId)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	tags := make([]string, 0)
	for _, subscriber := range subscribers {
		tags = append(tags, subscriber.Tag)
	}

	subscribers, err = service.db.FindSubscribersForTag(tags)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		if playerId != subscriber.PlayerId {
			targets[subscriber.PlayerId] = true
		}
	}

	res = MessageCreator.InitWebSocketMessage(S_REFRESH, "Player dropped. Please refresh.")
	err = service.db.DeleteSubscriber(playerId)
	return
}
