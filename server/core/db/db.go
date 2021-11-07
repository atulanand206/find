package db

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct{}

func (db DB) PlayersCollections() []string {
	return []string{TeamCollection, SubscriberCollection}
}

func (db DB) QuestionsCollections() []string {
	return []string{IndexCollection, QuestionCollection, AnswerCollection}
}

func (db DB) CreateCollections() (err error) {
	err = mongo.CreateCollections(Database, db.QuestionsCollections())
	return
}

func (db DB) DropCollections() (err error) {
	err = mongo.DropCollections(Database, db.QuestionsCollections())
	return
}

type Schemas struct{}

func (schemas Schemas) Subscriber() (jsonSchema bson.M) {
	return bson.M{
		"bsonType": "object",
		"required": []string{"tag, playerId, role"},
		"properties": bson.M{
			"tag": bson.M{
				"bsonType":    "string",
				"description": "subscriber must have a tag assigned.",
			},
			"player_id": bson.M{
				"bsonType":  "string",
				"describer": "subscriber must have a valid player id.",
			},
			"role": bson.M{
				"bsonType":  "string",
				"describer": "subscriber must have a role assigned.",
			},
		},
	}
}

func (db DB) CreateSubscriberCollection() (err error) {
	err = mongo.CreateCollection(Database, SubscriberCollection,
		options.CreateCollection().SetValidator(bson.M{
			"$jsonSchema": Schemas{}.Subscriber(),
		}))
	return
}

func (db DB) Init() (err error) {
	if err = db.DropCollections(); err != nil {
		return
	}
	if err = db.CreateSubscriberCollection(); err != nil {
		return
	}
	return
}

func (db DB) Create(request interface{}, collection string) (err error) {
	requestDto, err := mongo.Document(request)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, collection, *requestDto)
	return
}
