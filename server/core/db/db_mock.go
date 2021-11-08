package db

import (
	"fmt"

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

func NewMockDb() *MockDB {
	mockDb := &MockDB{}
	mockDb.CreateCollections()
	return mockDb
}

func (db *MockDB) CreateCollection(collection string) {
	if _, ok := db.Data[collection]; !ok {
		db.Data[collection] = make(map[int]interface{})
	}
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
	db.CreateCollection(collection)
	doc, _ := Document(request)
	db.Data[collection][len(db.Data[collection])] = doc
	return
}

func (db *MockDB) CreateMany(request []interface{}, collection string) (err error) {
	db.CreateCollection(collection)
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
		if db.isFilterMatch(entry, filters) {
			result, _ = bson.Marshal(entry)
			return
		}
	}
	err = fmt.Errorf("No document found")
	return
}

func (db *MockDB) Find(collection string, filters bson.M, findOptions *options.FindOptions) (results []bson.Raw, err error) {
	results = make([]bson.Raw, 0)
	for _, entry := range db.Data[collection] {
		if db.isFilterMatch(entry, filters) {
			result, _ := bson.Marshal(entry)
			results = append(results, result)
		}
	}
	if len(results) == 0 {
		err = fmt.Errorf("No document found")
	}
	return
}

func (db *MockDB) Delete(collection string, identifier bson.M) (result *mg.DeleteResult, err error) {
	for key, entry := range db.Data[collection] {
		if db.isFilterMatch(entry, identifier) {
			delete(db.Data[collection], key)
			return
		}
	}
	err = fmt.Errorf("No document found")
	return
}

func (db *MockDB) Update(collection string, identifier bson.M, v interface{}) (result *mg.UpdateResult, err error) {
	for _, entry := range db.Data[collection] {
		if db.isFilterMatch(entry, identifier) {
			doc, _ := Document(v)
			db.Data[collection][len(db.Data[collection])] = doc
			return
		}
	}
	err = fmt.Errorf("No document found")
	return
}

func (db *MockDB) isFilterMatch(entry interface{}, filters bson.M) (bool bool) {
	var x = 0
	for k, filter := range filters {
		if entry.(bson.M)[k] == filter {
			x++
		}
	}
	return x == len(filters)
}
