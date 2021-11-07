package db

import (
	"github.com/atulanand206/find/server/core/models"
	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionCrud struct {
	Db DB
}

func (crud QuestionCrud) SeedIndexes(indexes []models.Index) (err error) {
	var indexesDto []interface{}
	for _, t := range indexes {
		indexesDto = append(indexesDto, t)
	}
	return crud.Db.CreateMany(indexesDto, IndexCollection)
}

func (crud QuestionCrud) CreateQuestion(question models.Question) (err error) {
	return crud.Db.Create(question, QuestionCollection)
}

func (crud QuestionCrud) SeedQuestions(questions []models.Question) (err error) {
	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	return crud.Db.CreateMany(questionsDto, QuestionCollection)
}

func (crud QuestionCrud) CreateAnswer(answer models.Answer) (err error) {
	return crud.Db.Create(answer, AnswerCollection)
}

func (crud QuestionCrud) SeedAnswers(answers []models.Answer) (err error) {
	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	return crud.Db.CreateMany(answersDto, AnswerCollection)
}

func (crud QuestionCrud) FindQuestion(questionId string) (question models.Question, err error) {
	dto, err := crud.Db.FindOne(QuestionCollection, bson.M{"_id": questionId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	err = bson.Unmarshal(dto, &question)
	return
}

func (crud QuestionCrud) FindIndexForTag(tag string) (index models.Index, err error) {
	dto, err := crud.Db.FindOne(IndexCollection, bson.M{"tag": tag}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	err = bson.Unmarshal(dto, &index)
	return
}

func (crud QuestionCrud) FindAnswer(questionId string) (answer models.Answer, err error) {
	dto, err := crud.Db.FindOne(AnswerCollection, bson.M{"question_id": questionId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	err = bson.Unmarshal(dto, &answer)
	return
}

func (crud QuestionCrud) FindIndex() (indexes []models.Index, err error) {
	cursor, err := crud.Db.Find(IndexCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	indexes, err = DecodeIndexes(cursor)
	return
}

func (crud QuestionCrud) FindQuestionsForIndex(index models.Index, limit int64) (questions []models.Question, err error) {
	cursor, err := crud.Db.Find(QuestionCollection,
		bson.M{"tag": index.Id}, &options.FindOptions{Limit: pointer.Int64(limit)})
	if err != nil {
		return
	}
	questions, err = DecodeQuestions(cursor)
	return
}

func (crud QuestionCrud) FindQuestionsFromIndexes(indexes []models.Index, limit int64) (questions []models.Question, err error) {
	questions = make([]models.Question, 0)
	for _, indx := range indexes {
		indxQues, er := crud.FindQuestionsForIndex(indx, limit)
		if er != nil {
			return
		}
		questions = append(questions, indxQues...)
	}
	return
}
