package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubscriberCrud struct {
	db DB
}

func (crud SubscriberCrud) CreateSubscriber(subscriber Subscriber) error {
	return crud.db.Create(subscriber, SubscriberCollection)
}

func (crud SubscriberCrud) FindTeamPlayers(teams []Team) (teamPlayers []Subscriber, err error) {
	teamIds := make([]string, 0)
	for _, v := range teams {
		teamIds = append(teamIds, v.Id)
	}
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "tag", Value: 1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"tag": bson.M{"$in": teamIds}}, findOptions)
	if err != nil {
		return
	}
	teamPlayers, err = DecodeTeamPlayers(cursor)
	return
}

func (crud SubscriberCrud) FindIds(teamPlayers []Subscriber) (playerIds []string) {
	playerIds = make([]string, 0)
	for _, v := range teamPlayers {
		playerIds = append(playerIds, v.PlayerId)
	}
	return
}

func (crud SubscriberCrud) FindSubscriptionsForPlayerId(playerId string) (subscribers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"player_id": playerId}, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}

func (crud SubscriberCrud) FindSubscribers(tag string, role Role) (subscribers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"tag": tag, "role": role.String()}, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}

func (crud SubscriberCrud) FindSubscribersForTag(tags []string) (subscribers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"tag": bson.M{"$in": tags}}, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}

func (crud SubscriberCrud) FindSubscriberForTagAndPlayerId(tag string, playerId string) (subscriber Subscriber, err error) {
	dto := mongo.FindOne(Database, SubscriberCollection,
		bson.M{"tag": tag, "player_id": playerId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	subscriber, err = DecodeSubscriber(dto)
	return
}

func (crud SubscriberCrud) DeleteSubscriber(playerId string) (err error) {
	_, err = mongo.Delete(Database, SubscriberCollection, bson.M{
		"player_id": playerId,
		"active":    true})
	return
}
