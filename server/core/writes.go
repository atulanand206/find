package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateQuestionsCollections() (err error) {
	err = mongo.CreateCollections(Database, []string{IndexCollection, QuestionCollection, AnswerCollection})
	return
}

func SeedIndexes(indexes []Index) error {
	var indexesDto []interface{}
	for _, t := range indexes {
		indexesDto = append(indexesDto, t)
	}
	_, err := mongo.WriteMany(Database, IndexCollection, indexesDto)
	return err
}

func SeedQuestions(questions []Question) error {
	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	_, err := mongo.WriteMany(Database, QuestionCollection, questionsDto)
	return err
}

func SeedAnswers(answers []Answer) error {
	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	_, err := mongo.WriteMany(Database, AnswerCollection, answersDto)
	return err
}

func UpdateMatch(match Game, questions []Question) (err error) {
	for _, question := range questions {
		match.Tags = append(match.Tags, question.Tag)
	}
	requestDto, _ := mongo.Document(match)
	_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": match.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	return
}

func DropQuestionsCollections() (err error) {
	err = mongo.DropCollections(Database, []string{IndexCollection, QuestionCollection, AnswerCollection})
	return
}
