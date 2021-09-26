package core

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	Answer struct {
		Id         string `json:"id" bson:"_id"`
		QuestionId string `json:"question_id" bson:"question_id"`
		Answer     string `json:"answer" bson:"answer"`
	}

	FindAnswerRequest struct {
		QuestionId string `json:"question_id"`
	}

	FindAnswerResponse struct {
		QuestionId string `json:"question_id"`
		Answer     string `json:"answer"`
	}
)

func HandlerFindAnswer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody FindAnswerRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	dto := mongo.FindOne(Database, AnswerCollection, bson.M{"question_id": requestBody.QuestionId})
	err = dto.Err()
	if err != nil {
		http.Error(w, Err_AnswerNotPresent, http.StatusInternalServerError)
		return
	}
	answer, err := DecodeAnswer(dto)
	var response FindAnswerResponse
	response.QuestionId = requestBody.QuestionId
	response.Answer = answer.Answer
	json.NewEncoder(w).Encode(response)
}
