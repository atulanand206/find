package core

import "errors"

type PlayerService struct {
	crud PlayerCrud
}

func (service PlayerService) FindOrCreatePlayer(request Player) (player Player, err error) {
	player, err = service.crud.FindPlayer(request.Email)
	if err != nil {
		player = InitNewPlayer(request)
		if err = service.crud.CreatePlayer(player); err != nil {
			err = errors.New(Err_PlayerNotCreated)
			return
		}
	}
	return
}

func (service PlayerService) FindPlayerByEmail(email string) (player Player, err error) {
	player, err = service.crud.FindPlayer(email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	return
}
