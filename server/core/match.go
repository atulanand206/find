package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/atulanand206/go-mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	decoder := json.NewDecoder(r.Body)
	var requestBody EnterGameRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	player, err := CreateOrFindPlayer(w, requestBody.Person)
	if err != nil {
		return
	}
	match, err := CreateOrUpdateMatch(w, requestBody.MatchId, player)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(match)
}

func HandlerEnterGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody EnterGameRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	player, err := CreateOrFindPlayer(w, requestBody.Person)
	if err != nil {
		return
	}
	match, err := CreateOrUpdateMatch(w, requestBody.MatchId, player)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(match)
}

func CreateOrFindPlayer(w http.ResponseWriter, playerRequest Player) (player Player, err error) {
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": playerRequest.Email})
	err = dto.Err()
	if err == nil {
		player, err = DecodePlayer(dto)
		return
	}
	player = playerRequest
	player.Id = uuid.New().String()
	requestDto, _ := mongo.Document(player)
	if err != nil {
		_, err = mongo.Write(Database, PlayerCollection, *requestDto)
		if err != nil {
			http.Error(w, Err_PlayerNotCreated, http.StatusInternalServerError)
			return
		}
	}
	dto = mongo.FindOne(Database, PlayerCollection, bson.M{"email": player.Email})
	if err != nil {
		http.Error(w, Err_PlayerNotPresent, http.StatusInternalServerError)
		return
	}
	player, err = DecodePlayer(dto)
	return
}

func CreateOrUpdateMatch(w http.ResponseWriter, matchId string, person Player) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId})
	err = dto.Err()
	if err == nil {
		match, err = DecodeMatch(dto)
	}
	for _, v := range match.Players {
		if v.Id == person.Id {
			return
		}
	}
	if err != nil {
		match.Id = uuid.New().String()
		match.Players = make([]Player, 0)
		match.QuizMaster = person
		requestDto, _ := mongo.Document(match)
		insertMatchResult, er := mongo.Write(Database, MatchCollection, *requestDto)
		err = nil
		if er != nil {
			http.Error(w, Err_MatchNotCreated, http.StatusInternalServerError)
			return
		}
		matchId = fmt.Sprint(insertMatchResult.InsertedID)
	} else {
		if match.QuizMaster.Id == person.Id {
			err = errors.New(Err_QuizmasterCantPlay)
			http.Error(w, Err_QuizmasterCantPlay, http.StatusInternalServerError)
			return
		}
		match.Players = append(match.Players, person)
		requestDto, _ := mongo.Document(match)
		_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": matchId}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
		if err != nil {
			http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
			return
		}
	}
	match, err = FindMatch(matchId)
	return
}

func HandlerStartGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody StartGameRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": requestBody.MatchId})
	err = dto.Err()
	if err != nil {
		http.Error(w, Err_MatchNotPresent, http.StatusInternalServerError)
		return
	}
	match, err := DecodeMatch(dto)
	if len(match.Players) != PlayersCount {
		http.Error(w, Err_WaitingForPlayers, http.StatusInternalServerError)
		return
	}
	questions, err := FindQuestions(1, map[string]bool{}, 1)
	if err != nil {
		return
	}
	UpdateMatch(match, questions)
	var response StartGameResponse
	response.Match = match
	response.Prompt = questions
	json.NewEncoder(w).Encode(response)
}
