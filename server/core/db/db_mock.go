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
	doc, _ := Document(request)
	db.Data[collection][len(db.Data[collection])] = doc
	return
}

func (db *MockDB) CreateMany(request []interface{}, collection string) (err error) {
	for _, v := range request {
		doc, _ := Document(v)
		db.Data[collection][len(db.Data[collection])] = doc
	}
	return
}

func Document(request interface{}) (doc bson.M, err error) {
	data, err := bson.Marshal(request)
	if err != nil {
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func (db *MockDB) FindOne(collection string, filters bson.M, findOptions *options.FindOneOptions) (result bson.Raw, err error) {
	for _, entry := range db.Data[collection] {
		var x = 0
		for k, filter := range filters {
			if entry.(bson.M)[k] == filter {
				x++
			}
		}
		if x == len(filters) {
			result, _ = bson.Marshal(entry)
			return
		}
	}
	return
}

func (db *MockDB) Find(collection string, filters bson.M, findOptions *options.FindOptions) (result *mg.Cursor, err error) {
	return
}

func (db *MockDB) Delete(collection string, identifier bson.M) (result *mg.DeleteResult, err error) {
	for key, entry := range db.Data[collection] {
		var x = 0
		for k, filter := range identifier {
			if entry.(bson.M)[k] == filter {
				x++
			}
		}
		if x == len(identifier) {
			delete(db.Data[collection], key)
			return
		}
	}
	return
}

func (db *MockDB) Update(collection string, identifier bson.M, v interface{}) (result *mg.UpdateResult, err error) {
	for _, entry := range db.Data[collection] {
		var x = 0
		for k, filter := range identifier {
			if entry.(bson.M)[k] == filter {
				x++
			}
		}
		if x == len(identifier) {
			doc, _ := Document(v)
			db.Data[collection][len(db.Data[collection])] = doc
			return
		}
	}
	return
}
