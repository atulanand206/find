package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerCrud struct {
	db DB
}

func (crud PlayerCrud) CreatePlayer(player Player) error {
	return crud.db.Create(player, PlayerCollection)
}

func (crud PlayerCrud) FindPlayer(emailId string) (player Player, err error) {
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": emailId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	player, err = DecodePlayer(dto)
	return
}

func (crud PlayerCrud) FindPlayers(teamPlayers []Subscriber) (players []Player, err error) {
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
