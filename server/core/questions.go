package core

import (
	"encoding/json"
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

	index, err := FindIndexForTag(requestBody.Tag)
	if err != nil {
		http.Error(w, Err_IndexNotPresent, http.StatusInternalServerError)
		return
	}

	question := InitNewQuestion(index, requestBody.Question)
	if err = CreateQuestion(question); err != nil {
		http.Error(w, Err_QuestionNotCreated, http.StatusInternalServerError)
		return
	}

	answer := InitNewAnswer(question, requestBody.Question)
	if err = CreateAnswer(answer); err != nil {
		http.Error(w, Err_AnswerNotCreated, http.StatusInternalServerError)
		return
	}

	response := InitAddQuestionResponse(question, answer)
	json.NewEncoder(w).Encode(response)
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

	if err = SeedIndexes(indexes); err != nil {
		http.Error(w, Err_IndexNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = SeedQuestions(questions); err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = SeedAnswers(answers); err != nil {
		http.Error(w, Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}
