package db

import (
	"context"
	"encoding/json"

	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
)

func DecodePlayerJsonString(content string) (player models.Player, err error) {
	if err = json.Unmarshal([]byte(content), &player); err != nil {
		return
	}
	return
}

func DecodePermissions(cursor []bson.Raw) (scopes []models.Permission, err error) {
	for _, doc := range cursor {
		var scope models.Permission
		err = bson.Unmarshal(doc, &scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}

func DecodeMatches(cursor []bson.Raw) (games []models.Game, err error) {
	for _, doc := range cursor {
		var game models.Game
		err = bson.Unmarshal(doc, &game)
		if err != nil {
			return
		}
		games = append(games, game)
	}
	return
}

func DecodeIndexes(cursor []bson.Raw) (indexes []models.Index, err error) {
	for _, doc := range cursor {
		var index models.Index
		err = bson.Unmarshal(doc, &index)
		if err != nil {
			return
		}
		indexes = append(indexes, index)
	}
	return
}

func DecodeQuestions(cursor []bson.Raw) (questions []models.Question, err error) {
	for _, doc := range cursor {
		var question models.Question
		err = bson.Unmarshal(doc, &question)
		if err != nil {
			return
		}
		questions = append(questions, question)
	}
	return
}

func DecodeSnapshots(cursor []bson.Raw) (snapshots []models.Snapshot, err error) {
	for _, doc := range cursor {
		var snapshot models.Snapshot
		err = bson.Unmarshal(doc, &snapshot)
		if err != nil {
			return
		}
		snapshots = append(snapshots, snapshot)
	}
	return
}

func DecodeSubscribers(cursor []bson.Raw) (subscribers []models.Subscriber, err error) {
	for _, doc := range cursor {
		var subscriber models.Subscriber
		err = bson.Unmarshal(doc, &subscriber)
		if err != nil {
			return
		}
		subscribers = append(subscribers, subscriber)
	}
	return
}

func DecodeTeams(cursor []bson.Raw) (teams []models.Team, err error) {
	for _, doc := range cursor {
		var team models.Team
		err = bson.Unmarshal(doc, &team)
		if err != nil {
			return
		}
		teams = append(teams, team)
	}
	return
}

func DecodePlayers(cursor []bson.Raw) (players []models.Player, err error) {
	for _, doc := range cursor {
		var player models.Player
		err = bson.Unmarshal(doc, &player)
		if err != nil {
			return
		}
		players = append(players, player)
	}
	return
}

func DecodeRaw(cursor *mg.Cursor) (documents []bson.Raw, err error) {
	for cursor.Next(context.Background()) {
		documents = append(documents, cursor.Current)
	}
	return
}
