package services

import (
	"errors"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
)

type PlayerService struct {
	Crud db.PlayerCrud
}

func (service PlayerService) FindOrCreatePlayer(request models.Player) (player models.Player, err error) {
	return service.Crud.FindOrCreatePlayer(request)
}

func (service PlayerService) FindPlayerByEmail(email string) (player models.Player, err error) {
	player, err = service.Crud.FindPlayer(email)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	return
}
