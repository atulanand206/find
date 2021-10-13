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
	Err_RequestNotDecoded = "Request can't be decoded."

	Err_CollectionsNotCreated  = "Collections can't be created."
	Err_DeckDtosNotCreated     = "Deck dtos can't be created."
	Err_MatchRequestNotCreated = "New match request can't be created."
	Err_MatchNotCreated        = "New match can't be created."
	Err_TeamNotCreated         = "Team can't be created."
	Err_PlayerNotCreated       = "Player can't be created."
	Err_QuestionNotCreated     = "Question can't be created."
	Err_AnswerNotCreated       = "Answer can't be created."
	Err_SnapshotNotCreated     = "Snapshot can't be created."

	Err_IndexNotSeeded     = "Index can't be seeded."
	Err_QuestionsNotSeeded = "Questions can't be seeded."
	Err_AnswersNotSeeded   = "Answers can't be seeded."

	Err_IndexNotPresent       = "Index does not exist."
	Err_MatchNotPresent       = "Match with the given info does not exist."
	Err_SnapshotNotPresent    = "Snapshot with the given info does not exist."
	Err_PlayerNotPresent      = "Player with the given info does not exist."
	Err_QuestionNotPresent    = "Question does not exist."
	Err_AnswerNotPresent      = "Answer does not exist."
	Err_TeamNotPresent        = "Team does not exist."
	Err_TeamPlayersNotPresent = "Team players does not exist."

	Err_MatchNotUpdated = "Match can't be updated."
	Err_TeamNotUpdated  = "Team can't be updated."

	Err_CollectionsNotDropped = "Collections can't be dropped."

	Err_QuizmasterCantPlay     = "Quizmaster can't join the match as a player."
	Err_WaitingForPlayers      = "Waiting for more players to join."
	Err_QuestionsNotLeft       = "No remaining question in the quiz."
	Err_TeamsNotPresentInMatch = "Teams not present in the match."
	Err_PlayersFullInTeam      = "Players already full in the team."

	Err_SocketRequestFailed = "Sockets request failed by the server."
	Err_PlayerAlreadyInGame = "Player already present in the game."

	Err_SubscriberNotPresent = "Match with the given info does not exist."
)

func (creator ErrorMessageCreator) NotCreated(entity string, data interface{}) (errorMsg ErrorMessage) {
	errorMsg.code = http.StatusNotAcceptable
	errorMsg.msg = fmt.Sprintf("Unable to create %s %v", entity, data)
	return
}

func (creator ErrorMessageCreator) SubscriberNotCreated(v interface{}) (errorMsg ErrorMessage) {
	errorMsg = creator.NotCreated("subscriber", v)
	return
}
