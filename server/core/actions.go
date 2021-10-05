package core

import (
	"encoding/json"
	"fmt"
)

func HandleWSMessage(msg WebsocketMessage) (res WebsocketMessage, err error) {
	fmt.Println(msg)
	switch msg.Action {
	case ACTIVE.String():
		res, err = OnActive()
	case SPECS.String():
		res, err = OnCreate(msg.Content)
	case JOIN.String():
		res, err = OnJoin(msg.Content)
	case WATCH.String():
		res, err = OnWatch(msg.Content)
	case START.String():
		res, err = OnStart(msg.Content)
	case HINT.String():
		res, err = OnHint(msg.Content)
	case RIGHT.String():
		res, err = OnRight(msg.Content)
	case NEXT.String():
		res, err = OnNext(msg.Content)
	case PASS.String():
		res, err = OnPass(msg.Content)
	case SCORE.String():
		res, err = OnScore(msg.Content)
	}
	fmt.Println(res)
	fmt.Println()
	return
}

func OnBegin(content string) (res Player, err error) {
	request, err := DecodePlayerJsonString(content)
	if err != nil {
		// res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateBeginGameResponse(request)
	if err != nil {
		// res = InitWebSocketMessage(Failure, err.Error())
		return
	}
	res = response
	return
}

func OnActive() (res WebsocketMessage, err error) {
	response, err := GenerateActiveQuizResponse()
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_ACTIVE, response)
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

func OnHint(content string) (res WebsocketMessage, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateQuestionHintResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_HINT, response)
	return
}

func OnRight(content string) (res WebsocketMessage, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateQuestionAnswerResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_RIGHT, response)
	return
}

func OnNext(content string) (res WebsocketMessage, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateNextQuestionResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_NEXT, response)
	return
}

func OnPass(content string) (res WebsocketMessage, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GeneratePassQuestionResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_PASS, response)
	return
}

func OnScore(content string) (res WebsocketMessage, err error) {
	request, err := DecodeScoreRequestJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	response, err := GenerateScoreResponse(request)
	if err != nil {
		res = InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_SCORE, response)
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
