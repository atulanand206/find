package core

import (
	"encoding/json"
	"errors"
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

	if err = Controller.questionService.AddQuestion(requestBody.Tag, requestBody.Questions); err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
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

func HandlerTestAPI(w http.ResponseWriter, r *http.Request) {
	matches, err := Controller.matchService.crud.FindActiveMatches()
	if err != nil {
		err = errors.New(Err_MatchNotPresent)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(matches)
}
