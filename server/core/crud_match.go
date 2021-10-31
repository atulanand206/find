package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MatchCrud struct {
	db DB
}

func (crud MatchCrud) CreateMatch(match Game) error {
	return crud.db.Create(match, MatchCollection)
}

func (crud MatchCrud) FindMatch(matchId string) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	match, err = DecodeMatch(dto)
	return
}

func (crud MatchCrud) FindActiveMatches() (matches []Game, err error) {
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

func (crud MatchCrud) FindActiveTeamMatches(teams []Team) (matches []Game, err error) {
	matchIds := make([]string, 0)
	for _, v := range teams {
		matchIds = append(matchIds, v.QuizId)
	}
	findOptions := &options.FindOptions{}
	cursor, err := mongo.Find(Database, MatchCollection,
		bson.M{"_id": bson.M{"$in": matchIds}, "active": true}, findOptions)
	if err != nil {
		return
	}
	matches, err = DecodeMatches(cursor)
	return
}

func (crud MatchCrud) UpdateMatchQuestions(match Game, question Question) error {
	match.Tags = append(match.Tags, question.Tag)
	return crud.UpdateMatch(match)
}

func (crud MatchCrud) UpdateMatch(match Game) (err error) {
	requestDto, err := mongo.Document(match)
	if err != nil {
		return
	}
	_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": match.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	return
}
