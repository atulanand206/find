package core

import "errors"

type PlayerService struct {
	crud PlayerCrud
}

func (service PlayerService) FindOrCreatePlayer(request Player) (player Player, err error) {
	return service.crud.FindOrCreatePlayer(request)
}

func (service PlayerService) FindPlayerByEmail(email string) (player Player, err error) {
	player, err = service.crud.FindPlayer(email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	return
}
