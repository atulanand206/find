package db

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConn interface {
	AllCollections() []string
	CreateCollections() (err error)
	DropCollections() (err error)
	Create(request interface{}, collection string) (err error)
	CreateMany(request []interface{}, collection string) (err error)
	FindOne(collection string, filters bson.M, findOptions *options.FindOneOptions) (result bson.Raw, err error)
	Find(collection string, filters bson.M, findOptions *options.FindOptions) (result []bson.Raw, err error)
	Delete(collection string, identifier bson.M) (result *mg.DeleteResult, err error)
	Update(collection string, identifier bson.M, doc interface{}) (result *mg.UpdateResult, err error)
}

type DB struct{}

func NewDb() DB {
	mockDb := DB{}
	return mockDb
}

func (db DB) AllCollections() []string {
	return []string{MatchCollection, QuestionCollection, AnswerCollection, SnapshotCollection, TeamCollection, PlayerCollection, IndexCollection, SubscriberCollection}
}

func (db DB) CreateCollections() (err error) {
	err = mongo.CreateCollections(Database, db.AllCollections())
	return
}

func (db DB) DropCollections() (err error) {
	err = mongo.DropCollections(Database, db.AllCollections())
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

func (db DB) FindOne(collection string, filters bson.M, findOptions *options.FindOneOptions) (result bson.Raw, err error) {
	res := mongo.FindOne(Database, collection, filters, findOptions)
	if err = res.Err(); err != nil {
		return res.DecodeBytes()
	}
	return
}

func (db DB) Find(collection string, filters bson.M, findOptions *options.FindOptions) (result []bson.Raw, err error) {
	cursor, err := mongo.Find(Database, collection, filters, findOptions)
	if err != nil {
		return
	}
	result, err = DecodeRaw(cursor)
	return
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
