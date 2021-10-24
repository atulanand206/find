package core

import (
	"fmt"
)

func (hub *Hub) Handle(msg WebsocketMessage, client *Client) (res WebsocketMessage, targets map[string]bool, err error) {
	fmt.Println(msg)
	request, err := DecodeRequestJsonString(msg.Content)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	err = Controller.validator.ValidateRequest(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}
	switch request.Action {
	case BEGIN.String():
		res, targets, err = client.OnBegin(request)
	case ACTIVE.String():
		res, targets, err = client.OnActive()
	case SPECS.String():
		res, targets, err = client.OnCreate(request)
	case JOIN.String():
		res, targets, err = client.OnJoin(request)
	case WATCH.String():
		res, targets, err = client.OnWatch(request)
	case REFRESH.String():
		res, targets, err = client.OnRefresh(request)
	case START.String():
		res, targets, err = client.OnStart(request)
	case HINT.String():
		res, targets, err = client.OnHint(request)
	case RIGHT.String():
		res, targets, err = client.OnRight(request)
	case NEXT.String():
		res, targets, err = client.OnNext(request)
	case PASS.String():
		res, targets, err = client.OnPass(request)
	case SCORE.String():
		res, targets, err = client.OnScore(request)
	}
	fmt.Println(res)
	fmt.Println()
	return
}

func (client *Client) OnActive() (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateActiveQuizResponse()
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.selfResponse(client.playerId, S_ACTIVE, response)
	return
}

func (client *Client) OnBegin(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	player, err := Controller.GenerateBeginGameResponse(request.Person)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	client.setPlayerId(player.Id)
	res, targets = Controller.subscriberService.selfResponse(request.Person.Id, S_PLAYER, player)
	return
}

func (client *Client) OnCreate(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateCreateGameResponse(request.Person, request.Specs)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.selfResponse(request.Person.Id, S_JOIN, response)
	return
}

type Enter func(request Request) (res GameResponse, err error)

func (client *Client) requestToJoin(request Request, enter Enter) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := enter(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessageFailure()
		return
	}

	res, targets = Controller.subscriberService.joinResponse(request.QuizId, response)
	return
}

func (client *Client) OnJoin(content Request) (res WebsocketMessage, targets map[string]bool, err error) {
	return client.requestToJoin(content, func(request Request) (res GameResponse, err error) {
		return Controller.GenerateEnterGameResponse(request)
	})
}

func (client *Client) OnWatch(content Request) (res WebsocketMessage, targets map[string]bool, err error) {
	return client.requestToJoin(content, func(request Request) (res GameResponse, err error) {
		return Controller.GenerateWatchGameResponse(request)
	})
}

func (client *Client) OnRefresh(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateFullMatchResponse(request.QuizId)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.selfResponse(request.Person.Id, S_GAME, response)
	return
}

func (client *Client) OnStart(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateStartGameResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.quizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnHint(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateQuestionHintResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.quizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnRight(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateQuestionAnswerResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.quizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnNext(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateNextQuestionResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.quizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnPass(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GeneratePassQuestionResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.quizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnScore(request Request) (res WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateScoreResponse(request)
	if err != nil {
		res = MessageCreator.InitWebSocketMessage(Failure, err.Error())
		return
	}

	res, targets = Controller.subscriberService.selfResponse(request.QuizId, S_SCORE, response)
	return
}
