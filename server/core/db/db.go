package db

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mg "go.mongodb.org/mongo-driver/mongo"
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

func (db DB) Create(request interface{}, collection string) (err error) {
	requestDto, err := mongo.Document(request)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, collection, *requestDto)
	return
}

func (db DB) CreateMany(request []interface{}, collection string) (err error) {
	_, err = mongo.WriteMany(Database, collection, request)
	return
}

func (db DB) FindOne(collection string, filters bson.M, findOptions *options.FindOneOptions) (result *mg.SingleResult) {
	return mongo.FindOne(Database, collection, filters, findOptions)
}

func (db DB) Find(collection string, filters bson.M, findOptions *options.FindOptions) (result *mg.Cursor, err error) {
	return mongo.Find(Database, collection, filters, findOptions)
}

func (db DB) Delete(collection string, identifier bson.M) (result *mg.DeleteResult, err error) {
	return mongo.Delete(Database, collection, identifier)
}

func (db DB) Update(collection string, identifier bson.M, doc interface{}) (result *mg.UpdateResult, err error) {
	requestDto, err := mongo.Document(doc)
	if err != nil {
		return
	}
	return mongo.Update(Database, collection, identifier, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
}
