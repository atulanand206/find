package comms

import (
	"fmt"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
)

func (hub *Hub) Handle(msg models.WebsocketMessage, client *Client) (res models.WebsocketMessage, targets map[string]bool, err error) {
	fmt.Println(msg)
	request, err := db.DecodeRequestJsonString(msg.Content)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	err = Controller.Validator.ValidateRequest(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	switch request.Action {
	case actions.BEGIN.String():
		res, targets, err = client.OnBegin(request)
	case actions.SPECS.String():
		res, targets, err = client.OnCreate(request)
	case actions.JOIN.String():
		res, targets, err = client.OnJoin(request)
	case actions.WATCH.String():
		res, targets, err = client.OnWatch(request)
	case actions.REFRESH.String():
		res, targets, err = client.OnRefresh(request)
	case actions.START.String():
		res, targets, err = client.OnStart(request)
	case actions.HINT.String():
		res, targets, err = client.OnHint(request)
	case actions.RIGHT.String():
		res, targets, err = client.OnRight(request)
	case actions.NEXT.String():
		res, targets, err = client.OnNext(request)
	case actions.PASS.String():
		res, targets, err = client.OnPass(request)
	case actions.SCORE.String():
		res, targets, err = client.OnScore(request)
	}
	fmt.Println(targets)
	fmt.Println(res)
	fmt.Println()
	return
}

func (client *Client) OnBegin(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateBeginGameResponse(request.Person)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	client.SetPlayerId(response.Player.Id)
	res, targets = Controller.SubscriberService.SelfResponse(response.Player.Id, actions.S_PLAYER, response)
	return
}

func (client *Client) OnCreate(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateCreateGameResponse(request.Person, request.Specs)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.SelfResponse(request.Person.Id, actions.S_JOIN, response)
	return
}

type Enter func(request models.Request) (res models.GameResponse, err error)

func (client *Client) requestToJoin(request models.Request, enter Enter) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := enter(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}

	res, targets = Controller.SubscriberService.SelfResponse(request.Person.Id, actions.S_JOIN, response)
	return
}

func (client *Client) OnJoin(content models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	return client.requestToJoin(content, func(request models.Request) (res models.GameResponse, err error) {
		return Controller.GenerateEnterGameResponse(request)
	})
}

func (client *Client) OnWatch(content models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	return client.requestToJoin(content, func(request models.Request) (res models.GameResponse, err error) {
		return Controller.GenerateWatchGameResponse(request)
	})
}

func (client *Client) OnRefresh(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateFullMatchResponse(request.QuizId)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnStart(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateStartGameResponse(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnHint(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateQuestionHintResponse(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnRight(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateQuestionAnswerResponse(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnNext(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateNextQuestionResponse(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnPass(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GeneratePassQuestionResponse(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnScore(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateScoreResponse(request)
	if err != nil {
		res = Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = Controller.SubscriberService.SelfResponse(request.Person.Id, actions.S_SCORE, response)
	return
}
