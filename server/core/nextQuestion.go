package core

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type (
	NextQuestionRequest struct {
		MatchId string `json:"match_id"`
		Limit   int    `json:"limit"`
		Types   int    `json:"types"`
	}
)

func HandlerNextQuestion(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestBody NextQuestionRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	match, err := FindMatch(requestBody.MatchId)
	if err != nil {
		http.Error(w, Err_MatchNotPresent, http.StatusInternalServerError)
		return
	}

	questions, err := FindQuestions(requestBody.Limit, MapSansTags(match.Tags), requestBody.Types)
	if err != nil {
		http.Error(w, Err_QuestionNotPresent, http.StatusInternalServerError)
		return
	}

	err = UpdateMatch(match, questions)
	if err != nil {
		http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func FindQuestions(limit int, sans map[string]bool, types int) (questions []Question, err error) {
	index, err := FindIndex()
	if err != nil {
		return
	}
	indexes := FilterIndex(index, sans, types)
	questions, err = FindQuestionsFromIndexes(indexes, int64(limit))
	if err != nil {
		return
	}
	return
}

func FilterIndex(input []Index, sans map[string]bool, types int) (indexes []Index) {
	indxs := make([]Index, 0)
	for _, indx := range input {
		if !sans[indx.Id] {
			indxs = append(indxs, indx)
		}
	}
	mp := make(map[Index]bool)
	rand.Seed(time.Now().UnixNano())
	indexCount := len(indxs)
	for {
		if len(mp) < types {
			tag := indxs[rand.Intn(indexCount)]
			mp[tag] = true
		} else {
			break
		}
	}
	indexes = make([]Index, 0)
	for key := range mp {
		indexes = append(indexes, key)
	}
	return indexes
}

func FindQuestionsFromIndexes(indexes []Index, limit int64) (questions []Question, err error) {
	questions = make([]Question, 0)
	for _, indx := range indexes {
		indxQues, er := FindQuestionsFromIndex(indx, limit)
		if er != nil {
			return
		}
		questions = append(questions, indxQues...)
	}
	return
}
