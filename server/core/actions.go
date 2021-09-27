package core

import "encoding/json"

func HandleWSMessage(msg WebsocketMessage) (res WebsocketMessage, err error) {
	switch msg.Action {
	case Begin:
		res, err = OnBegin(msg.Content)
	}
	return
}

func OnBegin(content string) (res WebsocketMessage, err error) {
	quizmaster, err := DecodePlayerJsonString(content)
	if err != nil {
		res = InitWebSocketMessageFailure()
		return
	}

	match := InitNewMatch(quizmaster)
	if err = CreateMatch(match); err != nil {
		res = InitWebSocketMessage(Failure, Err_MatchNotCreated)
		return
	}

	resBytes, err := json.Marshal(match)
	if err != nil {
		res = InitWebSocketMessage(Failure, Err_RequestNotDecoded)
		return
	}

	res = InitWebSocketMessage(S_Game, string(resBytes))
	return
}
