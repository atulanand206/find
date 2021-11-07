package db

import (
	"context"
	"encoding/json"

	"github.com/atulanand206/find/server/core/models"
	mg "go.mongodb.org/mongo-driver/mongo"
)

func DecodeIndex(document *mg.SingleResult) (index models.Index, err error) {
	if err = document.Decode(&index); err != nil {
		return
	}
	return
}

func DecodeMatch(document *mg.SingleResult) (game models.Game, err error) {
	if err = document.Decode(&game); err != nil {
		return
	}
	return
}

func DecodeQuestion(document *mg.SingleResult) (question models.Question, err error) {
	if err = document.Decode(&question); err != nil {
		return
	}
	return
}

func DecodePlayer(document *mg.SingleResult) (player models.Player, err error) {
	if err = document.Decode(&player); err != nil {
		return
	}
	return
}

func DecodePlayerJsonString(content string) (player models.Player, err error) {
	if err = json.Unmarshal([]byte(content), &player); err != nil {
		return
	}
	return
}

func DecodeRequestJsonString(content string) (request models.Request, err error) {
	if err = json.Unmarshal([]byte(content), &request); err != nil {
		return
	}
	return
}

func DecodeAnswer(document *mg.SingleResult) (answer models.Answer, err error) {
	if err = document.Decode(&answer); err != nil {
		return
	}
	return
}

func DecodeSnapshot(document *mg.SingleResult) (snapshot models.Snapshot, err error) {
	if err = document.Decode(&snapshot); err != nil {
		return
	}
	return
}

func DecodeSubscriber(document *mg.SingleResult) (subscriber models.Subscriber, err error) {
	if err = document.Decode(&subscriber); err != nil {
		return
	}
	return
}

func DecodePermissions(cursor *mg.Cursor) (scopes []models.Permission, err error) {
	for cursor.Next(context.Background()) {
		var scope models.Permission
		err = cursor.Decode(&scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}

func DecodeMatches(cursor *mg.Cursor) (games []models.Game, err error) {
	for cursor.Next(context.Background()) {
		var game models.Game
		err = cursor.Decode(&game)
		if err != nil {
			return
		}
		games = append(games, game)
	}
	return
}

func DecodeIndexes(cursor *mg.Cursor) (indexes []models.Index, err error) {
	for cursor.Next(context.Background()) {
		var index models.Index
		err = cursor.Decode(&index)
		if err != nil {
			return
		}
		indexes = append(indexes, index)
	}
	return
}

func DecodeQuestions(cursor *mg.Cursor) (questions []models.Question, err error) {
	for cursor.Next(context.Background()) {
		var question models.Question
		err = cursor.Decode(&question)
		if err != nil {
			return
		}
		questions = append(questions, question)
	}
	return
}

func DecodeSnapshots(cursor *mg.Cursor) (snapshots []models.Snapshot, err error) {
	for cursor.Next(context.Background()) {
		var snapshot models.Snapshot
		err = cursor.Decode(&snapshot)
		if err != nil {
			return
		}
		snapshots = append(snapshots, snapshot)
	}
	return
}

func DecodeSubscribers(cursor *mg.Cursor) (subscribers []models.Subscriber, err error) {
	for cursor.Next(context.Background()) {
		var subscriber models.Subscriber
		err = cursor.Decode(&subscriber)
		if err != nil {
			return
		}
		subscribers = append(subscribers, subscriber)
	}
	return
}

func DecodeTeams(cursor *mg.Cursor) (teams []models.Team, err error) {
	for cursor.Next(context.Background()) {
		var team models.Team
		err = cursor.Decode(&team)
		if err != nil {
			return
		}
		teams = append(teams, team)
	}
	return
}

func DecodePlayers(cursor *mg.Cursor) (players []models.Player, err error) {
	for cursor.Next(context.Background()) {
		var player models.Player
		err = cursor.Decode(&player)
		if err != nil {
			return
		}
		players = append(players, player)
	}
	return
}

func DecodeTeamPlayers(cursor *mg.Cursor) (teamPlayers []models.Subscriber, err error) {
	for cursor.Next(context.Background()) {
		var teamPlayer models.Subscriber
		err = cursor.Decode(&teamPlayer)
		if err != nil {
			return
		}
		teamPlayers = append(teamPlayers, teamPlayer)
	}
	return
}
