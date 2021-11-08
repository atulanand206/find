package comms

import (
	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
)

func (hub *Hub) Handle(msg models.WebsocketMessage, client *Client) (res models.WebsocketMessage, targets map[string]bool, err error) {
	request, err := db.DecodeRequestJsonString(msg.Content)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	err = hub.Controller.Validator.ValidateRequest(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}
	switch request.Action {
	case actions.BEGIN.String():
		res, targets, err = hub.OnBegin(request, client)
	case actions.SPECS.String():
		res, targets, err = hub.OnCreate(request)
	case actions.JOIN.String():
		res, targets, err = hub.OnJoin(request)
	case actions.WATCH.String():
		res, targets, err = hub.OnWatch(request)
	case actions.REFRESH.String():
		res, targets, err = hub.OnRefresh(request)
	case actions.START.String():
		res, targets, err = hub.OnStart(request)
	case actions.HINT.String():
		res, targets, err = hub.OnHint(request)
	case actions.RIGHT.String():
		res, targets, err = hub.OnRight(request)
	case actions.NEXT.String():
		res, targets, err = hub.OnNext(request)
	case actions.PASS.String():
		res, targets, err = hub.OnPass(request)
	case actions.SCORE.String():
		res, targets, err = hub.OnScore(request)
	}
	return
}

func (hub *Hub) OnBegin(request models.Request, client *Client) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateBeginGameResponse(request.Person)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	client.SetPlayerId(response.Player.Id)
	res, targets = hub.Controller.SubscriberService.SelfResponse(response.Player.Id, actions.S_PLAYER, response)
	return
}

func (hub *Hub) OnCreate(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateCreateGameResponse(request.Person, request.Specs)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.SelfResponse(request.Person.Id, actions.S_JOIN, response)
	return
}

type Enter func(request models.Request) (res models.GameResponse, err error)

func (hub *Hub) requestToJoin(request models.Request, enter Enter) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := enter(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessageFailure()
		return
	}

	res, targets = hub.Controller.SubscriberService.SelfResponse(request.Person.Id, actions.S_JOIN, response)
	return
}

func (hub *Hub) OnJoin(content models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	return hub.requestToJoin(content, func(request models.Request) (res models.GameResponse, err error) {
		return hub.Controller.GenerateEnterGameResponse(request)
	})
}

func (hub *Hub) OnWatch(content models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	return hub.requestToJoin(content, func(request models.Request) (res models.GameResponse, err error) {
		return hub.Controller.GenerateWatchGameResponse(request)
	})
}

func (hub *Hub) OnRefresh(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateFullMatchResponse(request.QuizId)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (hub *Hub) OnStart(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateStartGameResponse(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (hub *Hub) OnHint(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateQuestionHintResponse(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (hub *Hub) OnRight(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateQuestionAnswerResponse(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (hub *Hub) OnNext(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateNextQuestionResponse(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (hub *Hub) OnPass(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GeneratePassQuestionResponse(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (hub *Hub) OnScore(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := hub.Controller.GenerateScoreResponse(request)
	if err != nil {
		res = hub.Controller.Creators.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = hub.Controller.SubscriberService.SelfResponse(request.Person.Id, actions.S_SCORE, response)
	return
}
