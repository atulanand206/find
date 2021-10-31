package core

import (
	"net/http"
	"os"
)

var (
	filePath  = "G:\\binge\\binge-questions\\%s.json"
	indexFile = "index"
)

func HandlerAddQuestion(w http.ResponseWriter, r *http.Request) {
	requestBody, err := DecodeAddQuestionRequest(r)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	index := InitNewIndex(requestBody.Tag)
	questions := make([]Question, 0)
	answers := make([]Answer, 0)

	for _, newQuestion := range requestBody.Questions {
		question := InitNewQuestion(index, newQuestion)
		questions = append(questions, question)

		answers = append(answers, InitNewAnswer(question, newQuestion))
	}

	if err = Controller.questionService.crud.SeedIndexes([]Index{index}); err != nil {
		http.Error(w, Err_IndexNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = Controller.questionService.crud.SeedQuestions(questions); err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = Controller.questionService.crud.SeedAnswers(answers); err != nil {
		http.Error(w, Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}

func HandlerSeedQuestions(w http.ResponseWriter, r *http.Request) {
	filePath = os.Getenv("SEED_FILES_PATH")

	var err error
	if err = Db.DropCollections(); err != nil {
		http.Error(w, Err_CollectionsNotDropped, http.StatusInternalServerError)
		return
	}

	if err = Db.CreateCollections(); err != nil {
		http.Error(w, Err_CollectionsNotCreated, http.StatusInternalServerError)
		return
	}

	indexes, questions, answers, err := CreateDeckDtos()
	if err != nil {
		http.Error(w, Err_DeckDtosNotCreated, http.StatusInternalServerError)
		return
	}

	if err = Controller.questionService.crud.SeedIndexes(indexes); err != nil {
		http.Error(w, Err_IndexNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = Controller.questionService.crud.SeedQuestions(questions); err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = Controller.questionService.crud.SeedAnswers(answers); err != nil {
		http.Error(w, Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}
