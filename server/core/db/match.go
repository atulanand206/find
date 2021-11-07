package db

import (
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MatchCrud struct {
	Db DB
}

func (crud MatchCrud) CreateMatch(match models.Game) error {
	return crud.Db.Create(match, MatchCollection)
}

func (crud MatchCrud) FindMatch(matchId string) (match models.Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	match, err = DecodeMatch(dto)
	return
}

func (crud MatchCrud) FindActiveMatches() (matches []models.Game, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "quizmaster.name", Value: 1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, MatchCollection,
		bson.M{"active": true}, findOptions)

	if err != nil {
		return
	}
	matches, err = DecodeMatches(cursor)
	return
}

func (crud MatchCrud) UpdateMatchQuestions(match models.Game, question models.Question) (bool, error) {
	match.Tags = append(match.Tags, question.Tag)
	return crud.UpdateMatch(match)
}

func (crud MatchCrud) UpdateMatch(match models.Game) (updated bool, err error) {
	requestDto, err := mongo.Document(match)
	if err != nil {
		return
	}
	res, err := mongo.Update(Database, MatchCollection, bson.M{"_id": match.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	updated = int(res.ModifiedCount) == 1
	return
}
