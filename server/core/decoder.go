package core

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	mg "go.mongodb.org/mongo-driver/mongo"
)

func DecodeIndex(document *mg.SingleResult) (index Index, err error) {
	if err = document.Decode(&index); err != nil {
		return
	}
	return
}

func DecodeMatch(document *mg.SingleResult) (game Game, err error) {
	if err = document.Decode(&game); err != nil {
		return
	}
	return
}

func DecodeQuestion(document *mg.SingleResult) (question Question, err error) {
	if err = document.Decode(&question); err != nil {
		return
	}
	return
}

func DecodePlayer(document *mg.SingleResult) (player Player, err error) {
	if err = document.Decode(&player); err != nil {
		return
	}
	return
}

func DecodePlayerJsonString(content string) (player Player, err error) {
	if err = json.Unmarshal([]byte(content), &player); err != nil {
		return
	}
	return
}

func DecodeAnswer(document *mg.SingleResult) (answer Answer, err error) {
	if err = document.Decode(&answer); err != nil {
		return
	}
	return
}

func DecodeSnapshot(document *mg.SingleResult) (snapshot Snapshot, err error) {
	if err = document.Decode(&snapshot); err != nil {
		return
	}
	return
}

func DecodeSubscriber(document *mg.SingleResult) (subscriber Subscriber, err error) {
	if err = document.Decode(&subscriber); err != nil {
		return
	}
	return
}

func DecodeMatches(cursor *mg.Cursor) (games []Game, err error) {
	for cursor.Next(context.Background()) {
		var game Game
		err = cursor.Decode(&game)
		if err != nil {
			return
		}
		games = append(games, game)
	}
	return
}

func DecodeIndexes(cursor *mg.Cursor) (indexes []Index, err error) {
	for cursor.Next(context.Background()) {
		var index Index
		err = cursor.Decode(&index)
		if err != nil {
			return
		}
		indexes = append(indexes, index)
	}
	return
}

func DecodeQuestions(cursor *mg.Cursor) (questions []Question, err error) {
	for cursor.Next(context.Background()) {
		var question Question
		err = cursor.Decode(&question)
		if err != nil {
			return
		}
		questions = append(questions, question)
	}
	return
}

func DecodeSnapshots(cursor *mg.Cursor) (snapshots []Snapshot, err error) {
	for cursor.Next(context.Background()) {
		var snapshot Snapshot
		err = cursor.Decode(&snapshot)
		if err != nil {
			return
		}
		snapshots = append(snapshots, snapshot)
	}
	return
}

func DecodeSubscribers(cursor *mg.Cursor) (subscribers []Subscriber, err error) {
	for cursor.Next(context.Background()) {
		var subscriber Subscriber
		err = cursor.Decode(&subscriber)
		if err != nil {
			return
		}
		subscribers = append(subscribers, subscriber)
	}
	return
}

func DecodeTeams(cursor *mg.Cursor) (teams []Team, err error) {
	for cursor.Next(context.Background()) {
		var team Team
		err = cursor.Decode(&team)
		if err != nil {
			return
		}
		teams = append(teams, team)
	}
	return
}

func DecodePlayers(cursor *mg.Cursor) (players []Player, err error) {
	for cursor.Next(context.Background()) {
		var player Player
		err = cursor.Decode(&player)
		if err != nil {
			return
		}
		players = append(players, player)
	}
	return
}

func DecodeTeamPlayers(cursor *mg.Cursor) (teamPlayers []Subscriber, err error) {
	for cursor.Next(context.Background()) {
		var teamPlayer Subscriber
		err = cursor.Decode(&teamPlayer)
		if err != nil {
			return
		}
		teamPlayers = append(teamPlayers, teamPlayer)
	}
	return
}

func DecodeAddQuestionRequest(r *http.Request) (request AddQuestionRequest, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	return
}

func DecodeEnterGameRequestJsonString(content string) (request EnterGameRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeCreateGameRequestJsonString(content string) (request CreateGameRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeStartGameRequestJsonString(content string) (request StartGameRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeGameSnapRequestJsonString(content string) (request GameSnapRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeScoreRequestJsonString(content string) (request ScoreRequest, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeWebSocketRequest(input []byte) (request WebsocketMessage, err error) {
	err = json.Unmarshal(input, &request)
	request.Content = strings.Replace(request.Content, "\\", "", -1)
	request.Content = strings.TrimLeft(request.Content, "\"")
	request.Content = strings.TrimRight(request.Content, "\"")
	return
}
