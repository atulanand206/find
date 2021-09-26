package core

import (
	"encoding/json"
	"net/http"
	"os"
)

type (
	Index struct {
		Id  string `json:"id" bson:"_id"`
		Tag string `json:"tag" bson:"tag"`
	}

	IndexWrapper struct {
		Indexes []Index `json:"indexes" bson:"indexes"`
	}

	IndexStore struct {
		Indexes []string `json:"indexes" bson:"indexes"`
	}

	QuestionBank struct {
		Questions []NewQuestion `json:"questions" bson:"questions"`
	}

	NewQuestion struct {
		Statements []string `json:"statements"`
		Answer     string   `json:"answer"`
	}

	Question struct {
		Id         string   `json:"id" bson:"_id"`
		Statements []string `json:"statements" bson:"statements"`
		Tag        string   `json:"-" bson:"tag"`
	}

	AddQuestionRequest struct {
		Question NewQuestion `json:"question"`
		Tag      string      `json:"tag"`
	}

	AddQuestionResponse struct {
		QuestionId string `json:"question_id"`
		AnswerId   string `json:"answer_id"`
	}

	NextQuestionRequest struct {
		MatchId string `json:"match_id"`
		Limit   int    `json:"limit"`
		Types   int    `json:"types"`
	}
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
	if err = DropQuestionsCollections(); err != nil {
		http.Error(w, Err_CollectionsNotDropped, http.StatusInternalServerError)
		return
	}

	if err = CreateQuestionsCollections(); err != nil {
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

func HandlerNextQuestion(w http.ResponseWriter, r *http.Request) {
	requestBody, err := DecodeNextQuestionRequest(r)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	match, err := FindMatch(requestBody.MatchId)
	if err != nil {
		http.Error(w, Err_MatchNotPresent, http.StatusInternalServerError)
		return
	}

	index, err := FindIndex()
	if err != nil {
		http.Error(w, Err_IndexNotPresent, http.StatusInternalServerError)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), requestBody.Types)

	questions, err := FindQuestionsFromIndexes(indexes, int64(requestBody.Limit))
	if err != nil {
		http.Error(w, Err_QuestionNotPresent, http.StatusInternalServerError)
		return
	}

	if err = UpdateMatchQuestions(match, questions); err != nil {
		http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}
