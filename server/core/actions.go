package core

import (
	"encoding/json"
	"fmt"
)

func HandleWSMessage(msg WebsocketMessage) (res WebsocketMessage, err error) {
	fmt.Println(msg.Content)
	switch msg.Action {
	case BEGIN.String():
		res, err = OnBegin(msg.Content)
	case SPECS.String():
		res, err = OnCreate(msg.Content)
	case JOIN.String():
		res, err = OnJoin(msg.Content)
	case WATCH.String():
		res, err = OnWatch(msg.Content)
	case REVEAL.String():
		res, err = OnReveal(msg.Content)
	}
	fmt.Println(res)
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

	res = WebSocketsResponse(S_PLAYER, response)
	return
}

func OnCreate(content string) (res WebsocketMessage, err error) {
	request, err := DecodeCreateGameRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateCreateGameResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_GAME, response)
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

	res = WebSocketsResponse(S_GAME, response)
	return
}

func OnWatch(content string) (res WebsocketMessage, err error) {
	request, err := DecodeEnterGameRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateWatchGameResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_GAME, response)
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

	res = WebSocketsResponse(S_START, response)
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
