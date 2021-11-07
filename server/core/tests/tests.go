package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/comms"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/go-mongo"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
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

func RunningHubWithClients(t *testing.T, n int) *comms.Hub {
	hub := RunningHub(t)
	for i := 0; i < n; i++ {
		NewClient(t, hub)
	}
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, n, len(hub.Clients), fmt.Sprintf("Hub should have %d clients", n))
	return hub
}

func RunningHub(t *testing.T) *comms.Hub {
	hub := comms.NewHub()
	assert.NotNil(t, hub, "Hub should not be nil")
	go hub.Run()
	return hub
}

func NewClient(t *testing.T, hub *comms.Hub) *comms.Client {
	client := comms.NewClient(hub, nil)
	testPlayerId, _ := gonanoid.New(10)
	client.SetPlayerId(testPlayerId)
	assert.NotNil(t, client, "Client should not be nil")
	hub.Register <- client
	return client
}

func TestMessage(action actions.Action, content string) models.WebsocketMessage {
	return models.WebsocketMessageCreator{}.InitWebSocketMessage(action, content)
}

func ClientX(hub *comms.Hub) *comms.Client {
	for client := range hub.Clients {
		return client
	}
	return nil
}
