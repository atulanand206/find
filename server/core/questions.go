package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/atulanand206/go-mongo"
	"github.com/google/uuid"
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

	AddQuestionResponse struct {
		QuestionId string `json:"question_id"`
		AnswerId   string `json:"answer_id"`
	}
)

var (
	filePath  = "G:\\binge\\binge-questions\\%s.json"
	indexFile = "index"
)

func HandlerAddQuestion(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody NewQuestion
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	var question Question
	question.Id = uuid.New().String()
	question.Statements = requestBody.Statements
	requestDto, _ := mongo.Document(question)
	insertedQuestionId, err := mongo.Write(Database, QuestionCollection, *requestDto)
	if err != nil {
		http.Error(w, Err_QuestionNotCreated, http.StatusInternalServerError)
		return
	}
	var answer Answer
	answer.Id = uuid.New().String()
	answer.QuestionId = fmt.Sprint(insertedQuestionId.InsertedID)
	answer.Answer = requestBody.Answer
	requestDto, _ = mongo.Document(answer)
	insertedAnswerId, err := mongo.Write(Database, AnswerCollection, *requestDto)
	if err != nil {
		http.Error(w, Err_AnswerNotCreated, http.StatusInternalServerError)
		return
	}
	var response AddQuestionResponse
	response.QuestionId = fmt.Sprint(insertedQuestionId.InsertedID)
	response.AnswerId = fmt.Sprint(insertedAnswerId.InsertedID)
	json.NewEncoder(w).Encode(response)
}

func HandlerSeedQuestions(w http.ResponseWriter, r *http.Request) {
	filePath = os.Getenv("SEED_FILES_PATH")

	err := DropQuestionsCollections()
	if err != nil {
		http.Error(w, Err_CollectionsNotDropped, http.StatusInternalServerError)
		return
	}

	err = CreateQuestionsCollections()
	if err != nil {
		http.Error(w, Err_CollectionsNotCreated, http.StatusInternalServerError)
		return
	}

	indexes, questions, answers, err := CreateDeckDtos()
	if err != nil {
		http.Error(w, Err_DeckDtosNotCreated, http.StatusInternalServerError)
		return
	}

	err = SeedIndexes(indexes)
	if err != nil {
		http.Error(w, Err_IndexNotSeeded, http.StatusInternalServerError)
		return
	}

	err = SeedQuestions(questions)
	if err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	err = SeedAnswers(answers)
	if err != nil {
		http.Error(w, Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}
