package db

import (
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubscriberCrud struct {
	Db DBConn
}

func (crud SubscriberCrud) CreateSubscriber(subscriber models.Subscriber) error {
	return crud.Db.Create(subscriber, SubscriberCollection)
}

func (crud SubscriberCrud) FindSubscribers(filters bson.M) (subscribers []models.Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "tag", Value: 1})
	findOptions.SetSort(sort)
	cursor, err := crud.Db.Find(SubscriberCollection, filters, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}

func (crud SubscriberCrud) FindSubscriber(tag string, playerId string) (subscriber models.Subscriber, err error) {
	dto, err := crud.Db.FindOne(SubscriberCollection,
		bson.M{"tag": tag, "player_id": playerId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	bson.Unmarshal(dto, &subscriber)
	return
}

func (crud SubscriberCrud) DeleteSubscriber(playerId string) (err error) {
	_, err = crud.Db.Delete(SubscriberCollection, bson.M{
		"player_id": playerId,
		"active":    true})
	return
}
