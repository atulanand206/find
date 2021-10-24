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
