package core

import (
	"github.com/atulanand206/go-mongo"
	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionCrud struct {
	db DB
}

func (crud QuestionCrud) SeedIndexes(indexes []Index) (err error) {
	var indexesDto []interface{}
	for _, t := range indexes {
		indexesDto = append(indexesDto, t)
	}
	_, err = mongo.WriteMany(Database, IndexCollection, indexesDto)
	return err
}

func (crud QuestionCrud) CreateQuestion(question Question) (err error) {
	requestDto, err := mongo.Document(question)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, QuestionCollection, *requestDto)
	return
}

func (crud QuestionCrud) SeedQuestions(questions []Question) (err error) {
	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	_, err = mongo.WriteMany(Database, QuestionCollection, questionsDto)
	return err
}

func (crud QuestionCrud) CreateAnswer(answer Answer) (err error) {
	requestDto, err := mongo.Document(answer)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, AnswerCollection, *requestDto)
	return
}

func (crud QuestionCrud) SeedAnswers(answers []Answer) (err error) {
	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	_, err = mongo.WriteMany(Database, AnswerCollection, answersDto)
	return err
}

func (crud QuestionCrud) FindQuestion(questionId string) (question Question, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": questionId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	question, err = DecodeQuestion(dto)
	return
}

func (crud QuestionCrud) FindIndexForTag(tag string) (index Index, err error) {
	dto := mongo.FindOne(Database, IndexCollection, bson.M{"tag": tag}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	index, err = DecodeIndex(dto)
	return
}

func (crud QuestionCrud) FindAnswer(questionId string) (answer Answer, err error) {
	dto := mongo.FindOne(Database, AnswerCollection, bson.M{"question_id": questionId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	answer, err = DecodeAnswer(dto)
	return
}

func (crud QuestionCrud) FindIndex() (indexes []Index, err error) {
	cursor, err := mongo.Find(Database, IndexCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	indexes, err = DecodeIndexes(cursor)
	return
}

func (crud QuestionCrud) FindQuestionsForIndex(index Index, limit int64) (questions []Question, err error) {
	cursor, err := mongo.Find(Database, QuestionCollection,
		bson.M{"tag": index.Id}, &options.FindOptions{Limit: pointer.Int64(limit)})
	if err != nil {
		return
	}
	questions, err = DecodeQuestions(cursor)
	return
}

func (crud QuestionCrud) FindQuestionsFromIndexes(indexes []Index, limit int64) (questions []Question, err error) {
	questions = make([]Question, 0)
	for _, indx := range indexes {
		indxQues, er := crud.FindQuestionsForIndex(indx, limit)
		if er != nil {
			return
		}
		questions = append(questions, indxQues...)
	}
	return
}
