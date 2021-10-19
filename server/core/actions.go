package core

import (
	"encoding/json"
	"fmt"
)

func (hub *Hub) Handle(msg WebsocketMessage, client *Client) (res WebsocketMessage, targets map[string]bool, err error) {
	fmt.Println(msg)
	switch msg.Action {
	case BEGIN.String():
		res, targets, err = client.OnBegin(msg.Content)
	case ACTIVE.String():
		res, targets, err = client.OnActive()
	case SPECS.String():
		res, targets, err = client.OnCreate(msg.Content)
	case JOIN.String():
		res, targets, err = client.OnJoin(msg.Content)
	case WATCH.String():
		res, targets, err = client.OnWatch(msg.Content)
	case REFRESH.String():
		res, targets, err = client.OnRefresh(msg.Content)
	case START.String():
		res, targets, err = OnStart(msg.Content)
	case HINT.String():
		res, targets, err = OnHint(msg.Content)
	case RIGHT.String():
		res, targets, err = OnRight(msg.Content)
	case NEXT.String():
		res, targets, err = OnNext(msg.Content)
	case PASS.String():
		res, targets, err = OnPass(msg.Content)
	case SCORE.String():
		res, targets, err = OnScore(msg.Content)
	}
	fmt.Println(res)
	fmt.Println()
	return
}

func (client *Client) OnBegin(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodePlayerJsonString(content)
	targets = make(map[string]bool)
	targets[request.Id] = true

	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	player, err := Controller.GenerateBeginGameResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	client.setPlayerId(player.Id)
	resBytes, er := json.Marshal(player)
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
	}

	res = MessageCreator.InitWebSocketMessage(S_PLAYER, string(resBytes))

	return
}

func (client *Client) OnActive() (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateActiveQuizResponse()
	targets = make(map[string]bool)
	targets[client.playerId] = true

	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_ACTIVE, response)
	return
}

func (client *Client) OnCreate(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeCreateGameRequestJsonString(content)
	targets = make(map[string]bool)
	targets[request.Quizmaster.Id] = true

	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateCreateGameResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res = WebSocketsResponse(S_GAME, response)
	return
}

type Enter func(request EnterGameRequest) (res Snapshot, err error)

func (client *Client) requestToJoin(content string, enter Enter) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeEnterGameRequestJsonString(content)
	targets = make(map[string]bool)
	targets[request.Person.Id] = true

	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := enter(request)

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_GAME, response)
	return
}

func (client *Client) OnJoin(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	return client.requestToJoin(content, func(request EnterGameRequest) (res Snapshot, err error) {
		return Controller.GenerateEnterGameResponse(request)
	})
}

func (client *Client) OnWatch(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	return client.requestToJoin(content, func(request EnterGameRequest) (res Snapshot, err error) {
		return Controller.GenerateWatchGameResponse(request)
	})
}

func (client *Client) OnRefresh(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeEnterGameRequestJsonString(content)
	targets = make(map[string]bool)
	targets[request.Person.Id] = true

	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateFullMatchResponse(request)
	res = WebSocketsResponse(S_GAME, response)
	return
}

func OnStart(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeStartGameRequestJsonString(content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateStartGameResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_START, response)
	return
}

func OnHint(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateQuestionHintResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_HINT, response)
	return
}

func OnRight(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateQuestionAnswerResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_RIGHT, response)
	return
}

func OnNext(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateNextQuestionResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_NEXT, response)
	return
}

func OnPass(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeGameSnapRequestJsonString(content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GeneratePassQuestionResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_PASS, response)
	return
}

func OnScore(content string) (res WebsocketMessage, targets map[string]bool, err error) {
	request, err := DecodeScoreRequestJsonString(content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	response, err := Controller.GenerateScoreResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	subscribers, er := Controller.subscriberService.FindSubscribersForTag([]string{request.QuizId})
	if er != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		targets[subscriber.PlayerId] = true
	}

	res = WebSocketsResponse(S_SCORE, response)
	return
}

func WebSocketsResponse(action Action, v interface{}) (res WebsocketMessage) {
	resBytes, err := json.Marshal(v)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	res = MessageCreator.InitWebSocketMessage(action, string(resBytes))
	return
}
