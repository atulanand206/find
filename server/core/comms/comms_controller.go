package comms

import (
	"errors"
	"fmt"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
)

func (hub *Hub) Handle(msg models.WebsocketMessage, client *Client) (res models.WebsocketMessage, targets map[string]bool, err error) {
	fmt.Println(msg)
	request, err := db.DecodeRequestJsonString(msg.Content)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	err = Controller.Validator.ValidateRequest(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessageFailure()
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
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	client.setPlayerId(response.Player.Id)
	res, targets = client.SelfResponse(response.Player.Id, actions.S_PLAYER, response)
	return
}

func (client *Client) OnCreate(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateCreateGameResponse(request.Person, request.Specs)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.SelfResponse(request.Person.Id, actions.S_JOIN, response)
	return
}

type Enter func(request models.Request) (res models.GameResponse, err error)

func (client *Client) requestToJoin(request models.Request, enter Enter) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := enter(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessageFailure()
		return
	}

	res, targets = client.SelfResponse(request.Person.Id, actions.S_JOIN, response)
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
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnStart(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateStartGameResponse(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnHint(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateQuestionHintResponse(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnRight(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateQuestionAnswerResponse(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnNext(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateNextQuestionResponse(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnPass(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GeneratePassQuestionResponse(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.QuizResponse(request.Action, request.QuizId, response)
	return
}

func (client *Client) OnScore(request models.Request) (res models.WebsocketMessage, targets map[string]bool, err error) {
	response, err := Controller.GenerateScoreResponse(request)
	if err != nil {
		res = Controller.MessageCreator.InitWebSocketMessage(actions.Failure, err.Error())
		return
	}

	res, targets = client.SelfResponse(request.Person.Id, actions.S_SCORE, response)
	return
}

func (client *Client) SelfResponse(playerId string, action actions.Action, response interface{}) (res models.WebsocketMessage, targets map[string]bool) {
	res = Controller.MessageCreator.WebSocketsResponse(action, response)
	targets = Controller.TargetService.TargetSelf(playerId)
	return
}

func (client *Client) QuizResponse(action string, quizId string, response models.Snapshot) (res models.WebsocketMessage, targets map[string]bool) {
	subscribers, er := Controller.SubscriberService.FindSubscribersForTag([]string{quizId})
	if er != nil {
		res = Controller.MessageCreator.InitWebSocketMessageFailure()
		return
	}
	targets = Controller.TargetService.TargetQuiz(quizId, subscribers)
	res = Controller.MessageCreator.WebSocketsResponse(actions.S_GAME, &models.SnapshotResponse{Action: action, Snapshot: response})
	return
}

func (client *Client) FindOrCreateSubscriber(tag string, audience models.Player, role actions.Role) (subscriber models.Subscriber, err error) {
	subscriber, err = Controller.SubscriberService.Crud.FindSubscriberForTagAndPlayerId(tag, audience.Id)
	if err != nil {
		subscriber = Controller.InstanceCreator.InitSubscriber(tag, audience, role.String())
		err = Controller.SubscriberService.Crud.CreateSubscriber(subscriber)
		if err != nil {
			err = errors.New(fmt.Sprint(Controller.ErrorCreator.SubscriberNotCreated(subscriber)))
		}
	}
	return
}

func (client *Client) SubscribeAndRespond(match models.Game, player models.Player, snapshot models.Snapshot, role actions.Role) (response models.GameResponse, err error) {
	_, err = client.FindOrCreateSubscriber(match.Id, player, role)
	if err != nil {
		return
	}

	response = models.GameResponse{Quiz: match, Snapshot: snapshot, Role: role.String()}
	return
}
