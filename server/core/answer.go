package core

import (
	"encoding/json"
	"net/http"
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
	request, err := DecodeFindAnswerRequest(r)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	answer, err := FindAnswer(request.QuestionId)
	if err != nil {
		http.Error(w, Err_AnswerNotPresent, http.StatusInternalServerError)
		return
	}

	response := InitFindAnswerResponse(request.QuestionId, answer)
	json.NewEncoder(w).Encode(response)
}
