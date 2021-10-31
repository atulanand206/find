package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SnapshotCrud struct {
	db DB
}

func (crud SnapshotCrud) CreateSnapshot(snapshot Snapshot) error {
	return crud.db.Create(snapshot, SnapshotCollection)
}

func (crud SnapshotCrud) FindSnapshotsForMatch(matchId string) (snapshots []Snapshot, err error) {
	cursor, err := mongo.Find(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId}, &options.FindOptions{})
	if err != nil {
		return
	}
	snapshots, err = DecodeSnapshots(cursor)
	return
}

func (crud SnapshotCrud) FindSnapshotsForQuestion(quizId string, questionId string, eventType string) (snapshots []Snapshot, err error) {
	cursor, err := mongo.Find(Database, SnapshotCollection,
		bson.M{"quiz_id": quizId, "question_id": questionId, "event_type": eventType}, &options.FindOptions{})
	if err != nil {
		return
	}
	snapshots, err = DecodeSnapshots(cursor)
	return
}

func (crud SnapshotCrud) FindLatestSnapshot(matchId string) (snapshot Snapshot, err error) {
	findOptions := &options.FindOneOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	dto := mongo.FindOne(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId}, findOptions)
	if err = dto.Err(); err != nil {
		return
	}
	snapshot, err = DecodeSnapshot(dto)
	return
}

func (crud SnapshotCrud) FindQuestionSnapshots(matchId string, questionId string) (snapshot []Snapshot, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId, "question_id": questionId}, findOptions)
	if err != nil {
		return
	}
	snapshot, err = DecodeSnapshots(cursor)
	return
}
