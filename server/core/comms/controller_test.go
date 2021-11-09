package comms_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/comms"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/tests"
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
}

func TestLogin(t *testing.T) {
	t.Run("Handle login", func(t *testing.T) {
		hub := tests.RunningHub(t)

		client := tests.NewClient(t, hub)
		player := tests.TestPlayer()
		LoginAndAssert(t, player, client)
	})

	t.Run("Handle login and create quiz", func(t *testing.T) {
		hub := tests.RunningHub(t)

		player, client := loggedInPlayer(t, hub)

		specs := tests.TestSpecs()
		CreateAndAssert(t, client, player, specs)
	})

	t.Run("Handle Login with 3 players", func(t *testing.T) {
		hub := tests.RunningHub(t)
		loggedInPlayer(t, hub)
		loggedInPlayer(t, hub)
		loggedInPlayer(t, hub)
	})
}

func TestJoinGame(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	t.Run("create a quiz with 2 players", func(t *testing.T) {
		hub := tests.RunningHub(t)

		quizmaster, client1 := loggedInPlayer(t, hub)
		player1, client2 := loggedInPlayer(t, hub)
		player2, client3 := loggedInPlayer(t, hub)

		specs := tests.TestSpecs()
		createRes := CreateAndAssert(t, client1, quizmaster, specs)
		JoinAndAssert(t, client2, player1, createRes.Quiz.Id)
		JoinAndAssert(t, client3, player2, createRes.Quiz.Id)
	})
}

func TestStartGame(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	t.Run("fail when starting a quiz with less players", func(t *testing.T) {
		hub := tests.RunningHub(t)

		quizmaster, client1 := loggedInPlayer(t, hub)
		player1, client2 := loggedInPlayer(t, hub)
		player2, client3 := loggedInPlayer(t, hub)

		specs := tests.TestSpecs()
		createRes := CreateAndAssert(t, client1, quizmaster, specs)
		JoinAndAssert(t, client2, player1, createRes.Quiz.Id)
		JoinAndAssert(t, client3, player2, createRes.Quiz.Id)
		StartAndAssertError(t, client1, quizmaster, createRes.Quiz.Id, "waiting for more players to join")
	})

	t.Run("fail when player starting the quiz", func(t *testing.T) {
		hub := tests.RunningHub(t)

		quizmaster, client1 := loggedInPlayer(t, hub)
		player1, client2 := loggedInPlayer(t, hub)
		player2, client3 := loggedInPlayer(t, hub)

		specs := tests.TestSpecs()
		specs.Players = 1
		createRes := CreateAndAssert(t, client1, quizmaster, specs)
		JoinAndAssert(t, client2, player1, createRes.Quiz.Id)
		JoinAndAssert(t, client3, player2, createRes.Quiz.Id)
		StartAndAssertError(t, client2, player1, createRes.Quiz.Id, "only quizmaster can start the match")
	})

	t.Run("start the quiz as quizmaster ", func(t *testing.T) {
		hub := tests.RunningHub(t)

		quizmaster, client1 := loggedInPlayer(t, hub)
		player1, client2 := loggedInPlayer(t, hub)
		player2, client3 := loggedInPlayer(t, hub)

		specs := tests.TestSpecs()
		specs.Players = 1
		createRes := CreateAndAssert(t, client1, quizmaster, specs)
		JoinAndAssert(t, client2, player1, createRes.Quiz.Id)
		JoinAndAssert(t, client3, player2, createRes.Quiz.Id)
		// StartAndAssert(t, client1, quizmaster, createRes.Quiz.Id)
	})
}

func loggedInPlayer(t *testing.T, hub *comms.Hub) (models.Player, *comms.Client) {
	client := tests.NewClient(t, hub)
	player := tests.TestPlayer()
	LoginAndAssert(t, player, client)
	return player, client
}

func LoginAndAssert(t *testing.T, player models.Player, client *comms.Client) models.LoginResponse {
	request := models.Request{
		Action: "BEGIN",
		Person: player,
	}
	res, _, err := client.Hub.Handle(tests.TestMessage(actions.BEGIN, string(comms.SerializeMessage(request))), client)
	assert.Nil(t, err, "error must be nil")
	assert.NotNil(t, res, "response must be present")
	assert.Equal(t, "S_PLAYER", res.Action)
	loginResponse, err := models.DecodeLoginResponseJsonString(res.Content)
	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, player.Id, loginResponse.Player.Id)
	assert.Equal(t, player.Name, loginResponse.Player.Name)
	return loginResponse
}

func CreateAndAssert(t *testing.T, client *comms.Client, player models.Player, specs models.Specs) models.GameResponse {
	request := models.Request{
		Action: "SPECS",
		Person: player,
		Specs:  specs,
	}
	msg := tests.TestMessage(actions.SPECS, string(comms.SerializeMessage(request)))
	res, _, err := client.Hub.Handle(msg, client)
	assert.Nil(t, err, "error must be nil")
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

func JoinAndAssert(t *testing.T, client *comms.Client, player models.Player, quizId string) models.GameResponse {
	request := models.Request{
		Action: "JOIN",
		Person: player,
		QuizId: quizId,
	}
	msg := tests.TestMessage(actions.JOIN, string(comms.SerializeMessage(request)))
	res, _, err := client.Hub.Handle(msg, client)
	assert.Nil(t, err, "error must be nil")
	assert.NotNil(t, res, "response must be present")
	assert.Equal(t, "S_JOIN", res.Action)
	gameResponse, err := models.DecodeGameResponseJsonString(res.Content)
	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, "PLAYER", gameResponse.Role)
	assert.Equal(t, quizId, gameResponse.Quiz.Id)
	return gameResponse
}

func StartAndAssert(t *testing.T, client *comms.Client, player models.Player, quizId string) models.Snapshot {
	request := models.Request{
		Action: "START",
		Person: player,
		QuizId: quizId,
	}
	msg := tests.TestMessage(actions.START, string(comms.SerializeMessage(request)))
	res, _, err := client.Hub.Handle(msg, client)
	assert.Nil(t, err, "error must be nil")
	assert.NotNil(t, res, "response must be present")
	assert.Equal(t, "S_GAME", res.Action)
	gameResponse, err := models.DecodeSnapshotJsonString(res.Content)
	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, quizId, gameResponse.QuizId)
	return gameResponse
}

func StartAndAssertError(t *testing.T, client *comms.Client, player models.Player, quizId string, errr string) {
	request := models.Request{
		Action: "START",
		Person: player,
		QuizId: quizId,
	}
	msg := tests.TestMessage(actions.START, string(comms.SerializeMessage(request)))
	_, _, err := client.Hub.Handle(msg, client)
	assert.NotNil(t, err, "error must be present")
	assert.Equal(t, errr, err.Error())
}
