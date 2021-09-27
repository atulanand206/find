package core

import (
	"context"
	"encoding/json"
	"net/http"

	mg "go.mongodb.org/mongo-driver/mongo"
)

func DecodeIndex(document *mg.SingleResult) (index Index, err error) {
	if err = document.Decode(&index); err != nil {
		return
	}
	return
}

func DecodeMatch(document *mg.SingleResult) (game Game, err error) {
	if err = document.Decode(&game); err != nil {
		return
	}
	return
}

func DecodePlayer(document *mg.SingleResult) (player Player, err error) {
	if err = document.Decode(&player); err != nil {
		return
	}
	return
}

func DecodePlayerJsonString(content string) (player Player, err error) {
	if err = json.Unmarshal([]byte(content), &player); err != nil {
		return
	}
	return
}

func DecodeAnswer(document *mg.SingleResult) (answer Answer, err error) {
	if err = document.Decode(&answer); err != nil {
		return
	}
	return
}

func DecodeIndexes(cursor *mg.Cursor) (indexes []Index, err error) {
	for cursor.Next(context.Background()) {
		var index Index
		err = cursor.Decode(&index)
		if err != nil {
			return
		}
		indexes = append(indexes, index)
	}
	return
}

func DecodeQuestions(cursor *mg.Cursor) (questions []Question, err error) {
	for cursor.Next(context.Background()) {
		var question Question
		err = cursor.Decode(&question)
		if err != nil {
			return
		}
		questions = append(questions, question)
	}
	return
}

func DecodeAddQuestionRequest(r *http.Request) (request AddQuestionRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeNextQuestionRequest(r *http.Request) (request NextQuestionRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeFindAnswerRequest(r *http.Request) (request FindAnswerRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeEnterGameRequest(r *http.Request) (request EnterGameRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeStartGameRequest(r *http.Request) (request StartGameRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeWebSocketRequest(input []byte) (request WebsocketMessage, err error) {
	err = json.Unmarshal(input, &request)
	return
}
