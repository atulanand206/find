package core

import (
	"context"

	"github.com/atulanand206/go-mongo"
	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindMatch(matchId string) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId})
	match, err = DecodeMatch(dto)
	return
}

func FindIndex() (indexes []Index, err error) {
	cursor, err := mongo.Find(Database, IndexCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	indexes, err = DecodeIndexes(cursor)
	return
}

func FindQuestionsFromIndex(index Index, limit int64) (questions []Question, err error) {
	cursor, err := mongo.Find(Database, QuestionCollection,
		bson.M{"tag": index.Id}, &options.FindOptions{Limit: pointer.Int64(limit)})
	if err != nil {
		return
	}
	for cursor.Next(context.Background()) {
		var question Question
		cursor.Decode(&question)
		questions = append(questions, question)
	}
	return
}
