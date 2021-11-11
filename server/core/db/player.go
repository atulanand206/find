package db

import (
	e "errors"

	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerCrud struct {
	Db DBConn
}

func (crud PlayerCrud) FindOrCreatePlayer(request models.Player) (player models.Player, err error) {
	player, err = crud.FindPlayer(request.Email)
	if err != nil {
		player = models.InitNewPlayer(request)
		if err = crud.Db.Create(player, PlayerCollection); err != nil {
			err = e.New(errors.Err_PlayerNotCreated)
			return
		}
	}
	return
}

func (crud PlayerCrud) FindPlayer(emailId string) (player models.Player, err error) {
	dto, err := crud.Db.FindOne(PlayerCollection, bson.M{"email": emailId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	err = bson.Unmarshal(dto, &player)
	return
}

func (crud PlayerCrud) FindPlayers(playerIds []string) (players []models.Player, err error) {
	findOptions := &options.FindOptions{}
	cursor, err := crud.Db.Find(PlayerCollection,
		bson.M{"_id": bson.M{"$in": playerIds}}, findOptions)
	if err != nil {
		return
	}
	players, err = DecodePlayers(cursor)
	return
}

func (crud PlayerCrud) UpdatePlayer(player models.Player) (updated bool, err error) {
	res, err := crud.Db.Update(PlayerCollection, bson.M{"_id": player.Id}, player)
	updated = int(res.ModifiedCount) == 1
	return
}
