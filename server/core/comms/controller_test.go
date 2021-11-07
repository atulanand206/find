package comms_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/comms"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleMessages(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	t.Run("Handle a string", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 2)
		msg := tests.TestMessage(actions.BEGIN, "Hello")
		_, _, err := hub.Handle(msg, tests.ClientX(hub))
		assert.NotNil(t, err, "error must be present")
	})

	t.Run("Handle a request with valid action", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 2)
		request := models.Request{
			Action: "EMPTY",
		}
		msg := tests.TestMessage(actions.DROP, string(comms.SerializeMessage(request)))
		msg, _, err := hub.Handle(msg, tests.ClientX(hub))
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, "", msg.Action, "action must be empty")
	})

	t.Run("Handle an invalid begin request", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 2)
		request := models.Request{
			Action: "BEGIN",
			Person: models.Player{
				Name:  "",
				Email: "",
				Id:    "",
			},
		}
		msg := tests.TestMessage(actions.BEGIN, string(comms.SerializeMessage(request)))
		client := tests.ClientX(hub)
		res, _, _ := hub.Handle(msg, client)
		assert.Equal(t, "failure", res.Action)
	})

	t.Run("Handle login", func(t *testing.T) {
		hub := tests.RunningHub(t)
		client := tests.NewClient(t, hub)

		testPlayerId, _ := gonanoid.New(10)
		player := models.Player{
			Name:  "Atul Khan",
			Email: "atul@zoom.in",
			Id:    testPlayerId,
		}
		Login(t, player, client)
	})

	t.Run("Handle login and create quiz", func(t *testing.T) {
		hub := tests.RunningHub(t)
		testPlayerId, _ := gonanoid.New(10)
		client := tests.NewClient(t, hub)
		player := models.Player{
			Name:  "Atul Khan",
			Email: "atul@zoom.io",
			Id:    testPlayerId,
		}
		Login(t, player, client)
		specs := models.Specs{
			Name:      "Binquiz #1",
			Players:   1,
			Teams:     2,
			Questions: 10,
		}
		Create(t, client, player, specs)
	})

	t.Run("login with 3 players", func(t *testing.T) {
		hub := tests.RunningHub(t)
		client1 := tests.NewClient(t, hub)
		testPlayerId, _ := gonanoid.New(10)
		quizmaster := models.Player{
			Name:  "Tony Stark",
			Email: "tony@stark.io",
			Id:    testPlayerId,
		}
		Login(t, quizmaster, client1)
		client2 := tests.NewClient(t, hub)
		testPlayerId, _ = gonanoid.New(10)
		player1 := models.Player{
			Name:  "Zara Khan",
			Email: "zara@khan.io",
			Id:    testPlayerId,
		}
		Login(t, player1, client2)
		client3 := tests.NewClient(t, hub)
		testPlayerId, _ = gonanoid.New(10)
		player2 := models.Player{
			Name:  "Mark Broody",
			Email: "mark@broody.io",
			Id:    testPlayerId,
		}
		Login(t, player2, client3)
	})
}

func TestJoinGame(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	t.Run("create a quiz with 2 players", func(t *testing.T) {
		hub := tests.RunningHub(t)
		client1 := tests.NewClient(t, hub)
		testPlayerId, _ := gonanoid.New(10)
		quizmaster := models.Player{
			Name:  "Tony Stark",
			Email: "mark@stark.io",
			Id:    testPlayerId,
		}
		Login(t, quizmaster, client1)
		client2 := tests.NewClient(t, hub)
		testPlayerId, _ = gonanoid.New(10)
		player1 := models.Player{
			Name:  "Zara Khan",
			Email: "zareen@khan.io",
			Id:    testPlayerId,
		}
		Login(t, player1, client2)
		client3 := tests.NewClient(t, hub)
		testPlayerId, _ = gonanoid.New(10)
		player2 := models.Player{
			Name:  "Mark Broody",
			Email: "martin@broody.io",
			Id:    testPlayerId,
		}
		Login(t, player2, client3)

		specs := models.Specs{
			Name:      "Binquiz #1",
			Players:   1,
			Teams:     2,
			Questions: 10,
		}
		createRes := Create(t, client1, quizmaster, specs)
		Join(t, client2, player1, createRes.Quiz.Id)
		Join(t, client3, player2, createRes.Quiz.Id)
	})
}

func Login(t *testing.T, player models.Player, client *comms.Client) models.LoginResponse {
	request := models.Request{
		Action: "BEGIN",
		Person: player,
	}
	res, _, _ := client.Hub.Handle(tests.TestMessage(actions.BEGIN, string(comms.SerializeMessage(request))), client)
	assert.NotNil(t, res, "response must be present")
	assert.Equal(t, "S_PLAYER", res.Action)
	loginResponse, err := models.DecodeLoginResponseJsonString(res.Content)
	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, player.Id, loginResponse.Player.Id)
	assert.Equal(t, player.Name, loginResponse.Player.Name)
	return loginResponse
}

func Create(t *testing.T, client *comms.Client, player models.Player, specs models.Specs) models.GameResponse {
	request := models.Request{
		Action: "SPECS",
		Person: player,
		Specs:  specs,
	}
	msg := tests.TestMessage(actions.SPECS, string(comms.SerializeMessage(request)))
	res, _, _ := client.Hub.Handle(msg, client)
	assert.NotNil(t, res, "response must be present")
	assert.Equal(t, "S_JOIN", res.Action)
	gameResponse, err := models.DecodeGameResponseJsonString(res.Content)
	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, player.Id, gameResponse.Quiz.QuizMaster.Id)
	assert.Equal(t, player.Name, gameResponse.Quiz.QuizMaster.Name)
	assert.Equal(t, "QUIZMASTER", gameResponse.Role)
	assert.Equal(t, specs.Name, gameResponse.Quiz.Specs.Name)
	assert.Equal(t, specs.Players, gameResponse.Quiz.Specs.Players)
	assert.Equal(t, specs.Teams, gameResponse.Quiz.Specs.Teams)
	assert.Equal(t, 2, gameResponse.Quiz.Specs.Rounds)
	assert.Equal(t, specs.Questions, gameResponse.Quiz.Specs.Questions)
	assert.Equal(t, 16, gameResponse.Quiz.Specs.Points)
	return gameResponse
}

func Join(t *testing.T, client *comms.Client, player models.Player, quizId string) models.GameResponse {
	request := models.Request{
		Action: "JOIN",
		Person: player,
		QuizId: quizId,
	}
	msg := tests.TestMessage(actions.SPECS, string(comms.SerializeMessage(request)))
	res, _, _ := client.Hub.Handle(msg, client)
	assert.NotNil(t, res, "response must be present")
	assert.Equal(t, "S_JOIN", res.Action)
	gameResponse, err := models.DecodeGameResponseJsonString(res.Content)
	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, "PLAYER", gameResponse.Role)
	assert.Equal(t, quizId, gameResponse.Quiz.Id)
	return gameResponse
}
