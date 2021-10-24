package core

import "errors"

type QuestionService struct {
	db DB
}

func (repo QuestionService) FindQuestionForMatch(match Game) (question Question, err error) {
	index, err := repo.db.FindIndex()
	if err != nil {
		err = errors.New(Err_IndexNotPresent)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), 1)
	questions, err := repo.db.FindQuestionsFromIndexes(indexes, int64(1))
	if len(questions) != 1 || err != nil {
		err = errors.New(Err_QuestionNotPresent)
		return
	}

	question = questions[0]
	return
}
