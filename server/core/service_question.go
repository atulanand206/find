package core

import "errors"

type QuestionService struct {
	crud QuestionCrud
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
