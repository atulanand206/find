package services

import (
	e "errors"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/utils"
)

type QuestionService struct {
	Crud db.QuestionCrud
}

func (service QuestionService) AddQuestion(tag string, newQuestions []models.NewQuestion) (err error) {
	index := models.InitNewIndex(tag)
	questions := make([]models.Question, 0)
	answers := make([]models.Answer, 0)

	for _, newQuestion := range newQuestions {
		question := models.InitNewQuestion(index, newQuestion)
		questions = append(questions, question)

		answers = append(answers, models.InitNewAnswer(question, newQuestion))
	}

	if err = service.Crud.SeedIndexes([]models.Index{index}); err != nil {
		err = e.New(errors.Err_IndexNotSeeded)
		return
	}

	if err = service.Crud.SeedQuestions(questions); err != nil {
		err = e.New(errors.Err_QuestionsNotSeeded)
		return
	}

	if err = service.Crud.SeedAnswers(answers); err != nil {
		err = e.New(errors.Err_AnswersNotSeeded)
		return
	}
	return
}

func (service QuestionService) FindQuestionForMatch(match models.Game) (question models.Question, err error) {
	index, err := service.Crud.FindIndex()
	if err != nil {
		err = e.New(errors.Err_IndexNotPresent)
		return
	}

	indexes := utils.FilterIndex(index, utils.MapSansTags(match.Tags), 1)
	if len(indexes) == 0 {
		err = e.New(errors.Err_QuestionNotPresent)
		return
	}

	questions, err := service.Crud.FindQuestionsFromIndexes(indexes, int64(1))
	if len(questions) != 1 || err != nil {
		err = e.New(errors.Err_QuestionNotPresent)
		return
	}

	question = questions[0]
	return
}
