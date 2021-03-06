package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/comms"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/go-mongo"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

func Setup(t *testing.T) func(t *testing.T) {
	mongo.ConfigureMongoClient("mongodb://localhost:27017")
	db.Database = "binquiz-test"
	os.Setenv("CLIENT_SECRET", "aedsaddsad05442b3fe16f2e72d1d497bf14ea9cb")
	os.Setenv("REFRESH_CLIENT_SECRET", "ae05442b3fe16f2e72d1d497bf14ea9dsafdsadsacb")
	os.Setenv("TOKEN_EXPIRE_MINUTES", "2")
	os.Setenv("REFRESH_TOKEN_EXPIRE_MINUTES", "10")
	return func(t *testing.T) {
	}
}

func TestPlayer() models.Player {
	playerId, _ := gonanoid.New(16)
	playerName, _ := gonanoid.New(16)
	email, _ := gonanoid.New(16)
	player := models.Player{
		Id:    playerId,
		Name:  playerName,
		Email: email,
	}
	return player
}

func TestRandomString() string {
	id, _ := gonanoid.New(16)
	return id
}

func TestTeams(quizId string, count int) []models.Team {
	teams := make([]models.Team, count)
	for i := 0; i < count; i++ {
		teams[i] = TestTeam(quizId)
	}
	return teams
}

func TestTeam(quizId string) models.Team {
	teamId, _ := gonanoid.New(16)
	teamName, _ := gonanoid.New(16)
	return models.Team{
		Id:     teamId,
		QuizId: quizId,
		Name:   teamName,
		Score:  0,
	}
}

func TestSpecs() models.Specs {
	gameName, _ := gonanoid.New(16)
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
	playerId, _ := gonanoid.New(16)
	playerName, _ := gonanoid.New(16)
	email, _ := gonanoid.New(16)
	quizmaster := models.Player{
		Id:    playerId,
		Name:  playerName,
		Email: email,
	}
	specs := TestSpecs()
	game := models.InitNewMatch(quizmaster, specs)
	return game
}

func TestIndex() models.Index {
	return models.Index{
		Id:  TestRandomString(),
		Tag: TestRandomString(),
	}
}

func TestIndexes(count int) []models.Index {
	indexes := make([]models.Index, count)
	for i := 0; i < count; i++ {
		indexes[i] = TestIndex()
	}
	return indexes
}

func TestNewQuestion() models.NewQuestion {
	return models.NewQuestion{
		Statements: []string{
			TestRandomString(),
			TestRandomString(),
			TestRandomString(),
		},
		Answer: []string{
			TestRandomString(),
			TestRandomString(),
			TestRandomString(),
		},
	}
}

func TestQuestion(tag string, question models.NewQuestion) models.Question {
	return models.Question{
		Id:         TestRandomString(),
		Tag:        tag,
		Statements: question.Statements,
	}
}

func TestAnswer(questionId string, answer []string) models.Answer {
	return models.Answer{
		Id:         TestRandomString(),
		QuestionId: questionId,
		Answer:     answer,
		Hint: []string{
			TestRandomString(),
			TestRandomString(),
			TestRandomString(),
		},
	}
}

func TestSubscriber() models.Subscriber {
	return models.Subscriber{
		Tag:      TestRandomString(),
		PlayerId: TestRandomString(),
		Role:     actions.PLAYER.String(),
		Active:   true,
	}
}

func TestSnapshot(quizId string) models.Snapshot {
	return models.InitialSnapshot(quizId, []models.TeamRoster{})
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
	hub := comms.NewHub(services.Init(db.NewMockDb(true)))
	assert.NotNil(t, hub, "Hub should not be nil")
	go hub.Run()
	return hub
}

func NewClient(t *testing.T, hub *comms.Hub) *comms.Client {
	client := comms.NewClient(hub, nil)
	assert.NotNil(t, client, "Client should not be nil")
	hub.Register <- client
	return client
}

func TestMessage(action actions.Action, content string) models.WebsocketMessage {
	return models.WebsocketMessageCreator{}.InitWebSocketMessage(action, content)
}

func TestBeginMessage(player models.Player) models.WebsocketMessage {
	request := models.Request{
		Action: "BEGIN",
		Person: player,
	}
	return TestMessage(actions.BEGIN, string(comms.SerializeMessage(request)))
}

func ClientX(hub *comms.Hub) *comms.Client {
	for client := range hub.Clients {
		return client
	}
	return nil
}
