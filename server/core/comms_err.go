package core

import (
	"fmt"
	"net/http"
)

type ErrorMessageCreator struct{}

type ErrorMessage struct {
	msg  string
	code int
}

const (
	Err_RequestNotDecoded = "request can't be decoded"

	Err_CollectionsNotCreated  = "collections can't be created"
	Err_DeckDtosNotCreated     = "deck dtos can't be created"
	Err_MatchRequestNotCreated = "new match request can't be created"
	Err_MatchNotCreated        = "new match can't be created"
	Err_TeamNotCreated         = "team can't be created"
	Err_PlayerNotCreated       = "player can't be created"
	Err_QuestionNotCreated     = "question can't be created"
	Err_AnswerNotCreated       = "answer can't be created"
	Err_SnapshotNotCreated     = "snapshot can't be created"

	Err_IndexNotSeeded     = "index can't be seeded"
	Err_QuestionsNotSeeded = "questions can't be seeded"
	Err_AnswersNotSeeded   = "answers can't be seeded"

	Err_IndexNotPresent       = "index does not exist"
	Err_MatchNotPresent       = "match with the given info does not exist"
	Err_SnapshotNotPresent    = "snapshot with the given info does not exist"
	Err_PlayerNotPresent      = "player with the given info does not exist"
	Err_QuestionNotPresent    = "question does not exist"
	Err_AnswerNotPresent      = "answer does not exist"
	Err_TeamNotPresent        = "team does not exist"
	Err_TeamPlayersNotPresent = "team players does not exist"

	Err_MatchNotUpdated = "match can't be updated"
	Err_TeamNotUpdated  = "team can't be updated"

	Err_CollectionsNotDropped = "collections can't be dropped"

	Err_QuizmasterCantPlay     = "quizmaster can't join the match as a player"
	Err_WaitingForPlayers      = "waiting for more players to join"
	Err_QuestionsNotLeft       = "no remaining question in the quiz"
	Err_TeamsNotPresentInMatch = "teams not present in the match"
	Err_PlayersFullInTeam      = "players already full in the team"

	Err_SocketRequestFailed = "sockets request failed by the server"
	Err_PlayerAlreadyInGame = "player already present in the game"

	Err_SubscriberNotPresent = "match with the given info does not exist"
)

func (creator ErrorMessageCreator) NotCreated(entity string, data interface{}) (errorMsg ErrorMessage) {
	errorMsg.code = http.StatusNotAcceptable
	errorMsg.msg = fmt.Sprintf("unable to create %s %v", entity, data)
	return
}

func (creator ErrorMessageCreator) SubscriberNotCreated(v interface{}) (errorMsg ErrorMessage) {
	errorMsg = creator.NotCreated("subscriber", v)
	return
}

func (creator ErrorMessageCreator) RequestInvalid() (errorMsg ErrorMessage) {
	errorMsg.msg = "request invalid"
	errorMsg.code = http.StatusBadRequest
	return
}
