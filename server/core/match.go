package core

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	PlayersCount = 2
)

type (
	Player struct {
		Id    string `json:"id" bson:"_id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
	}

	Game struct {
		Id         string   `json:"id" bson:"_id"`
		Players    []Player `json:"players" bson:"players"`
		QuizMaster Player   `json:"quizmaster" bson:"quizmaster"`
		Tags       []string `bson:"tags"`
	}

	EnterGameRequest struct {
		Person  Player `json:"person"`
		MatchId string `json:"match_id"`
	}

	StartGameRequest struct {
		MatchId string `json:"match_id"`
	}

	StartGameResponse struct {
		Match  Game       `json:"game"`
		Prompt []Question `json:"prompt"`
	}
)

func HandlerBeginGame(w http.ResponseWriter, r *http.Request) {
	requestBody, err := DecodeEnterGameRequest(r)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	player, err := FindPlayer(requestBody.Person.Email)

	if err != nil {
		player = InitNewPlayer(requestBody.Person)
		if err = CreatePlayer(player); err != nil {
			http.Error(w, Err_PlayerNotCreated, http.StatusInternalServerError)
			return
		}
	}

	match, err := FindMatch(requestBody.MatchId)
	for _, v := range match.Players {
		if v.Id == player.Id {
			return
		}
	}

	match = InitNewMatch(requestBody.Person)
	if err = CreateMatch(match); err != nil {
		http.Error(w, Err_MatchNotCreated, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(match)
}

func HandlerEnterGame(w http.ResponseWriter, r *http.Request) {
	requestBody, err := DecodeEnterGameRequest(r)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	player, err := FindPlayer(requestBody.Person.Email)

	if err != nil {
		player = InitNewPlayer(requestBody.Person)
		if err = CreatePlayer(player); err != nil {
			http.Error(w, Err_PlayerNotCreated, http.StatusInternalServerError)
			return
		}
	}

	match, err := FindMatch(requestBody.MatchId)
	for _, v := range match.Players {
		if v.Id == player.Id {
			return
		}
	}

	if match.QuizMaster.Id == requestBody.Person.Id {
		err = errors.New(Err_QuizmasterCantPlay)
		http.Error(w, Err_QuizmasterCantPlay, http.StatusInternalServerError)
		return
	}

	if err = UpdateMatchPlayer(match, requestBody.Person); err != nil {
		http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(match)
}

func HandlerStartGame(w http.ResponseWriter, r *http.Request) {
	requestBody, err := DecodeStartGameRequest(r)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	match, err := FindMatch(requestBody.MatchId)
	if err != nil {
		http.Error(w, Err_MatchNotPresent, http.StatusInternalServerError)
		return
	}

	if len(match.Players) != PlayersCount {
		http.Error(w, Err_WaitingForPlayers, http.StatusInternalServerError)
		return
	}

	index, err := FindIndex()
	if err != nil {
		http.Error(w, Err_IndexNotPresent, http.StatusInternalServerError)
		return
	}

	indexes := FilterIndex(index, MapSansTags(match.Tags), 1)

	questions, err := FindQuestionsFromIndexes(indexes, int64(1))
	if err != nil {
		http.Error(w, Err_QuestionNotPresent, http.StatusInternalServerError)
		return
	}

	if err = UpdateMatchQuestions(match, questions); err != nil {
		http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
	}

	response := InitStartGameResponse(match, questions)
	json.NewEncoder(w).Encode(response)
}
