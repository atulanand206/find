package core

import (
	"encoding/json"
)

func HandleWSMessage(msg WebsocketMessage) (res WebsocketMessage, err error) {
	switch msg.Action {
	case Begin:
		res, err = OnBegin(msg.Content)
	case Join:
		res, err = OnJoin(msg.Content)
	case Start:
		res, err = OnStart(msg.Content)
	case Next:
		res, err = OnNext(msg.Content)
	case Reveal:
		res, err = OnReveal(msg.Content)
	}
	return
}

func OnBegin(content string) (res WebsocketMessage, err error) {
	request, err := DecodePlayerJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateBeginGameResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_Game, response)
	return
}

func OnJoin(content string) (res WebsocketMessage, err error) {
	request, err := DecodeEnterGameRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateEnterGameResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_Game, response)
	return
}

func OnStart(content string) (res WebsocketMessage, err error) {
	request, err := DecodeStartGameRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateStartGameResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_Start, response)
	return
}

func OnNext(content string) (res WebsocketMessage, err error) {
	request, err := DecodeNextQuestionRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateNextQuestionResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_Question, response)
	return
}

func OnReveal(content string) (res WebsocketMessage, err error) {
	request, err := DecodeFindAnswerRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateFindAnswerResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_Answer, response)
	return
}

func WebSocketsResponse(action Action, v interface{}) (res WebsocketMessage) {
	resBytes, err := json.Marshal(v)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}
	res = InitWebSocketMessage(action, string(resBytes))
	return
}
