package core

import "github.com/google/uuid"

func InitNewIndex(tag string) (index Index) {
	index.Tag = tag
	index.Id = uuid.New().String()
	return
}

func InitNewQuestion(index Index, newQuestion NewQuestion) (question Question) {
	question.Statements = newQuestion.Statements
	question.Tag = index.Id
	question.Id = uuid.New().String()
	return
}

func InitNewAnswer(question Question, newQuestion NewQuestion) (answer Answer) {
	answer.Answer = newQuestion.Answer
	answer.QuestionId = question.Id
	answer.Id = uuid.New().String()
	return
}
