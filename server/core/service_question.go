package core

import "errors"

type QuestionService struct {
	crud QuestionCrud
}

func (service QuestionService) AddQuestion(tag string, newQuestions []NewQuestion) (err error) {
	index := InitNewIndex(tag)
	questions := make([]Question, 0)
	answers := make([]Answer, 0)

	for _, newQuestion := range newQuestions {
		question := InitNewQuestion(index, newQuestion)
		questions = append(questions, question)

		answers = append(answers, InitNewAnswer(question, newQuestion))
	}

	if err = service.crud.SeedIndexes([]Index{index}); err != nil {
		err = errors.New(Err_IndexNotSeeded)
		return
	}

	if err = service.crud.SeedQuestions(questions); err != nil {
		err = errors.New(Err_QuestionsNotSeeded)
		return
	}

	if err = service.crud.SeedAnswers(answers); err != nil {
		err = errors.New(Err_AnswersNotSeeded)
		return
	}
	return
}

func (service QuestionService) FindQuestionForMatch(match Game) (question Question, err error) {
	index, err := service.crud.FindIndex()
	if err != nil {
		err = errors.New(Err_IndexNotPresent)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), 1)
	if len(indexes) == 0 {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	questions, err := service.crud.FindQuestionsFromIndexes(indexes, int64(1))
	if len(questions) != 1 || err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	question = questions[0]
	return
}
