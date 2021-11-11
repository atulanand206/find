package db

import (
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SnapshotCrud struct {
	Db DBConn
}

func (crud SnapshotCrud) CreateSnapshot(snapshot models.Snapshot) error {
	return crud.Db.Create(snapshot, SnapshotCollection)
}

func (crud SnapshotCrud) FindSnapshots(filters bson.M) (subscribers []models.Snapshot, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := crud.Db.Find(SnapshotCollection, filters, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSnapshots(cursor)
	return
}

func (crud SnapshotCrud) FindLatestSnapshot(matchId string) (snapshot models.Snapshot, err error) {
	findOptions := &options.FindOneOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	dto, err := crud.Db.FindOne(SnapshotCollection,
		bson.M{"quiz_id": matchId}, findOptions)
	if err != nil {
		return
	}
	bson.Unmarshal(dto, &snapshot)
	return
}
