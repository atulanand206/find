package db

import (
	e "errors"

	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerCrud struct {
	Db DB
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
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": emailId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	player, err = DecodePlayer(dto)
	return
}

func (crud PlayerCrud) FindPlayers(teamPlayers []models.Subscriber) (players []models.Player, err error) {
	playerIds := make([]string, 0)
	for _, v := range teamPlayers {
		playerIds = append(playerIds, v.PlayerId)
	}
	findOptions := &options.FindOptions{}
	cursor, err := mongo.Find(Database, PlayerCollection,
		bson.M{"_id": bson.M{"$in": playerIds}}, findOptions)
	if err != nil {
		return
	}
	players, err = DecodePlayers(cursor)
	return
}

func (crud PlayerCrud) UpdatePlayer(player models.Player) (updated bool, err error) {
	requestDto, err := mongo.Document(player)
	if err != nil {
		return
	}
	res, err := mongo.Update(Database, PlayerCollection, bson.M{"_id": player.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	updated = int(res.ModifiedCount) == 1
	return
}