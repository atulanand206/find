package db

import (
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SnapshotCrud struct {
	Db DB
}

func (crud SnapshotCrud) CreateSnapshot(snapshot models.Snapshot) error {
	return crud.Db.Create(snapshot, SnapshotCollection)
}

func (crud SnapshotCrud) FindSnapshotsForMatch(matchId string) (snapshots []models.Snapshot, err error) {
	cursor, err := crud.Db.Find(SnapshotCollection,
		bson.M{"quiz_id": matchId}, &options.FindOptions{})
	if err != nil {
		return
	}
	snapshots, err = DecodeSnapshots(cursor)
	return
}

func (crud SnapshotCrud) FindSnapshotsForQuestion(quizId string, questionId string, eventType string) (snapshots []models.Snapshot, err error) {
	cursor, err := crud.Db.Find(SnapshotCollection,
		bson.M{"quiz_id": quizId, "question_id": questionId, "event_type": eventType}, &options.FindOptions{})
	if err != nil {
		return
	}
	snapshots, err = DecodeSnapshots(cursor)
	return
}

func (crud SnapshotCrud) FindLatestSnapshot(matchId string) (snapshot models.Snapshot, err error) {
	findOptions := &options.FindOneOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	dto := crud.Db.FindOne(SnapshotCollection,
		bson.M{"quiz_id": matchId}, findOptions)
	if err = dto.Err(); err != nil {
		return
	}
	snapshot, err = DecodeSnapshot(dto)
	return
}

func (crud SnapshotCrud) FindQuestionSnapshots(matchId string, questionId string) (snapshot []models.Snapshot, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := crud.Db.Find(SnapshotCollection,
		bson.M{"quiz_id": matchId, "question_id": questionId}, findOptions)
	if err != nil {
		return
	}
	snapshot, err = DecodeSnapshots(cursor)
	return
}
