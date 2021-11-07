package db

import (
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MatchCrud struct {
	Db DBConn
}

func (crud MatchCrud) CreateMatch(match models.Game) error {
	return crud.Db.Create(match, MatchCollection)
}

func (crud MatchCrud) FindMatch(matchId string) (match models.Game, err error) {
	dto := crud.Db.FindOne(MatchCollection, bson.M{"_id": matchId}, &options.FindOneOptions{})
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
	cursor, err := crud.Db.Find(MatchCollection,
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
	res, err := crud.Db.Update(MatchCollection, bson.M{"_id": match.Id}, match)
	updated = int(res.ModifiedCount) == 1
	return
}
