package tests

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/go-mongo"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Setup(t *testing.T) func(t *testing.T) {
	mongo.ConfigureMongoClient("mongodb://localhost:27017")
	db.Database = "binquiz-test"
	db.MatchCollection = "matches"
	db.QuestionCollection = "questions"
	db.AnswerCollection = "answers"
	db.SnapshotCollection = "snapshots"
	db.TeamCollection = "teams"
	db.PlayerCollection = "players"
	db.IndexCollection = "indexes"
	db.SubscriberCollection = "subscribers"
	return func(t *testing.T) {
	}
}

func TestPlayer() models.Player {
	playerId, _ := gonanoid.New(8)
	playerName, _ := gonanoid.New(8)
	email, _ := gonanoid.New(8)
	player := models.Player{
		Id:    playerId,
		Name:  playerName,
		Email: email,
	}
	return player
}

func TestSpecs() models.Specs {
	gameName, _ := gonanoid.New(8)
	return models.Specs{
		Name:      gameName,
		Teams:     2,
		Players:   2,
		Questions: 10,
		Rounds:    2,
		Points:    10,
	}
}

func TestGame() models.Game {
	playerId, _ := gonanoid.New(8)
	playerName, _ := gonanoid.New(8)
	email, _ := gonanoid.New(8)
	quizmaster := models.Player{
		Id:    playerId,
		Name:  playerName,
		Email: email,
	}
	specs := TestSpecs()
	game := models.InitNewMatch(quizmaster, specs)
	return game
}
