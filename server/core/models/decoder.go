package models

import (
	"encoding/json"
	"net/http"
	"strings"
)

func DecodeAddQuestionRequest(r *http.Request) (request AddQuestionRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeLoginResponseJsonString(content string) (request LoginResponse, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeEnterGameRequestJsonString(content string) (request EnterGameRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeCreateGameRequestJsonString(content string) (request CreateGameRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeGameResponseJsonString(content string) (request GameResponse, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeStartGameRequestJsonString(content string) (request StartGameRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeGameSnapRequestJsonString(content string) (request GameSnapRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeScoreRequestJsonString(content string) (request ScoreRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeWebSocketRequest(input []byte) (request WebsocketMessage, err error) {
	err = json.Unmarshal(input, &request)
	request.Content = strings.Replace(request.Content, "\\", "", -1)
	request.Content = strings.TrimLeft(request.Content, "\"")
	request.Content = strings.TrimRight(request.Content, "\"")
	return
}
