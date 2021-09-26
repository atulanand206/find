package core

func CreateDeckDtos() (indexs []Index, questions []Question, answers []Answer, err error) {
	indexsWrapper, err := ReadIndexes()
	if err != nil {
		return
	}
	indexs = make([]Index, 0)
	questions = make([]Question, 0)
	answers = make([]Answer, 0)
	for _, index := range indexsWrapper.Indexes {
		idx := InitNewIndex(index)
		indexs = append(indexs, idx)
		newQuestions := ReadQuestionsForIndex(idx)
		for _, newQuestion := range newQuestions {
			question := InitNewQuestion(idx, newQuestion)
			questions = append(questions, question)
			answers = append(answers, InitNewAnswer(question, newQuestion))
		}
	}
	return
}

func MapSansTags(tags []string) (tagsMap map[string]bool) {
	tagsMap = make(map[string]bool)
	for _, tag := range tags {
		tagsMap[tag] = true
	}
	return
}
