package db

import (
	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockDB struct {
	Data map[string]map[int]interface{}
}

func (db *MockDB) PlayersCollections() []string {
	return []string{TeamCollection, SubscriberCollection}
}
func (db *MockDB) QuestionsCollections() []string {
	return []string{TeamCollection, SubscriberCollection}
}

func (db *MockDB) Init() {
	db.Data = make(map[string]map[int]interface{})
}

func (db *MockDB) CreateCollection(collection string) {
	db.Data[collection] = make(map[int]interface{})
}

func (db *MockDB) CreateCollections() (err error) {
	db.Data = make(map[string]map[int]interface{})
	return
}

func (db *MockDB) DropCollections() (err error) {
	db.Data = make(map[string]map[int]interface{})
	return
}

func (db *MockDB) Create(request interface{}, collection string) (err error) {
	db.Data[collection][len(db.Data[collection])] = request
	return
}

func (db *MockDB) CreateMany(request []interface{}, collection string) (err error) {
	for _, v := range request {
		db.Data[collection][len(db.Data[collection])] = v
	}
	return
}

func (db *MockDB) FindOne(collection string, filters bson.M, findOptions *options.FindOneOptions) (result *mg.SingleResult) {
	for _, v := range db.Data[collection] {
		var x = 0
		for k := range filters {
			if v.(bson.M)[k] == k {
				x++
			}
		}
		if x == len(filters) {
			result = &mg.SingleResult{}
			result.Decode(v)
			return
		}
	}
	return
}

func (db *MockDB) Find(collection string, filters bson.M, findOptions *options.FindOptions) (result *mg.Cursor, err error) {
	return
}

func (db *MockDB) Delete(collection string, identifier bson.M) (result *mg.DeleteResult, err error) {
	return
}

func (db *MockDB) Update(collection string, identifier bson.M, doc interface{}) (result *mg.UpdateResult, err error) {
	return
}
